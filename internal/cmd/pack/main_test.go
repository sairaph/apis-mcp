package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestGeneratePacksIsDeterministicAndCanonical(t *testing.T) {
	source := t.TempDir()
	writeManifest(t, source, "zeta/v1", "Zeta", "v1", "Zeta docs.", []string{"tools"})
	writePage(t, source, "zeta/v1/reference.md", "Reference", "reference", "Reference body.")
	writeManifest(t, source, "alpha-api/v2", "Alpha API", "v2", "Alpha docs.", []string{"developer_tools", "examples"})
	writePage(t, source, "alpha-api/v2/operations/list.md", "List", "list", "List body.")
	writeManifest(t, source, "alpha-api/v1", "Alpha API", "v1", "Alpha docs.", []string{"examples", "developer_tools"})
	writePage(t, source, "alpha-api/v1/overview.md", "Overview", "overview", "Overview body.")

	first := filepath.Join(t.TempDir(), "packs")
	second := filepath.Join(t.TempDir(), "packs")
	if count, err := generatePacks(source, first); err != nil || count != 2 {
		t.Fatalf("first generation: count=%d err=%v", count, err)
	}
	if count, err := generatePacks(source, second); err != nil || count != 2 {
		t.Fatalf("second generation: count=%d err=%v", count, err)
	}
	compareDirectories(t, first, second)
	if _, err := verifyPacks(source, first); err != nil {
		t.Fatalf("verify generated packs: %v", err)
	}

	rawCatalog := readSource(t, first, "catalog.json")
	var got catalog
	if err := json.Unmarshal(rawCatalog, &got); err != nil {
		t.Fatal(err)
	}
	if got.SchemaVersion != 1 || len(got.Packs) != 2 || got.Packs[0].ID != "alpha-api" || got.Packs[1].ID != "zeta" {
		t.Fatalf("unexpected catalog: %+v", got)
	}
	alpha := got.Packs[0]
	assetID, assetHash, ok := parseContentAddressedName(alpha.Asset)
	if alpha.Name != "Alpha API" || alpha.Description != "Alpha docs." || !ok || assetID != alpha.ID || assetHash != alpha.SHA256 {
		t.Fatalf("unexpected alpha metadata: %+v", alpha)
	}
	if !reflect.DeepEqual(alpha.Versions, []string{"v1", "v2"}) || !reflect.DeepEqual(alpha.Collections, []string{"developer_tools", "examples"}) {
		t.Fatalf("unsorted catalog arrays: %+v", alpha)
	}
	if alpha.Files != 4 || alpha.Pages != 2 || alpha.UncompressedBytes <= 0 {
		t.Fatalf("unexpected alpha counts: %+v", alpha)
	}

	archive := readSource(t, first, alpha.Asset)
	sum := sha256.Sum256(archive)
	if alpha.SHA256 != hex.EncodeToString(sum[:]) || alpha.Bytes != int64(len(archive)) {
		t.Fatalf("incorrect archive digest or size: %+v", alpha)
	}
	reader, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil {
		t.Fatal(err)
	}
	wantNames := []string{
		"alpha-api/v1/_index.md",
		"alpha-api/v1/overview.md",
		"alpha-api/v2/_index.md",
		"alpha-api/v2/operations/list.md",
	}
	var names []string
	var uncompressed int64
	for _, file := range reader.File {
		names = append(names, file.Name)
		uncompressed += int64(file.UncompressedSize64)
		if filepath.Ext(file.Name) != ".md" || file.Method != zip.Deflate {
			t.Fatalf("noncanonical ZIP entry: %+v", file.FileHeader)
		}
		if file.Mode().Perm() != 0o644 || !file.Modified.Equal(archiveTime) {
			t.Fatalf("unstable ZIP metadata for %s: mode=%v modified=%v", file.Name, file.Mode(), file.Modified)
		}
	}
	if !reflect.DeepEqual(names, wantNames) {
		t.Fatalf("ZIP paths = %v, want %v", names, wantNames)
	}
	if uncompressed != alpha.UncompressedBytes {
		t.Fatalf("ZIP uncompressed bytes = %d, catalog = %d", uncompressed, alpha.UncompressedBytes)
	}
}

