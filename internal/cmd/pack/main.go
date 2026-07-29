package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/sairaph/apis-mcp/internal/docpacks"
	productionlibrary "github.com/sairaph/apis-mcp/library"
	"gopkg.in/yaml.v3"
)

const catalogSchemaVersion = 1

var archiveTime = time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)

type catalog struct {
	SchemaVersion int         `json:"schema_version"`
	Packs         []packEntry `json:"packs"`
}

type packEntry struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Asset             string   `json:"asset"`
	SHA256            string   `json:"sha256"`
	Bytes             int64    `json:"bytes"`
	UncompressedBytes int64    `json:"uncompressed_bytes"`
	Files             int      `json:"files"`
	Pages             int      `json:"pages"`
	Versions          []string `json:"versions"`
	Collections       []string `json:"collections"`
}

type manifest struct {
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Collections []string `yaml:"collections"`
}

type sourceFile struct {
	name string
	raw  []byte
}

type artifact struct {
	name string
	raw  []byte
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUsage(stderr)
		return 2
	}

	mode := args[0]
	flags := flag.NewFlagSet(mode, flag.ContinueOnError)
	flags.SetOutput(stderr)
	source := flags.String("source", "library/builtin", "canonical Markdown source root")
	output := flags.String("out", "packs", "generated pack directory")
	if err := flags.Parse(args[1:]); err != nil {
		return 2
	}
	if flags.NArg() != 0 || mode != "generate" && mode != "verify" {
		printUsage(stderr)
		return 2
	}

	var (
		count int
		err   error
	)
	if mode == "generate" {
		count, err = generatePacks(*source, *output)
	} else {
		count, err = verifyPacks(*source, *output)
	}
	if err != nil {
		fmt.Fprintln(stderr, "pack:", err)
		return 1
	}
	fmt.Fprintf(stdout, "%s %d packs in %s\n", modePastTense(mode), count, *output)
	return 0
}

func printUsage(writer io.Writer) {
	fmt.Fprintln(writer, "usage: pack <generate|verify> [-source DIR] [-out DIR]")
}

func modePastTense(mode string) string {
	if mode == "generate" {
		return "generated"
	}
	return "verified"
}

func generatePacks(sourceRoot, outputRoot string) (int, error) {
	artifacts, count, err := buildArtifacts(sourceRoot)
	if err != nil {
		return 0, err
	}
	if err := validateArtifactBytes(artifacts); err != nil {
		return 0, err
	}
	if err := writeArtifacts(outputRoot, artifacts); err != nil {
		return 0, err
	}
	return count, nil
}

func verifyPacks(sourceRoot, outputRoot string) (int, error) {
	artifacts, count, err := buildArtifacts(sourceRoot)
	if err != nil {
		return 0, err
	}
	if err := validateArtifactBytes(artifacts); err != nil {
		return 0, err
	}
	entries, err := os.ReadDir(outputRoot)
	if err != nil {
		return 0, fmt.Errorf("read generated directory: %w", err)
	}
	expected := make(map[string][]byte, len(artifacts))
	var currentArchives []string
	for _, item := range artifacts {
		expected[item.name] = item.raw
	}
	for _, entry := range entries {
		want, ok := expected[entry.Name()]
		if !ok {
			if err := validateContentAddressedArchive(filepath.Join(outputRoot, entry.Name()), entry); err != nil {
				return 0, err
			}
			continue
		}
		if entry.Type()&os.ModeSymlink != 0 || entry.IsDir() {
			return 0, fmt.Errorf("artifact is not a regular file: %s", filepath.Join(outputRoot, entry.Name()))
		}
		got, err := os.ReadFile(filepath.Join(outputRoot, entry.Name()))
		if err != nil {
			return 0, fmt.Errorf("read artifact %s: %w", entry.Name(), err)
		}
		if !bytes.Equal(got, want) {
			return 0, fmt.Errorf("artifact is out of date: %s; run pack generate", filepath.Join(outputRoot, entry.Name()))
		}
		if path.Ext(entry.Name()) == ".zip" {
			currentArchives = append(currentArchives, filepath.Join(outputRoot, entry.Name()))
		}
		delete(expected, entry.Name())
	}
	if len(expected) != 0 {
		missing := make([]string, 0, len(expected))
		for name := range expected {
			missing = append(missing, filepath.Join(outputRoot, name))
		}
		sort.Strings(missing)
		return 0, fmt.Errorf("missing artifacts: %s; run pack generate", strings.Join(missing, ", "))
	}
	if err := validatePackArchives(currentArchives); err != nil {
		return 0, err
	}
	return count, nil
}

