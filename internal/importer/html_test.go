package importer

import (
	"net/url"
	"strings"
	"testing"
)

func TestHTMLRefreshRequiresImmediateNonEmptyTarget(t *testing.T) {
	base, err := url.Parse("https://docs.example.test/guide/")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name string
		html string
		want string
	}{
		{name: "delayed", html: `<meta http-equiv="refresh" content="5;url=../target/">`, want: "zero-second"},
		{name: "empty", html: `<meta http-equiv="refresh" content="0;url=">`, want: "empty URL"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := htmlRefreshRedirect([]byte(test.html), base); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %q error, got %v", test.want, err)
			}
		})
	}
	ordinary := `<script>const sample = '<meta http-equiv="refresh" content="0;url=../target/">'</script><!-- <meta http-equiv="refresh" content="0;url=../target/"> -->`
	if target, err := htmlRefreshRedirect([]byte(ordinary), base); err != nil || target != nil {
		t.Fatalf("ordinary text detected as refresh: %v, %v", target, err)
	}
}

func TestVitePressStartInventoryAliases(t *testing.T) {
	for _, test := range []struct {
		start     string
		inventory string
	}{
		{start: "https://docs.test/guide/start.html?source=test", inventory: "https://docs.test/guide/start"},
		{start: "https://docs.test/guide/index.html", inventory: "https://docs.test/guide/"},
		{start: "https://docs.test/guide/index", inventory: "https://docs.test/guide/"},
		{start: "https://docs.test/guide/", inventory: "https://docs.test/guide/index.html"},
	} {
		start, err := url.Parse(test.start)
		if err != nil {
			t.Fatal(err)
		}
		alias, err := vitepressStartInventoryAlias(start, "/", []string{test.inventory})
		if err != nil || alias != test.inventory {
			t.Errorf("alias %s: %q, %v", test.start, alias, err)
		}
	}
}

func TestDetectAstroStarlightGenerator(t *testing.T) {
	for _, test := range []struct {
		generator string
		want      string
	}{
		{generator: "Starlight v0.41.4", want: "astro-starlight"},
		{generator: "Starlight v1.0.0-beta.1+build.2", want: "astro-starlight"},
		{generator: "Astro v7.0.2"},
		{generator: "Starlight"},
	} {
		document, err := parseHTML([]byte(`<!doctype html><html><head><meta name="generator" content="` + test.generator + `"></head><body></body></html>`))
		if err != nil {
			t.Fatal(err)
		}
		if got := detectHTMLFramework(document); got != test.want {
			t.Errorf("generator %q: got %q, want %q", test.generator, got, test.want)
		}
	}
}

func TestStarlightSitemapRecordUsesHreflangIdentity(t *testing.T) {
	pageURL, err := url.Parse("https://docs.test/guide/")
	if err != nil {
		t.Fatal(err)
	}
	locale := starlightLocaleScope{language: "en", root: "/", excluded: []string{"/es/"}}
	alternates := []starlightSitemapAlternate{
		{Rel: "alternate", Language: "en", Href: "https://docs.test/guide/"},
		{Rel: "alternate", Language: "de", Href: "https://docs.test/de/guide/"},
		{Rel: "alternate", Language: "x-default", Href: "https://docs.test/guide/"},
	}
	selected, err := starlightSitemapRecordLocale(alternates, "https://docs.test/de/guide/", pageURL, "/", locale)
	if err != nil || selected {
		t.Fatalf("unadvertised locale selected: selected=%v err=%v", selected, err)
	}
	selected, err = starlightSitemapRecordLocale(alternates, "https://docs.test/guide/", pageURL, "/", locale)
	if err != nil || !selected {
		t.Fatalf("default locale not selected: selected=%v err=%v", selected, err)
	}
}