func TestVerifyPacksRejectsStaleMissingAndUnexpectedArtifacts(t *testing.T) {
	source := t.TempDir()
	writeManifest(t, source, "alpha/v1", "Alpha", "v1", "Alpha docs.", []string{"examples"})
	writePage(t, source, "alpha/v1/page.md", "Page", "page", "Page body.")
	output := filepath.Join(t.TempDir(), "packs")
	if _, err := generatePacks(source, output); err != nil {
		t.Fatal(err)
	}
	current := readCatalog(t, output).Packs[0].Asset

	writeSource(t, output, current, "stale")
	if _, err := verifyPacks(source, output); err == nil || !strings.Contains(err.Error(), "out of date") {
		t.Fatalf("stale verification error = %v", err)
	}
	if _, err := generatePacks(source, output); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(output, "catalog.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := verifyPacks(source, output); err == nil || !strings.Contains(err.Error(), "missing artifacts") {
		t.Fatalf("missing verification error = %v", err)
	}
	if _, err := generatePacks(source, output); err != nil {
		t.Fatal(err)
	}
	writeSource(t, output, "unexpected.txt", "unexpected")
	if _, err := verifyPacks(source, output); err == nil || !strings.Contains(err.Error(), "unexpected or malformed") {
		t.Fatalf("unexpected verification error = %v", err)
	}
}

func TestGeneratePreservesPriorContentAddressedArchives(t *testing.T) {
	source := t.TempDir()
	writeManifest(t, source, "alpha/v1", "Alpha", "v1", "Alpha docs.", []string{"examples"})
	writePage(t, source, "alpha/v1/page.md", "Page", "page", "First body.")
	output := filepath.Join(t.TempDir(), "packs")
	if _, err := generatePacks(source, output); err != nil {
		t.Fatal(err)
	}
	oldAsset := readCatalog(t, output).Packs[0].Asset
	oldArchive := readSource(t, output, oldAsset)

	writePage(t, source, "alpha/v1/page.md", "Page", "page", "Second body.")
	writeSource(t, output, "alpha.zip", "legacy")
	if _, err := generatePacks(source, output); err != nil {
		t.Fatal(err)
	}
	newAsset := readCatalog(t, output).Packs[0].Asset
	if newAsset == oldAsset {
		t.Fatal("source change did not produce a new content-addressed asset")
	}
	if !bytes.Equal(readSource(t, output, oldAsset), oldArchive) {
		t.Fatal("prior content-addressed archive was not preserved exactly")
	}
	if _, err := os.Stat(filepath.Join(output, "alpha.zip")); !os.IsNotExist(err) {
		t.Fatalf("legacy archive still exists: %v", err)
	}
	if _, err := verifyPacks(source, output); err != nil {
		t.Fatalf("verify with preserved archive: %v", err)
	}

	writeSource(t, output, oldAsset, "corrupt")
	if _, err := verifyPacks(source, output); err == nil || !strings.Contains(err.Error(), "has SHA-256") {
		t.Fatalf("corrupt preserved archive error = %v", err)
	}
	if err := os.WriteFile(filepath.Join(output, oldAsset), oldArchive, 0o644); err != nil {
		t.Fatal(err)
	}
	writeSource(t, output, "alpha-not-a-hash.zip", "unmanaged")
	if _, err := verifyPacks(source, output); err == nil || !strings.Contains(err.Error(), "unexpected or malformed") {
		t.Fatalf("malformed archive error = %v", err)
	}
}

