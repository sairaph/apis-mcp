package importer

import (
	"context"
	"net/http"
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

func TestDetectMDBookRequiresGeneratedStructure(t *testing.T) {
	for _, test := range []struct {
		name string
		html string
		want string
	}{
		{name: "current", html: `<!-- Book generated using mdBook --><nav id="mdbook-sidebar"></nav><div id="mdbook-content"><main><h1>Book</h1></main></div>`, want: "mdbook"},
		{name: "legacy unsupported", html: `<!-- Book generated using mdBook --><nav id="sidebar"></nav><div id="content"><main><h1>Book</h1></main></div>`},
		{name: "comment only", html: `<!-- Book generated using mdBook --><main><h1>Not a book</h1></main>`},
		{name: "structure only", html: `<nav id="mdbook-sidebar"></nav><div id="mdbook-content"><main><h1>Not generated</h1></main></div>`},
	} {
		t.Run(test.name, func(t *testing.T) {
			document, err := parseHTML([]byte(`<!doctype html><html><body>` + test.html + `</body></html>`))
			if err != nil {
				t.Fatal(err)
			}
			if got := detectHTMLFramework(document); got != test.want {
				t.Fatalf("framework = %q, want %q", got, test.want)
			}
		})
	}
}

func TestMDBookStartInventoryAlias(t *testing.T) {
	first := "https://docs.test/book/title-page.html"
	for _, value := range []string{"https://docs.test/book/", "https://docs.test/book/index.html"} {
		pageURL, err := url.Parse(value)
		if err != nil {
			t.Fatal(err)
		}
		alias, err := mdbookStartInventoryAlias(pageURL, "/book/", first)
		if err != nil || alias != first {
			t.Errorf("alias %s: %q, %v", value, alias, err)
		}
	}
	nested, _ := url.Parse("https://docs.test/book/chapter.html")
	if _, err := mdbookStartInventoryAlias(nested, "/book/", first); err == nil {
		t.Fatal("nested chapter accepted as root alias")
	}
	queried, _ := url.Parse("https://docs.test/book/?variant=test")
	if _, err := mdbookStartInventoryAlias(queried, "/book/", first); err == nil {
		t.Fatal("query-bearing book root accepted as alias")
	}
}

func TestMDBookSiteRootUsesScopedTOCAssets(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><head></head><body><nav id="mdbook-sidebar"><noscript><iframe class="sidebar-iframe-outer" src="../toc.html"></iframe></noscript></nav><div id="mdbook-content"><main><script src="toc-decoy.js"></script></main></div></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	pageURL, _ := url.Parse("https://docs.test/book/chapter/start.html")
	root, tocURL, err := mdbookSiteRoot(document, pageURL)
	if err != nil || root != "/book/" || tocURL.String() != "https://docs.test/book/toc.html" {
		t.Fatalf("mdBook root: root=%q toc=%v err=%v", root, tocURL, err)
	}
}

func TestMDBookPageTitleUsesHeadOverride(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><head><title>Custom Chapter - Fixture Book</title></head><body><h1 class="menu-title">Fixture Book</h1></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	if title := mdbookPageTitle(document); title != "Custom Chapter" {
		t.Fatalf("mdBook title = %q", title)
	}
}

func TestHTMLHeadingTitleDoesNotPersistMarkdownEscapes(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><body><main><h1>API_v2 [beta]</h1></main></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	base, _ := url.Parse("https://docs.test/")
	title, markdown := htmlToMarkdown(firstHTMLTag(document, "main"), base, false)
	if title != "API_v2 [beta]" || markdown != `# API\_v2 \[beta\]` {
		t.Fatalf("title=%q markdown=%q", title, markdown)
	}
}

func TestImportHTMLRejectsURLCredentials(t *testing.T) {
	requests := 0
	options := Options{
		LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil },
		HTTPClient: &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
			requests++
			return docsifyTestResponse(request, http.StatusOK, `<html></html>`), nil
		})},
	}
	_, err := ImportHTML(context.Background(), "Private", "v1", "https://user:secret@docs.test/book/", options)
	if err == nil {
		t.Fatal("credential-bearing HTML URL accepted")
	}
	if _, err := DetectURL(context.Background(), "https://user:secret@docs.test/book/", options); err == nil {
		t.Fatal("credential-bearing detection URL accepted")
	}
	if requests != 0 {
		t.Fatalf("credential-bearing URL made %d requests", requests)
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