func buildArtifacts(sourceRoot string) ([]artifact, int, error) {
	entries, err := os.ReadDir(sourceRoot)
	if err != nil {
		return nil, 0, fmt.Errorf("read source root: %w", err)
	}
	if len(entries) == 0 {
		return nil, 0, errors.New("source root contains no API families")
	}

	generated := make([]artifact, 0, len(entries)+1)
	catalogEntries := make([]packEntry, 0, len(entries))
	packIDs := make(map[string]string, len(entries))
	for _, entry := range entries {
		if entry.Type()&os.ModeSymlink != 0 || !entry.IsDir() {
			return nil, 0, fmt.Errorf("source root may contain only family directories: %s", entry.Name())
		}
		if !validFamilyID(entry.Name()) {
			return nil, 0, fmt.Errorf("invalid family directory %q", entry.Name())
		}
		pack, raw, err := buildPack(sourceRoot, entry.Name())
		if err != nil {
			return nil, 0, err
		}
		if previous, exists := packIDs[pack.ID]; exists {
			return nil, 0, fmt.Errorf("source families %q and %q produce duplicate ID %q", previous, entry.Name(), pack.ID)
		}
		packIDs[pack.ID] = entry.Name()
		sum := sha256.Sum256(raw)
		pack.SHA256 = hex.EncodeToString(sum[:])
		pack.Bytes = int64(len(raw))
		pack.Asset = pack.ID + "-" + pack.SHA256 + ".zip"
		generated = append(generated, artifact{name: pack.Asset, raw: raw})
		catalogEntries = append(catalogEntries, pack)
	}
	sort.Slice(catalogEntries, func(i, j int) bool { return catalogEntries[i].ID < catalogEntries[j].ID })

	var catalogJSON bytes.Buffer
	encoder := json.NewEncoder(&catalogJSON)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(catalog{SchemaVersion: catalogSchemaVersion, Packs: catalogEntries}); err != nil {
		return nil, 0, fmt.Errorf("encode catalog: %w", err)
	}
	generated = append(generated, artifact{name: "catalog.json", raw: catalogJSON.Bytes()})
	sort.Slice(generated, func(i, j int) bool { return generated[i].name < generated[j].name })
	return generated, len(catalogEntries), nil
}