func TestGenerateAndVerifyUseProductionLibraryValidation(t *testing.T) {
	source := t.TempDir()
	writeManifest(t, source, "alpha/v1", "Alpha", "v1", "Alpha docs.", nil)
	writePage(t, source, "alpha/v1/one.md", "One", "duplicate", "First.")
	writePage(t, source, "alpha/v1/two.md", "Two", "duplicate", "Second.")
	output := filepath.Join(t.TempDir(), "packs")

	if _, err := generatePacks(source, output); err == nil || !strings.Contains(err.Error(), "duplicate explicit page_id") {
		t.Fatalf("generation semantic validation error = %v", err)
	}
	artifacts, _, err := buildArtifacts(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := writeArtifacts(output, artifacts); err != nil {
		t.Fatal(err)
	}
	if _, err := verifyPacks(source, output); err == nil || !strings.Contains(err.Error(), "duplicate explicit page_id") {
		t.Fatalf("verification semantic validation error = %v", err)
	}
}

func TestBuildArtifactsValidatesCanonicalFamilies(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(*testing.T, string)
		want    string
	}{
		{
			name: "version mismatch",
			prepare: func(t *testing.T, root string) {
				writeManifest(t, root, "alpha/v1", "Alpha", "v2", "Docs.", nil)
			},
			want: "declares version",
		},
		{
			name: "family metadata mismatch",
			prepare: func(t *testing.T, root string) {
				writeManifest(t, root, "alpha/v1", "Alpha", "v1", "First.", nil)
				writeManifest(t, root, "alpha/v2", "Alpha", "v2", "Second.", nil)
			},
			want: "disagree",
		},
		{
			name: "non Markdown source",
			prepare: func(t *testing.T, root string) {
				writeManifest(t, root, "alpha/v1", "Alpha", "v1", "Docs.", nil)
				writeSource(t, root, "alpha/v1/image.png", "not markdown")
			},
			want: "non-Markdown",
		},
		{
			name: "nested manifest",
			prepare: func(t *testing.T, root string) {
				writeManifest(t, root, "alpha/v1", "Alpha", "v1", "Docs.", nil)
				writeSource(t, root, "alpha/v1/nested/_index.md", "---\nname: Nested\nversion: v1\n---\n")
			},
			want: "nested manifest",
		},
		{
			name: "invalid family ID",
			prepare: func(t *testing.T, root string) {
				writeManifest(t, root, "Alpha/v1", "Alpha", "v1", "Docs.", nil)
			},
			want: "invalid family directory",
		},
		{
			name: "invalid collection",
			prepare: func(t *testing.T, root string) {
				writeManifest(t, root, "alpha/v1", "Alpha", "v1", "Docs.", []string{"Developer Tools"})
			},
			want: "invalid collection",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := t.TempDir()
			test.prepare(t, root)
			if _, _, err := buildArtifacts(root); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("build error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestBuildPackDerivesFamilyDirectoryIDFromManifestName(t *testing.T) {
	source := t.TempDir()
	writeManifest(t, source, "source-group/v1", "Alpha API", "v1", "Docs.", nil)
	writePage(t, source, "source-group/v1/page.md", "Page", "page", "Body.")
	artifacts, _, err := buildArtifacts(source)
	if err != nil {
		t.Fatal(err)
	}
	var got catalog
	for _, item := range artifacts {
		if item.name == "catalog.json" {
			if err := json.Unmarshal(item.raw, &got); err != nil {
				t.Fatal(err)
			}
		}
	}
	if len(got.Packs) != 1 || got.Packs[0].ID != "alpha-api" {
		t.Fatalf("catalog packs = %+v", got.Packs)
	}
	archive := artifactNamed(t, artifacts, got.Packs[0].Asset)
	reader, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range reader.File {
		if !strings.HasPrefix(file.Name, "alpha-api/") {
			t.Fatalf("archive family directory does not match manifest slug: %s", file.Name)
		}
	}
}

func writeManifest(t *testing.T, root, relative, name, version, description string, collections []string) {
	t.Helper()
	var text strings.Builder
	text.WriteString("---\nname: " + name + "\nversion: " + version + "\ndescription: " + description + "\ncollections:\n")
	for _, collection := range collections {
		text.WriteString("  - " + collection + "\n")
	}
	text.WriteString("---\n\nManifest.\n")
	writeSource(t, root, filepath.Join(relative, "_index.md"), text.String())
}

func writePage(t *testing.T, root, relative, title, pageID, body string) {
	t.Helper()
	writeSource(t, root, relative, "---\ntitle: "+title+"\npage_id: "+pageID+"\n---\n\n# "+title+"\n\n"+body+"\n")
}

func writeSource(t *testing.T, root, relative, content string) {
	t.Helper()
	name := filepath.Join(root, relative)
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readSource(t *testing.T, root, relative string) []byte {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(root, relative))
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func readCatalog(t *testing.T, root string) catalog {
	t.Helper()
	var result catalog
	if err := json.Unmarshal(readSource(t, root, "catalog.json"), &result); err != nil {
		t.Fatal(err)
	}
	return result
}

func artifactNamed(t *testing.T, artifacts []artifact, name string) []byte {
	t.Helper()
	for _, item := range artifacts {
		if item.name == name {
			return item.raw
		}
	}
	t.Fatalf("artifact %q not found", name)
	return nil
}

func compareDirectories(t *testing.T, left, right string) {
	t.Helper()
	leftEntries, err := os.ReadDir(left)
	if err != nil {
		t.Fatal(err)
	}
	rightEntries, err := os.ReadDir(right)
	if err != nil {
		t.Fatal(err)
	}
	if len(leftEntries) != len(rightEntries) {
		t.Fatalf("directory sizes differ: %d != %d", len(leftEntries), len(rightEntries))
	}
	for index, entry := range leftEntries {
		if entry.Name() != rightEntries[index].Name() {
			t.Fatalf("artifact names differ: %s != %s", entry.Name(), rightEntries[index].Name())
		}
		if !bytes.Equal(readSource(t, left, entry.Name()), readSource(t, right, entry.Name())) {
			t.Fatalf("artifact bytes differ: %s", entry.Name())
		}
	}
}
