// Package docpacks manages downloadable official documentation packs.
package docpacks

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"
)

const (
	// ProductionCatalogURL is the official documentation pack catalog.
	ProductionCatalogURL = "https://raw.githubusercontent.com/sairaph/apis-mcp/main/packs/catalog.json"

	catalogSchemaVersion = 1
	maxCatalogBytes      = 8 << 20
	maxCatalogPacks      = 10_000
	maxArchiveBytes      = 256 << 20
	maxUncompressedBytes = 1 << 30
	maxArchiveFiles      = 100_000
	defaultHTTPTimeout   = 10 * time.Minute
)

// Catalog describes the available official documentation packs.
type Catalog struct {
	SchemaVersion int    `json:"schema_version"`
	Packs         []Pack `json:"packs"`
}

// Pack is validated catalog metadata for one documentation family archive.
type Pack struct {
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

// ValidateCatalogBytes applies the production catalog decoder and validation
// rules to a generated catalog.
func ValidateCatalogBytes(raw []byte) (Catalog, error) {
	return decodeCatalog(raw)
}

func decodeCatalog(raw []byte) (Catalog, error) {
	var catalog Catalog
	if err := decodeStrict(raw, &catalog); err != nil {
		return Catalog{}, fmt.Errorf("parse pack catalog: %w", err)
	}
	if err := validateCatalog(catalog); err != nil {
		return Catalog{}, err
	}
	return catalog, nil
}

func decodeStrict(raw []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("multiple JSON values")
		}
		return err
	}
	return nil
}

func validateCatalog(catalog Catalog) error {
	if catalog.SchemaVersion != catalogSchemaVersion {
		return fmt.Errorf("unsupported pack catalog schema %d", catalog.SchemaVersion)
	}
	if catalog.Packs == nil {
		return errors.New("pack catalog is missing packs")
	}
	if len(catalog.Packs) > maxCatalogPacks {
		return fmt.Errorf("pack catalog has too many entries: %d", len(catalog.Packs))
	}
	ids := make(map[string]bool, len(catalog.Packs))
	assets := make(map[string]bool, len(catalog.Packs))
	for index, pack := range catalog.Packs {
		if err := validatePack(pack); err != nil {
			return fmt.Errorf("pack %d: %w", index, err)
		}
		if ids[pack.ID] {
			return fmt.Errorf("duplicate pack ID %q", pack.ID)
		}
		if assets[pack.Asset] {
			return fmt.Errorf("duplicate pack asset %q", pack.Asset)
		}
		ids[pack.ID] = true
		assets[pack.Asset] = true
	}
	return nil
}

func validatePack(pack Pack) error {
	if !validID(pack.ID) {
		return fmt.Errorf("invalid ID %q", pack.ID)
	}
	if pack.Name == "" || strings.TrimSpace(pack.Name) != pack.Name {
		return errors.New("name must be non-empty and trimmed")
	}
	if familyID(pack.Name) != pack.ID {
		return errors.New("name disagrees with the pack ID")
	}
	if strings.TrimSpace(pack.Description) != pack.Description {
		return errors.New("description must be trimmed")
	}
	decodedHash, err := hex.DecodeString(pack.SHA256)
	if err != nil || len(decodedHash) != 32 || strings.ToLower(pack.SHA256) != pack.SHA256 {
		return fmt.Errorf("invalid SHA-256 %q", pack.SHA256)
	}
	expectedAsset := pack.ID + "-" + pack.SHA256 + ".zip"
	if pack.Asset != expectedAsset || path.Base(pack.Asset) != pack.Asset || strings.ContainsAny(pack.Asset, `/\`) {
		return fmt.Errorf("invalid asset %q", pack.Asset)
	}
	if pack.Bytes <= 0 || pack.Bytes > maxArchiveBytes {
		return fmt.Errorf("invalid compressed byte count %d", pack.Bytes)
	}
	if pack.UncompressedBytes <= 0 || pack.UncompressedBytes > maxUncompressedBytes {
		return fmt.Errorf("invalid uncompressed byte count %d", pack.UncompressedBytes)
	}
	if pack.Files <= 0 || pack.Files > maxArchiveFiles || pack.Pages < 0 || pack.Pages >= pack.Files {
		return fmt.Errorf("invalid file/page counts %d/%d", pack.Files, pack.Pages)
	}
	if len(pack.Versions) == 0 || pack.Files != pack.Pages+len(pack.Versions) {
		return errors.New("file count must equal pages plus version manifests")
	}
	if err := validateSortedUnique(pack.Versions, validVersion, "version"); err != nil {
		return err
	}
	if err := validateSortedUnique(pack.Collections, validCollection, "collection"); err != nil {
		return err
	}
	return nil
}

func validateSortedUnique(values []string, valid func(string) bool, label string) error {
	for index, value := range values {
		if !valid(value) {
			return fmt.Errorf("invalid %s %q", label, value)
		}
		if index > 0 && values[index-1] >= value {
			return fmt.Errorf("%ss must be sorted and unique", label)
		}
	}
	return nil
}

func validID(value string) bool {
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

func catalogPackMap(catalog Catalog) map[string]Pack {
	packs := make(map[string]Pack, len(catalog.Packs))
	for _, pack := range catalog.Packs {
		packs[pack.ID] = pack
	}
	return packs
}

func validateCatalogURL(value string) (*url.URL, error) {
	parsed, err := url.Parse(value)
	if err != nil || parsed.Scheme != "http" && parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil || parsed.Fragment != "" {
		return nil, fmt.Errorf("invalid catalog URL %q", value)
	}
	if path.Base(parsed.Path) == "." || path.Base(parsed.Path) == "/" {
		return nil, fmt.Errorf("catalog URL must identify a file: %q", value)
	}
	return parsed, nil
}

func siblingURL(catalogURL *url.URL, asset string) string {
	result := *catalogURL
	result.Path = path.Join(path.Dir(catalogURL.Path), asset)
	if strings.HasPrefix(catalogURL.Path, "/") && !strings.HasPrefix(result.Path, "/") {
		result.Path = "/" + result.Path
	}
	result.RawPath = ""
	result.RawQuery = ""
	result.Fragment = ""
	return result.String()
}

func sameCatalogDirectory(base, candidate *url.URL) bool {
	return base.Scheme == candidate.Scheme && base.Host == candidate.Host && path.Dir(base.Path) == path.Dir(candidate.Path)
}

func clientForCatalog(source *http.Client, catalogURL *url.URL) *http.Client {
	if source == nil {
		source = &http.Client{Timeout: defaultHTTPTimeout}
	}
	client := *source
	originalRedirect := source.CheckRedirect
	client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if !sameCatalogDirectory(catalogURL, request.URL) {
			return errors.New("pack request redirect left the catalog directory")
		}
		if originalRedirect != nil {
			return originalRedirect(request, via)
		}
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return nil
	}
	return &client
}

func sortedPackIDs(packs map[string]Pack) []string {
	ids := make([]string, 0, len(packs))
	for id := range packs {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
