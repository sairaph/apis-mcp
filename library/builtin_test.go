package library

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"testing"
)

func TestBuiltinCatalogContainsOnlyMarkdown(t *testing.T) {
	err := fs.WalkDir(os.DirFS("."), "builtin", func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() && !strings.EqualFold(path.Ext(name), ".md") {
			return fmt.Errorf("unexpected built-in catalog file %s", name)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
