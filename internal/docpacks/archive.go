package docpacks

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

const maxManifestBytes = 1 << 20

type packManifest struct {
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Collections []string `yaml:"collections"`
}

// ValidateArchiveFile applies the production integrity and archive validation
// rules to a generated pack file.
func ValidateArchiveFile(name string, pack Pack) error {
	if err := validatePack(pack); err != nil {
		return err
	}
	return verifyBlob(name, pack)
}

func verifyBlob(name string, pack Pack) error {
	info, err := os.Lstat(name)
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() || info.Size() != pack.Bytes {
		return fmt.Errorf("pack blob is not a regular %d-byte file", pack.Bytes)
	}
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	hash := sha256.New()
	_, copyErr := io.Copy(hash, file)
	closeErr := file.Close()
	if err := errors.Join(copyErr, closeErr); err != nil {
		return err
	}
	if actual := hex.EncodeToString(hash.Sum(nil)); actual != pack.SHA256 {
		return fmt.Errorf("pack SHA-256 is %s, expected %s", actual, pack.SHA256)
	}
	return validateArchive(name, pack)
}

func validateArchive(name string, pack Pack) error {
	reader, err := zip.OpenReader(name)
	if err != nil {
		return fmt.Errorf("open pack ZIP: %w", err)
	}
	defer reader.Close()
	if len(reader.File) != pack.Files {
		return fmt.Errorf("pack ZIP has %d files, expected %d", len(reader.File), pack.Files)
	}

	versions := make(map[string]bool, len(pack.Versions))
	for _, version := range pack.Versions {
		versions[version] = true
	}
	manifests := make(map[string]bool, len(pack.Versions))
	seen := make(map[string]bool, len(reader.File))
	var uncompressed int64
	pages := 0
	for _, file := range reader.File {
		if err := validateArchiveName(file, pack.ID, versions); err != nil {
			return err
		}
		if seen[file.Name] {
			return fmt.Errorf("pack ZIP contains duplicate entry %q", file.Name)
		}
		seen[file.Name] = true
		if file.UncompressedSize64 > uint64(maxUncompressedBytes) || uncompressed > maxUncompressedBytes-int64(file.UncompressedSize64) {
			return errors.New("pack ZIP exceeds the uncompressed byte limit")
		}

		entry, err := file.Open()
		if err != nil {
			return fmt.Errorf("open pack ZIP entry %s: %w", file.Name, err)
		}
		parts := strings.Split(file.Name, "/")
		isManifest := path.Base(file.Name) == "_index.md"
		remaining := pack.UncompressedBytes - uncompressed
		if remaining < 0 {
			entry.Close()
			return errors.New("pack ZIP exceeds its declared uncompressed size")
		}
		var read int64
		if isManifest {
			limit := min(remaining, int64(maxManifestBytes))
			raw, readErr := io.ReadAll(io.LimitReader(entry, limit+1))
			read = int64(len(raw))
			if readErr == nil && read > limit {
				readErr = errors.New("manifest exceeds its byte limit")
			}
			if readErr == nil {
				readErr = validateManifest(raw, pack, parts[1])
			}
			err = errors.Join(readErr, entry.Close())
			if err == nil {
				manifests[parts[1]] = true
			}
		} else {
			read, err = io.Copy(io.Discard, io.LimitReader(entry, remaining+1))
			if err == nil && read > remaining {
				err = errors.New("pack ZIP exceeds its declared uncompressed size")
			}
			err = errors.Join(err, entry.Close())
			pages++
		}
		if err != nil {
			return fmt.Errorf("read pack ZIP entry %s: %w", file.Name, err)
		}
		if read != int64(file.UncompressedSize64) {
			return fmt.Errorf("pack ZIP entry %s size is %d, expected %d", file.Name, read, file.UncompressedSize64)
		}
		uncompressed += read
	}
	if pages != pack.Pages {
		return fmt.Errorf("pack ZIP has %d pages, expected %d", pages, pack.Pages)
	}
	if uncompressed != pack.UncompressedBytes {
		return fmt.Errorf("pack ZIP has %d uncompressed bytes, expected %d", uncompressed, pack.UncompressedBytes)
	}
	for _, version := range pack.Versions {
		if !manifests[version] {
			return fmt.Errorf("pack ZIP has no manifest for version %q", version)
		}
	}
	return nil
}

func validateArchiveName(file *zip.File, id string, versions map[string]bool) error {
	name := file.Name
	if name == "." || !fs.ValidPath(name) || strings.Contains(name, "\\") || path.Clean(name) != name {
		return fmt.Errorf("pack ZIP contains invalid path %q", name)
	}
	if file.FileInfo().IsDir() || !file.Mode().IsRegular() {
		return fmt.Errorf("pack ZIP entry is not a regular file: %q", name)
	}
	if path.Ext(name) != ".md" {
		return fmt.Errorf("pack ZIP contains non-Markdown entry %q", name)
	}
	parts := strings.Split(name, "/")
	if len(parts) < 3 || parts[0] != id || !versions[parts[1]] {
		return fmt.Errorf("pack ZIP entry is outside declared family versions: %q", name)
	}
	if path.Base(name) == "_index.md" && len(parts) != 3 {
		return fmt.Errorf("pack ZIP contains nested manifest %q", name)
	}
	return nil
}

func validateManifest(raw []byte, pack Pack, version string) error {
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return errors.New("manifest frontmatter is missing")
	}
	closing := -1
	for index := 1; index < len(lines); index++ {
		if strings.TrimSpace(lines[index]) == "---" {
			closing = index
			break
		}
	}
	if closing < 0 {
		return errors.New("manifest frontmatter is not closed")
	}
	var manifest packManifest
	if err := yaml.Unmarshal([]byte(strings.Join(lines[1:closing], "\n")), &manifest); err != nil {
		return fmt.Errorf("parse manifest frontmatter: %w", err)
	}
	manifest.Name = strings.TrimSpace(manifest.Name)
	manifest.Version = strings.TrimSpace(manifest.Version)
	manifest.Description = strings.TrimSpace(manifest.Description)
	collections := cleanCollections(manifest.Collections)
	if manifest.Name != pack.Name || manifest.Version != version || manifest.Description != pack.Description || !equalStrings(collections, pack.Collections) {
		return errors.New("manifest metadata disagrees with the pack catalog")
	}
	if familyID(manifest.Name) != pack.ID {
		return errors.New("manifest name disagrees with the pack ID")
	}
	return nil
}

func cleanCollections(values []string) []string {
	seen := make(map[string]bool, len(values))
	cleaned := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		cleaned = append(cleaned, value)
	}
	sort.Strings(cleaned)
	return cleaned
}

func familyID(value string) string {
	var result strings.Builder
	dash := false
	for _, character := range strings.ToLower(strings.TrimSpace(value)) {
		if character >= 'a' && character <= 'z' || character >= '0' && character <= '9' {
			if dash && result.Len() > 0 {
				result.WriteByte('-')
			}
			result.WriteRune(character)
			dash = false
			continue
		}
		if unicode.IsLetter(character) || unicode.IsDigit(character) {
			dash = true
			continue
		}
		dash = true
	}
	return strings.Trim(result.String(), "-")
}

func equalStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