func buildPack(sourceRoot, familyDirectory string) (packEntry, []byte, error) {
	familyRoot := filepath.Join(sourceRoot, familyDirectory)
	versionEntries, err := os.ReadDir(familyRoot)
	if err != nil {
		return packEntry{}, nil, fmt.Errorf("read family %s: %w", familyDirectory, err)
	}
	if len(versionEntries) == 0 {
		return packEntry{}, nil, fmt.Errorf("family %s contains no versions", familyDirectory)
	}

	var familyManifest manifest
	var familyID string
	versions := make([]string, 0, len(versionEntries))
	for index, entry := range versionEntries {
		if entry.Type()&os.ModeSymlink != 0 || !entry.IsDir() {
			return packEntry{}, nil, fmt.Errorf("family %s may contain only version directories: %s", familyDirectory, entry.Name())
		}
		if !validVersion(entry.Name()) {
			return packEntry{}, nil, fmt.Errorf("family %s has invalid version directory %q", familyDirectory, entry.Name())
		}
		manifestPath := filepath.Join(familyRoot, entry.Name(), "_index.md")
		info, err := os.Lstat(manifestPath)
		if err != nil {
			return packEntry{}, nil, fmt.Errorf("family %s version %s requires _index.md: %w", familyDirectory, entry.Name(), err)
		}
		if !info.Mode().IsRegular() {
			return packEntry{}, nil, fmt.Errorf("manifest is not a regular file: %s", manifestPath)
		}
		raw, err := os.ReadFile(manifestPath)
		if err != nil {
			return packEntry{}, nil, fmt.Errorf("read manifest %s: %w", manifestPath, err)
		}
		metadata, err := parseManifest(raw)
		if err != nil {
			return packEntry{}, nil, fmt.Errorf("parse manifest %s: %w", manifestPath, err)
		}
		if metadata.Version != entry.Name() {
			return packEntry{}, nil, fmt.Errorf("manifest %s declares version %q, expected %q", manifestPath, metadata.Version, entry.Name())
		}
		metadataID := slug(metadata.Name)
		if !validFamilyID(metadataID) {
			return packEntry{}, nil, fmt.Errorf("manifest %s name %q does not produce a valid family ID", manifestPath, metadata.Name)
		}
		if index == 0 {
			familyManifest = metadata
			familyID = metadataID
		} else if familyManifest.Name != metadata.Name || familyManifest.Description != metadata.Description || !equalStrings(familyManifest.Collections, metadata.Collections) {
			return packEntry{}, nil, fmt.Errorf("versions of family %s disagree on name, description, or collections", familyDirectory)
		}
		versions = append(versions, metadata.Version)
	}
	sort.Strings(versions)

	var files []sourceFile
	var uncompressedBytes int64
	pages := 0
	err = filepath.WalkDir(familyRoot, func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(familyRoot, name)
		if err != nil {
			return err
		}
		if relative == "." {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("symbolic links are not allowed: %s", name)
		}
		portable := filepath.ToSlash(relative)
		if !fs.ValidPath(portable) || strings.Contains(portable, "\\") || path.Clean(portable) != portable {
			return fmt.Errorf("invalid source path: %s", name)
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("source is not a regular file: %s", name)
		}
		if path.Ext(portable) != ".md" {
			return fmt.Errorf("source contains non-Markdown file: %s", name)
		}
		if path.Base(portable) == "_index.md" && len(strings.Split(portable, "/")) != 2 {
			return fmt.Errorf("nested manifest is not allowed: %s", name)
		}
		raw, err := os.ReadFile(name)
		if err != nil {
			return err
		}
		archiveName := path.Join(familyID, portable)
		files = append(files, sourceFile{name: archiveName, raw: raw})
		uncompressedBytes += int64(len(raw))
		if path.Base(portable) != "_index.md" {
			pages++
		}
		return nil
	})
	if err != nil {
		return packEntry{}, nil, fmt.Errorf("walk family %s: %w", familyDirectory, err)
	}
	sort.Slice(files, func(i, j int) bool { return files[i].name < files[j].name })
	raw, err := writeZIP(files)
	if err != nil {
		return packEntry{}, nil, fmt.Errorf("build family %s: %w", familyDirectory, err)
	}
	return packEntry{
		ID:                familyID,
		Name:              familyManifest.Name,
		Description:       familyManifest.Description,
		UncompressedBytes: uncompressedBytes,
		Files:             len(files),
		Pages:             pages,
		Versions:          versions,
		Collections:       append([]string(nil), familyManifest.Collections...),
	}, raw, nil
}

func parseManifest(raw []byte) (manifest, error) {
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return manifest{}, errors.New("YAML frontmatter must start on the first line")
	}
	closing := -1
	for index := 1; index < len(lines); index++ {
		if strings.TrimSpace(lines[index]) == "---" {
			closing = index
			break
		}
	}
	if closing < 0 {
		return manifest{}, errors.New("YAML frontmatter is missing its closing delimiter")
	}
	var metadata manifest
	if err := yaml.Unmarshal([]byte(strings.Join(lines[1:closing], "\n")), &metadata); err != nil {
		return manifest{}, err
	}
	metadata.Name = strings.TrimSpace(metadata.Name)
	metadata.Version = strings.TrimSpace(metadata.Version)
	metadata.Description = strings.TrimSpace(metadata.Description)
	if metadata.Name == "" || metadata.Version == "" {
		return manifest{}, errors.New("manifest requires name and version")
	}
	collections, err := cleanCollections(metadata.Collections)
	if err != nil {
		return manifest{}, err
	}
	metadata.Collections = collections
	return metadata, nil
}

func cleanCollections(values []string) ([]string, error) {
	seen := make(map[string]bool, len(values))
	cleaned := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		if !validCollection(value) {
			return nil, fmt.Errorf("invalid collection %q", value)
		}
		seen[value] = true
		cleaned = append(cleaned, value)
	}
	sort.Strings(cleaned)
	return cleaned, nil
}

