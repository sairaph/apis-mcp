package budget

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCountO200kBase(t *testing.T) {
	tests := map[string]int{
		"hello world":  2,
		"hello  world": 3,
		"お誕生日おめでとう":    8,
	}
	for text, want := range tests {
		got, err := Count(text)
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Errorf("Count(%q) = %d, want %d", text, got, want)
		}
	}
}

func TestTruncateTokenAndUnicodeBoundaries(t *testing.T) {
	text := "hello お誕生日おめでとう world"
	prefix, tokens, truncated, err := Truncate(text, 4, 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	if !truncated || !utf8.ValidString(prefix) || tokens > 4 {
		t.Fatalf("prefix/tokens/truncated = %q / %d / %v", prefix, tokens, truncated)
	}
	actual, err := Count(prefix)
	if err != nil || actual != tokens {
		t.Fatalf("actual tokens = %d, %v; reported %d", actual, err, tokens)
	}

	prefix, tokens, truncated, err = Truncate("🙂🙂", 100, 5)
	if err != nil || prefix != "🙂" || tokens == 0 || !truncated {
		t.Fatalf("byte-bounded Unicode prefix = %q / %d / %v / %v", prefix, tokens, truncated, err)
	}

	prefix, tokens, truncated, err = Truncate("𓀀", 1, 1<<20)
	if err != nil || !utf8.ValidString(prefix) || tokens > 1 || !truncated {
		t.Fatalf("split-token Unicode prefix = %q / %d / %v / %v", prefix, tokens, truncated, err)
	}
}

func TestPaginateExactBoundaryAndOversizedFirstRecord(t *testing.T) {
	render := func(records []string) (string, error) { return strings.Join(records, ""), nil }
	records := []string{"hello", " world", "supercalifragilistic"}

	first, pages, err := Paginate(records, 1, 2, render)
	if err != nil || pages != 2 || len(first) != 2 {
		t.Fatalf("exact-boundary page = %v / %d / %v", first, pages, err)
	}
	second, pages, err := Paginate(records, 2, 2, render)
	if err != nil || pages != 2 || len(second) != 1 || second[0] != "supercalifragilistic" {
		t.Fatalf("oversized page = %v / %d / %v", second, pages, err)
	}
}