func validCollection(value string) bool {
	if value == "" || value[0] == '_' || value[len(value)-1] == '_' {
		return false
	}
	underscore := false
	for _, character := range value {
		if character == '_' {
			if underscore {
				return false
			}
			underscore = true
			continue
		}
		if character < 'a' || character > 'z' {
			if character < '0' || character > '9' {
				return false
			}
		}
		underscore = false
	}
	return true
}

func validFamilyID(value string) bool {
	if value == "" || value[0] == '-' || value[len(value)-1] == '-' {
		return false
	}
	dash := false
	for _, character := range value {
		if character == '-' {
			if dash {
				return false
			}
			dash = true
			continue
		}
		if character < 'a' || character > 'z' {
			if character < '0' || character > '9' {
				return false
			}
		}
		dash = false
	}
	return true
}

func slug(value string) string {
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

func validVersion(value string) bool {
	if value == "" || value == "." || value == ".." || value[0] == '.' {
		return false
	}
	for _, character := range value {
		if character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' || character >= '0' && character <= '9' {
			continue
		}
		if character != '.' && character != '_' && character != '-' {
			return false
		}
	}
	return true
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

func writeZIP(files []sourceFile) ([]byte, error) {
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	writer.RegisterCompressor(zip.Deflate, func(output io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(output, flate.BestCompression)
	})
	for _, file := range files {
		header := &zip.FileHeader{Name: file.name, Method: zip.Deflate}
		header.SetModTime(archiveTime)
		header.SetMode(0o644)
		entry, err := writer.CreateHeader(header)
		if err != nil {
			return nil, err
		}
		if _, err := entry.Write(file.raw); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func validateArtifactBytes(artifacts []artifact) error {
	directory, err := os.MkdirTemp("", "apis-mcp-packs-*")
	if err != nil {
		return fmt.Errorf("create pack validation directory: %w", err)
	}
	defer os.RemoveAll(directory)
	var catalogRaw []byte
	archivePaths := make(map[string]string)
	for _, item := range artifacts {
		if item.name == "catalog.json" {
			catalogRaw = item.raw
			continue
		}
		if path.Ext(item.name) != ".zip" {
			continue
		}
		name := filepath.Join(directory, item.name)
		if err := os.WriteFile(name, item.raw, 0o600); err != nil {
			return fmt.Errorf("write pack for validation: %w", err)
		}
		archivePaths[item.name] = name
	}
	catalog, err := docpacks.ValidateCatalogBytes(catalogRaw)
	if err != nil {
		return fmt.Errorf("validate catalog with production rules: %w", err)
	}
	archives := make([]string, 0, len(catalog.Packs))
	for _, pack := range catalog.Packs {
		name, ok := archivePaths[pack.Asset]
		if !ok {
			return fmt.Errorf("catalog references missing archive %s", pack.Asset)
		}
		if err := docpacks.ValidateArchiveFile(name, pack); err != nil {
			return fmt.Errorf("validate %s with production rules: %w", pack.Asset, err)
		}
		archives = append(archives, name)
	}
	return validatePackArchives(archives)
}

func validatePackArchives(archives []string) error {
	if len(archives) == 0 {
		return errors.New("no current pack archives to validate")
	}
	directory, err := os.MkdirTemp("", "apis-mcp-pack-index-*")
	if err != nil {
		return fmt.Errorf("create pack index directory: %w", err)
	}
	defer os.RemoveAll(directory)
	snapshot, err := productionlibrary.Open(context.Background(), productionlibrary.Options{
		IndexPath:    filepath.Join(directory, "library.sqlite"),
		PackArchives: archives,
	})
	if err != nil {
		return fmt.Errorf("validate packs with production library: %w", err)
	}
	if err := snapshot.Close(); err != nil {
		return fmt.Errorf("close pack validation index: %w", err)
	}
	return nil
}

func validateContentAddressedArchive(name string, entry fs.DirEntry) error {
	id, declaredHash, ok := parseContentAddressedName(entry.Name())
	if !ok {
		return fmt.Errorf("unexpected or malformed artifact %s; run pack generate", name)
	}
	if entry.Type()&os.ModeSymlink != 0 || entry.IsDir() {
		return fmt.Errorf("content-addressed artifact is not a regular file: %s", name)
	}
	info, err := entry.Info()
	if err != nil {
		return fmt.Errorf("inspect content-addressed artifact %s: %w", name, err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("content-addressed artifact is not a regular file: %s", name)
	}
	raw, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read content-addressed artifact %s: %w", name, err)
	}
	sum := sha256.Sum256(raw)
	if actual := hex.EncodeToString(sum[:]); actual != declaredHash {
		return fmt.Errorf("content-addressed artifact %s has SHA-256 %s, expected %s", name, actual, declaredHash)
	}
	reader, err := zip.OpenReader(name)
	if err != nil {
		return fmt.Errorf("open content-addressed artifact %s: %w", name, err)
	}
	defer reader.Close()
	if len(reader.File) == 0 {
		return fmt.Errorf("content-addressed artifact is empty: %s", name)
	}
	previous := ""
	for _, file := range reader.File {
		if file.Name <= previous || !fs.ValidPath(file.Name) || strings.Contains(file.Name, "\\") || path.Clean(file.Name) != file.Name || path.Ext(file.Name) != ".md" || !strings.HasPrefix(file.Name, id+"/") || file.FileInfo().IsDir() || !file.Mode().IsRegular() {
			return fmt.Errorf("content-addressed artifact %s has invalid entry %q", name, file.Name)
		}
		previous = file.Name
		content, err := file.Open()
		if err != nil {
			return fmt.Errorf("open content-addressed artifact entry %s: %w", file.Name, err)
		}
		_, copyErr := io.Copy(io.Discard, content)
		if err := errors.Join(copyErr, content.Close()); err != nil {
			return fmt.Errorf("read content-addressed artifact entry %s: %w", file.Name, err)
		}
	}
	return nil
}

func parseContentAddressedName(name string) (string, string, bool) {
	if path.Ext(name) != ".zip" || path.Base(name) != name {
		return "", "", false
	}
	base := strings.TrimSuffix(name, ".zip")
	if len(base) <= sha256.Size*2+1 || base[len(base)-sha256.Size*2-1] != '-' {
		return "", "", false
	}
	id := base[:len(base)-sha256.Size*2-1]
	hash := base[len(base)-sha256.Size*2:]
	decoded, err := hex.DecodeString(hash)
	if !validFamilyID(id) || err != nil || len(decoded) != sha256.Size || strings.ToLower(hash) != hash {
		return "", "", false
	}
	return id, hash, true
}

func legacyArchiveName(name string) bool {
	return path.Ext(name) == ".zip" && path.Base(name) == name && validFamilyID(strings.TrimSuffix(name, ".zip"))
}

func writeArtifacts(outputRoot string, artifacts []artifact) error {
	info, err := os.Stat(outputRoot)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("inspect output directory: %w", err)
	}
	if err == nil && !info.IsDir() {
		return fmt.Errorf("output is not a directory: %s", outputRoot)
	}
	if err := os.MkdirAll(outputRoot, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	expected := make(map[string]bool, len(artifacts))
	for _, item := range artifacts {
		expected[item.name] = true
	}
	entries, err := os.ReadDir(outputRoot)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if expected[entry.Name()] {
			continue
		}
		name := filepath.Join(outputRoot, entry.Name())
		if _, _, contentAddressed := parseContentAddressedName(entry.Name()); contentAddressed {
			if err := validateContentAddressedArchive(name, entry); err != nil {
				return err
			}
			continue
		}
		if !legacyArchiveName(entry.Name()) {
			return fmt.Errorf("output directory contains unmanaged entry: %s", name)
		}
		if err := os.Remove(name); err != nil {
			return fmt.Errorf("remove stale artifact %s: %w", entry.Name(), err)
		}
	}
	for _, item := range artifacts {
		if err := writeFileAtomic(filepath.Join(outputRoot, item.name), item.raw); err != nil {
			return err
		}
	}
	return nil
}

func writeFileAtomic(name string, raw []byte) error {
	temporary, err := os.CreateTemp(filepath.Dir(name), ".pack-*")
	if err != nil {
		return fmt.Errorf("create temporary artifact: %w", err)
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)
	if err := temporary.Chmod(0o644); err != nil {
		temporary.Close()
		return err
	}
	if _, err := temporary.Write(raw); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Remove(name); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("replace artifact %s: %w", name, err)
	}
	if err := os.Rename(temporaryName, name); err != nil {
		return fmt.Errorf("publish artifact %s: %w", name, err)
	}
	return nil
}
