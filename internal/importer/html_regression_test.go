package importer

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestParseHTMLAcceptsDocumentJustOverOldNodeLimit(t *testing.T) {
	raw := []byte(`<!doctype html><html><body><main>` + strings.Repeat(`<span>x</span>`, 25_001) + `</main></body></html>`)
	if _, err := parseHTML(raw); err != nil {
		t.Fatalf("parse dense bounded HTML: %v", err)
	}
}

func TestDeterministicHTMLRouteAlias(t *testing.T) {
	for _, pair := range [][2]string{
		{"https://docs.test/CONTRIBUTING", "https://docs.test/contributing"},
		{"https://docs.test/api/ErrorTypes.html", "https://docs.test/api/errortypes"},
		{"https://docs.test/guide/index.html", "https://docs.test/guide/"},
	} {
		first, _ := url.Parse(pair[0])
		second, _ := url.Parse(pair[1])
		if !deterministicHTMLRouteAlias(first, second) {
			t.Errorf("expected deterministic alias %s -> %s", first, second)
		}
	}
	first, _ := url.Parse("https://docs.test/advertised.html")
	second, _ := url.Parse("https://docs.test/unlisted.html")
	if deterministicHTMLRouteAlias(first, second) {
		t.Fatal("unrelated in-root route accepted as an alias")
	}
}

func TestHTMLSitemapInventorySupportsIndexAssetsAndLocaleAlternates(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<sitemapindex><sitemap><loc>%s/sitemap-0.xml</loc></sitemap></sitemapindex>`, server.URL)
		case "/sitemap-0.xml":
			fmt.Fprintf(writer, `<urlset xmlns:xhtml="http://www.w3.org/1999/xhtml"><url><loc>%[1]s/docs/page</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US/docs/page"/><xhtml:link rel="alternate" hreflang="ja" href="%[1]s/ja-JP/docs/page"/></url><url><loc>%[1]s/apple-icon.png</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	pageURL, _ := url.Parse(server.URL + "/en-US/docs/start")
	reader := &sourceReader{client: server.Client(), perSource: 1 << 20, total: 4 << 20}
	var exclusions []string
	inventory, err := htmlSitemapInventory(context.Background(), "Nextra", pageURL, "/", "/en-US/", reader, func(source string) {
		exclusions = append(exclusions, source)
	})
	if err != nil {
		t.Fatal(err)
	}
	want := server.URL + "/en-US/docs/page"
	if len(inventory) != 1 || inventory[0] != want || len(exclusions) != 1 || exclusions[0] != server.URL+"/apple-icon.png" {
		t.Fatalf("inventory = %v, exclusions = %v, want [%s] and one asset", inventory, exclusions, want)
	}
}

func TestHTMLSitemapInventoryRejectsNonPagesForAuthoritativeFrameworks(t *testing.T) {
	for _, framework := range []string{"MkDocs", "MkDocs Material", "VitePress"} {
		t.Run(framework, func(t *testing.T) {
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				if request.URL.Path != "/docs/sitemap.xml" {
					http.NotFound(writer, request)
					return
				}
				fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/docs/start/</loc></url><url><loc>%[1]s/docs/logo.png</loc></url></urlset>`, server.URL)
			}))
			defer server.Close()
			pageURL, _ := url.Parse(server.URL + "/docs/start/")
			reader := &sourceReader{client: server.Client(), perSource: 1 << 20, total: 2 << 20}
			if _, err := htmlSitemapInventory(context.Background(), framework, pageURL, "/docs/", "/docs/", reader, nil); err == nil || !strings.Contains(err.Error(), "unsupported non-page URL") {
				t.Fatalf("%s non-page sitemap record was not rejected: %v", framework, err)
			}
		})
	}
}

func TestHTMLSitemapInventoryAccountsSixNextraAssetsAndRejectsInvalidRecords(t *testing.T) {
	assets := []string{"icon.png", "diagram.svg", "robots.txt", "preview.jpg", "social.png", "flow.svg"}
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/sitemap.xml" {
			http.NotFound(writer, request)
			return
		}
		fmt.Fprintf(writer, `<urlset><url><loc>%s/docs/start/</loc></url>`, server.URL)
		for _, asset := range assets {
			fmt.Fprintf(writer, `<url><loc>%s/%s</loc></url>`, server.URL, asset)
		}
		fmt.Fprint(writer, `</urlset>`)
	}))
	defer server.Close()
	pageURL, _ := url.Parse(server.URL + "/docs/start/")
	reader := &sourceReader{client: server.Client(), perSource: 1 << 20, total: 2 << 20}
	var exclusions []string
	inventory, err := htmlSitemapInventory(context.Background(), "Nextra", pageURL, "/", "/docs/", reader, func(source string) {
		exclusions = append(exclusions, source)
	})
	if err != nil || len(inventory) != 1 || len(exclusions) != 6 {
		t.Fatalf("Nextra sitemap accounting = pages %v, exclusions %v, err %v", inventory, exclusions, err)
	}

	for _, test := range []struct {
		name string
		xml  string
	}{
		{name: "duplicate", xml: `<urlset><url><loc>%[1]s/docs/start/</loc></url><url><loc>%[1]s/icon.png</loc></url><url><loc>%[1]s/icon.png</loc></url></urlset>`},
		{name: "unknown extension", xml: `<urlset><url><loc>%[1]s/docs/start/</loc></url><url><loc>%[1]s/manual.pdf</loc></url></urlset>`},
		{name: "asset alternate", xml: `<urlset><url><loc>%[1]s/docs/start/</loc></url><url><loc>%[1]s/icon.png</loc><link rel="alternate" hreflang="en" href="%[1]s/icon.png"/></url></urlset>`},
	} {
		t.Run(test.name, func(t *testing.T) {
			var fixture *httptest.Server
			fixture = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(writer, test.xml, fixture.URL)
			}))
			defer fixture.Close()
			fixtureURL, _ := url.Parse(fixture.URL + "/docs/start/")
			fixtureReader := &sourceReader{client: fixture.Client(), perSource: 1 << 20, total: 2 << 20}
			if _, inventoryErr := htmlSitemapInventory(context.Background(), "Nextra", fixtureURL, "/", "/docs/", fixtureReader, func(string) {}); inventoryErr == nil {
				t.Fatal("invalid Nextra non-page record was accepted")
			}
		})
	}
}

func TestImportHTMLNextraReportsCompletePageAndAssetAccounting(t *testing.T) {
	assets := []string{"icon.png", "diagram.svg", "robots.txt", "preview.jpg", "social.png", "flow.svg"}
	assetRequests := 0
	var progress []Progress
	var server *httptest.Server
	page := func(title string) string {
		return `<!doctype html><html><head><script src="/_next/static/chunks/app.js"></script></head><body><a id="nextra-skip-nav"></a><aside class="nextra-sidebar"><a href="/docs/start/">Start</a><a href="/docs/second/">Second</a></aside><main data-pagefind-body="true"><h1>` + title + `</h1><p>Static Nextra content.</p></main></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start/":
			fmt.Fprint(writer, page("Start"))
		case "/docs/second/":
			fmt.Fprint(writer, page("Second"))
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/docs/start/</loc></url><url><loc>%[1]s/docs/second/</loc></url>`, server.URL)
			for _, asset := range assets {
				fmt.Fprintf(writer, `<url><loc>%s/%s</loc></url>`, server.URL, asset)
			}
			fmt.Fprint(writer, `</urlset>`)
		default:
			assetRequests++
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	options := htmlRegressionOptions(t, server.Client())
	options.Progress = func(update Progress) { progress = append(progress, update) }
	result, err := ImportHTML(context.Background(), "Nextra Assets", "v1", server.URL+"/docs/start/", options)
	if err != nil || result.Pages != 2 || result.Truncated || assetRequests != 0 {
		t.Fatalf("Nextra asset import = %+v, asset requests=%d, err=%v", result, assetRequests, err)
	}
	want := "accounted 8 of 8 sitemap records: 2 fetched pages, 6 validated non-page exclusions"
	found := false
	for _, update := range progress {
		found = found || update.Stage == "inventory" && update.Message == want && update.Pages == 2
	}
	if !found {
		t.Fatalf("missing Nextra sitemap accounting %q in %+v", want, progress)
	}
}

func TestInspectElementPlusVitePressHome(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><head><meta name="generator" content="VitePress v1.6.4"></head><body><div id="app"><div class="App"><header class="navbar">Chrome</header><main id="page-content" class="page-content"><div class="hero-content"><h1>Element Plus</h1><p>Static home content.</p></div><footer>Chrome</footer></main></div></div></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	view := inspectHTMLDocument(document)
	if view.framework != "vitepress" || view.profileError != nil || view.content == nil || view.content.tag != "main" || view.content.attrs["id"] != "page-content" {
		t.Fatalf("Element Plus view = %+v", view)
	}
}

func TestStarlightLocaleSupportsScopedMonolingualSite(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html lang="en"><head><link rel="canonical" href="https://docs.test/docs/start/"></head><body><div id="starlight__sidebar"><a href="/docs/start/">Start</a><a href="/docs/second/">Second</a><a href="/changelog/">Changelog</a></div></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	pageURL, _ := url.Parse("https://docs.test/docs/start/")
	sitemapURL, _ := url.Parse("https://docs.test/sitemap-index.xml")
	locale, err := starlightLocale(document, pageURL, sitemapURL)
	if err != nil || locale.language != "en" || locale.root != "/docs/" || locale.multilingual {
		t.Fatalf("monolingual locale = %+v, %v", locale, err)
	}
	selected, err := starlightSitemapRecordLocale(nil, "https://docs.test/docs/second/", pageURL, "/", locale)
	if err != nil || !selected {
		t.Fatalf("documentation route not selected: %v, %v", selected, err)
	}
	selected, err = starlightSitemapRecordLocale(nil, "https://docs.test/pricing/", pageURL, "/", locale)
	if err != nil || selected {
		t.Fatalf("global marketing route selected: %v, %v", selected, err)
	}
}

func TestStarlightMonolingualScopeCanSpanDocumentSections(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html lang="en"><head><link rel="canonical" href="https://docs.test/docs/"></head><body><div id="starlight__sidebar"><a href="/docs/">Overview</a><a href="/supporters/overview/">Supporters</a><a href="/publishers/overview/">Publishers</a></div></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	pageURL, _ := url.Parse("https://docs.test/docs/")
	sitemapURL, _ := url.Parse("https://docs.test/sitemap-index.xml")
	locale, err := starlightLocale(document, pageURL, sitemapURL)
	if err != nil || locale.root != "/" {
		t.Fatalf("cross-section monolingual locale = %+v, %v", locale, err)
	}
}

func TestStarlightMonolingualSitemapExcludesNonPagesAndOutsideScope(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/sitemap-index.xml":
			fmt.Fprintf(writer, `<sitemapindex><sitemap><loc>%s/sitemap-0.xml</loc></sitemap></sitemapindex>`, server.URL)
		case "/sitemap-0.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/docs/start/</loc></url><url><loc>%[1]s/docs/llms.txt</loc></url><url><loc>%[1]s/pricing/</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	pageURL, _ := url.Parse(server.URL + "/docs/start/")
	sitemapURL, _ := url.Parse(server.URL + "/sitemap-index.xml")
	reader := &sourceReader{client: server.Client(), perSource: 1 << 20, total: 4 << 20}
	locale := starlightLocaleScope{language: "en", root: "/docs/"}
	inventory, err := starlightSitemapInventory(context.Background(), pageURL, sitemapURL, "/", locale, reader)
	if err != nil || len(inventory) != 2 || inventory[0] != server.URL+"/docs/start/" || inventory[1] != server.URL+"/pricing/" {
		t.Fatalf("monolingual Starlight inventory = %v, %v", inventory, err)
	}
}

func TestMDBookTOCInventoryAcceptsRootAndCleanRoutes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/toc.html" {
			http.NotFound(writer, request)
			return
		}
		fmt.Fprint(writer, `<!doctype html><html><body class="sidebar-iframe-inner"><ol class="chapter"><li><a href="/">Home</a></li><li><a href="/book/">Book</a></li><li><a href="/book/chapter">Chapter</a></li></ol></body></html>`)
	}))
	defer server.Close()
	pageURL, _ := url.Parse(server.URL + "/book/")
	tocURL, _ := url.Parse(server.URL + "/toc.html")
	reader := &sourceReader{client: server.Client(), perSource: 1 << 20, total: 2 << 20}
	inventory, first, err := mdbookTOCInventory(context.Background(), pageURL, "/", tocURL, reader)
	if err != nil || len(inventory) != 3 || first != server.URL+"/" {
		t.Fatalf("clean mdBook inventory = %v, first=%q, err=%v", inventory, first, err)
	}
}

func TestImportHTMLMDBookGeneratesEveryCleanTOCChapter(t *testing.T) {
	var server *httptest.Server
	page := func(title string) string {
		return `<!doctype html><html><head><!-- Book generated using mdBook --><title>` + title + ` - Book</title><script src="/toc.js"></script></head><body><nav id="mdbook-sidebar"></nav><h1 class="menu-title">Book</h1><div id="mdbook-content"><main><h1>` + title + `</h1><p>Chapter content.</p></main></div></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			fmt.Fprint(writer, page("Home"))
		case "/book/":
			fmt.Fprint(writer, page("Book"))
		case "/book/chapter":
			fmt.Fprint(writer, page("Chapter"))
		case "/toc.html":
			fmt.Fprint(writer, `<!doctype html><html><body class="sidebar-iframe-inner"><ol class="chapter"><li><a href="/">Home</a></li><li><a href="/book/">Book</a></li><li><a href="/book/chapter">Chapter</a></li></ol></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	result, err := ImportHTML(context.Background(), "Aya", "v1", server.URL+"/book/", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 3 || result.Truncated {
		t.Fatalf("clean mdBook import = %+v, %v", result, err)
	}
}

func TestMDBookEmptyIndexAliasesUniqueChildButNotOrdinaryEmptyChapter(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><head><title>Group - Book</title></head><body><h1 class="menu-title">Book</h1><a rel="next" href="child.html">Next</a><a rel="next prefetch" href="child.html">Next again</a></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	pageURL, _ := url.Parse("https://docs.test/group/index.html")
	inventory := map[string]bool{"https://docs.test/group/index.html": true, "https://docs.test/group/child.html": true}
	alias, err := mdbookEmptyChapterAlias(document, pageURL, inventory)
	if err != nil || alias != "https://docs.test/group/child.html" {
		t.Fatalf("empty group alias = %q, %v", alias, err)
	}
	ordinaryURL, _ := url.Parse("https://docs.test/empty.html")
	if _, err := mdbookEmptyChapterAlias(document, ordinaryURL, inventory); err == nil {
		t.Fatal("ordinary empty chapter accepted as an alias")
	}
}

func TestImportHTMLMDBookTreatsEmptyIndexGroupAsBoundedAlias(t *testing.T) {
	var server *httptest.Server
	page := func(title, body, next string) string {
		navigation := ""
		if next != "" {
			navigation = `<nav><a rel="next" href="` + next + `">Next</a></nav>`
		}
		return `<!doctype html><html><head><!-- Book generated using mdBook --><title>` + title + ` - Book</title><script src="/book/toc.js"></script></head><body><nav id="mdbook-sidebar"></nav><h1 class="menu-title">Book</h1><div id="mdbook-content"><main>` + body + `</main>` + navigation + `</div></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/book/start.html":
			fmt.Fprint(writer, page("Start", `<h1>Start</h1><p>Start content.</p>`, ""))
		case "/book/group/index.html":
			fmt.Fprint(writer, page("Group", "", "child.html"))
		case "/book/group/child.html":
			fmt.Fprint(writer, page("Child", `<h1>Child</h1><p>Child content.</p>`, ""))
		case "/book/toc.html":
			fmt.Fprint(writer, `<!doctype html><html><body class="sidebar-iframe-inner"><ol class="chapter"><li><a href="start.html">Start</a></li><li><a href="group/index.html">Group</a><ol><li><a href="group/child.html">Child</a></li></ol></li></ol></body></html>`)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	result, err := ImportHTML(context.Background(), "Servo", "v1", server.URL+"/book/start.html", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 2 || result.Truncated {
		t.Fatalf("empty mdBook group import = %+v, %v", result, err)
	}
}

func TestImportHTMLMkDocsAcceptsDirectoryIndexInventoryAlias(t *testing.T) {
	var server *httptest.Server
	page := `<!doctype html><html><head><meta name="generator" content="mkdocs-1.6.1"><script>var base_url = ".";</script></head><body><div id="navbar-collapse"><a href="index.html">Home</a></div><main role="main"><h1>LORIS</h1><p>Classic MkDocs content.</p></main></body></html>`
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/en/latest/":
			fmt.Fprint(writer, page)
		case "/en/latest/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%s/en/latest/index.html</loc></url></urlset>`, server.URL)
		case "/en/latest/index.html":
			http.Error(writer, "directory alias fetched twice", http.StatusInternalServerError)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	result, err := ImportHTML(context.Background(), "LORIS", "v1", server.URL+"/en/latest/", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 1 || result.Truncated {
		t.Fatalf("MkDocs index alias import = %+v, %v", result, err)
	}
}

func TestImportHTMLVitePressAccountsForCaseNormalizedRedirectAlias(t *testing.T) {
	var server *httptest.Server
	page := func(title string) string {
		return `<!doctype html><html><head><meta name="generator" content="VitePress v1.6.4"><link rel="stylesheet" href="/vp-icons.css"></head><body><div class="VPDoc"><div class="vp-doc"><h1>` + title + `</h1><p>Static content.</p></div></div></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			fmt.Fprint(writer, page("Home"))
		case "/API/ErrorTypes.html":
			http.Redirect(writer, request, "/api/errortypes", http.StatusMovedPermanently)
		case "/api/errortypes":
			fmt.Fprint(writer, page("Error Types"))
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/</loc></url><url><loc>%[1]s/API/ErrorTypes.html</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	result, err := ImportHTML(context.Background(), "Case Redirect", "v1", server.URL+"/", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 2 || result.Truncated {
		t.Fatalf("case redirect import = %+v, %v", result, err)
	}
}

func TestImportHTMLRejectsSharedCaseNormalizedRedirectAlias(t *testing.T) {
	var server *httptest.Server
	page := func(title string) string {
		return `<!doctype html><html><head><meta name="generator" content="VitePress v1.6.4"><link rel="stylesheet" href="/vp-icons.css"></head><body><div class="VPDoc"><div class="vp-doc"><h1>` + title + `</h1><p>Static content.</p></div></div></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			fmt.Fprint(writer, page("Home"))
		case "/API/ErrorTypes.html", "/api/ERRORTYPES.html":
			http.Redirect(writer, request, "/api/errortypes", http.StatusMovedPermanently)
		case "/api/errortypes":
			fmt.Fprint(writer, page("Error Types"))
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/</loc></url><url><loc>%[1]s/API/ErrorTypes.html</loc></url><url><loc>%[1]s/api/ERRORTYPES.html</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	_, err := ImportHTML(context.Background(), "Ambiguous Redirect", "v1", server.URL+"/", htmlRegressionOptions(t, server.Client()))
	if err == nil || !strings.Contains(err.Error(), "shared by multiple inventory entries") {
		t.Fatalf("expected shared redirect alias rejection, got %v", err)
	}
}

func TestImportHTMLNextraFollowsRelativeLocaleRedirectAndSelectsAlternates(t *testing.T) {
	var server *httptest.Server
	page := func(title string) string {
		return `<!doctype html><html><head><script src="/_next/static/chunks/app.js"></script></head><body><a id="nextra-skip-nav"></a><aside class="nextra-sidebar"><a href="/en-US">Home</a><a href="/en-US/page">Page</a></aside><main data-pagefind-body="true"><h1>` + title + `</h1><p>Localized static content.</p></main></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			writer.Header().Set("Location", "/en-US")
			writer.WriteHeader(http.StatusTemporaryRedirect)
		case "/en-US":
			fmt.Fprint(writer, page("Home"))
		case "/en-US/page":
			fmt.Fprint(writer, page("Page"))
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset xmlns:xhtml="http://www.w3.org/1999/xhtml"><url><loc>%[1]s</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US"/><xhtml:link rel="alternate" hreflang="ja" href="%[1]s/ja-JP"/></url><url><loc>%[1]s/page</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US/page"/><xhtml:link rel="alternate" hreflang="ja" href="%[1]s/ja-JP/page"/></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()
	result, err := ImportHTML(context.Background(), "imgix", "v1", server.URL, htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 2 || result.Source != server.URL+"/en-US" || result.Truncated {
		t.Fatalf("locale redirect import = %+v, %v", result, err)
	}
}

func TestImportHTMLNextraEstablishesNestedScopeAfterInitialLocaleRedirect(t *testing.T) {
	var server *httptest.Server
	page := func(title string) string {
		return `<!doctype html><html><head><script src="/_next/static/chunks/app.js"></script></head><body><a id="nextra-skip-nav"></a><aside class="nextra-sidebar"><a href="/en-US/apis/rendering/overview">Overview</a><a href="/en-US/apis/rendering/format">Format</a></aside><main data-pagefind-body="true"><h1>` + title + `</h1><p>Localized rendering API content.</p></main></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/apis/rendering/overview":
			writer.Header().Set("Location", "/en-US/apis/rendering/overview")
			writer.WriteHeader(http.StatusTemporaryRedirect)
		case "/en-US/apis/rendering/overview":
			fmt.Fprint(writer, page("Rendering overview"))
		case "/en-US/apis/rendering/format":
			fmt.Fprint(writer, page("Rendering format"))
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset xmlns:xhtml="http://www.w3.org/1999/xhtml"><url><loc>%[1]s/apis/rendering/overview</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US/apis/rendering/overview"/><xhtml:link rel="alternate" hreflang="ja" href="%[1]s/ja-JP/apis/rendering/overview"/></url><url><loc>%[1]s/apis/rendering/format</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US/apis/rendering/format"/><xhtml:link rel="alternate" hreflang="ja" href="%[1]s/ja-JP/apis/rendering/format"/></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	result, err := ImportHTML(context.Background(), "imgix nested", "v1", server.URL+"/apis/rendering/overview", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 2 || result.Source != server.URL+"/en-US/apis/rendering/overview" || result.Truncated {
		t.Fatalf("nested locale redirect import = %+v, %v", result, err)
	}
}

func TestInitialHTMLLocaleRedirectIsExactAndUnambiguous(t *testing.T) {
	start, _ := url.Parse("https://docs.test/apis/rendering/overview")
	accepted, _ := url.Parse("https://docs.test/en-US/apis/rendering/overview")
	if !initialHTMLLocaleRedirect(start, accepted) {
		t.Fatal("exact nested locale-prefix redirect rejected")
	}
	for _, target := range []string{
		"https://docs.test/marketing/apis/rendering/overview",
		"https://docs.test/en-us/apis/rendering/overview",
		"https://docs.test/en-US/apis/rendering/other",
		"https://docs.test/en-US/apis/rendering/overview?preview=1",
		"https://docs.test/en-US/apis/rendering/overview#section",
		"https://user@docs.test/en-US/apis/rendering/overview",
		"https://other.test/en-US/apis/rendering/overview",
	} {
		redirected, _ := url.Parse(target)
		if initialHTMLLocaleRedirect(start, redirected) {
			t.Errorf("ambiguous locale redirect accepted: %s", target)
		}
	}
}

func TestImportHTMLNextraRejectsLocaleAlternateCollisionAfterInitialRedirect(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/apis/rendering/overview":
			http.Redirect(writer, request, "/en-US/apis/rendering/overview", http.StatusTemporaryRedirect)
		case "/en-US/apis/rendering/overview":
			fmt.Fprint(writer, `<!doctype html><html><head><script src="/_next/static/chunks/app.js"></script></head><body><a id="nextra-skip-nav"></a><aside class="nextra-sidebar"><a href="/en-US/apis/rendering/overview">Overview</a><a href="/en-US/apis/rendering/format">Format</a></aside><main data-pagefind-body="true"><h1>Overview</h1><p>Content.</p></main></body></html>`)
		case "/sitemap.xml":
			fmt.Fprintf(writer, `<urlset xmlns:xhtml="http://www.w3.org/1999/xhtml"><url><loc>%[1]s/apis/rendering/overview</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US/apis/rendering/overview"/></url><url><loc>%[1]s/apis/rendering/alias</loc><xhtml:link rel="alternate" hreflang="en" href="%[1]s/en-US/apis/rendering/overview"/></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	_, err := ImportHTML(context.Background(), "imgix collision", "v1", server.URL+"/apis/rendering/overview", htmlRegressionOptions(t, server.Client()))
	if err == nil || !strings.Contains(err.Error(), "repeats selected URL") {
		t.Fatalf("expected locale alias collision rejection, got %v", err)
	}
}

func TestImportHTMLMonolingualStarlightAccountsForMixedLayouts(t *testing.T) {
	var server *httptest.Server
	requests := make(map[string]int)
	var progress []Progress
	starlightPage := func(route, title string) string {
		return `<!doctype html><html lang="en"><head><meta name="generator" content="Astro v6.3.8"><meta name="generator" content="Starlight v0.39.2"><link rel="canonical" href="` + server.URL + route + `"><link rel="sitemap" href="/sitemap-index.xml"></head><body><div id="starlight__sidebar"><a href="/docs/">Docs</a><a href="/developers/guide/">Developer guide</a></div><main data-pagefind-body><h1>` + title + `</h1><div class="sl-markdown-content"><p>Static Starlight content.</p></div></main></body></html>`
	}
	landingPage := func(route, title string) string {
		return `<!doctype html><html lang="en"><head><meta name="generator" content="Astro v6.3.8"><link rel="canonical" href="` + server.URL + route + `"></head><body><main><h1>` + title + `</h1><p>Static landing content.</p></main></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests[request.URL.Path]++
		switch request.URL.Path {
		case "/docs/":
			fmt.Fprint(writer, starlightPage("/docs/", "Docs"))
		case "/developers/guide/":
			fmt.Fprint(writer, starlightPage("/developers/guide/", "Developer guide"))
		case "/":
			fmt.Fprint(writer, landingPage("/", "Home"))
		case "/developers/":
			fmt.Fprint(writer, landingPage("/developers/", "Developers"))
		case "/sitemap-index.xml":
			fmt.Fprintf(writer, `<sitemapindex><sitemap><loc>%s/sitemap-0.xml</loc></sitemap></sitemapindex>`, server.URL)
		case "/sitemap-0.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/</loc></url><url><loc>%[1]s/developers/</loc></url><url><loc>%[1]s/developers/guide/</loc></url><url><loc>%[1]s/docs/</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	options := htmlRegressionOptions(t, server.Client())
	options.Progress = func(update Progress) { progress = append(progress, update) }
	result, err := ImportHTML(context.Background(), "Mixed Starlight", "v1", server.URL+"/docs/", options)
	if err != nil || result.Pages != 2 || result.Truncated {
		t.Fatalf("mixed-layout Starlight import = %+v, %v", result, err)
	}
	for _, route := range []string{"/", "/developers/", "/developers/guide/", "/docs/"} {
		if requests[route] != 1 {
			t.Errorf("page requests for %s = %d, want 1", route, requests[route])
		}
	}
	wantProgress := "fetched 4 of 4 sitemap pages: 2 generated, 2 non-framework exclusions"
	foundProgress := false
	for _, update := range progress {
		if update.Stage == "inventory" && update.Message == wantProgress && update.Pages == 2 {
			foundProgress = true
		}
	}
	if !foundProgress {
		t.Fatalf("missing final inventory accounting %q in %+v", wantProgress, progress)
	}
}

func TestImportHTMLMonolingualStarlightRejectsInvalidMixedLayoutExclusions(t *testing.T) {
	tests := []struct {
		name    string
		route   string
		landing func(string) string
	}{
		{name: "missing canonical", route: "/landing/", landing: func(string) string {
			return `<!doctype html><html><body><main><h1>Landing</h1><p>Content.</p></main></body></html>`
		}},
		{name: "empty static body", route: "/landing/", landing: func(base string) string {
			return `<!doctype html><html><head><link rel="canonical" href="` + base + `/landing/"></head><body></body></html>`
		}},
		{name: "residual Starlight fingerprint", route: "/landing/", landing: func(base string) string {
			return `<!doctype html><html><head><meta name="generator" content="Starlight preview"><link rel="canonical" href="` + base + `/landing/"></head><body><main><h1>Landing</h1><p>Content.</p></main></body></html>`
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var server *httptest.Server
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				switch request.URL.Path {
				case "/docs/":
					fmt.Fprintf(writer, `<!doctype html><html lang="en"><head><meta name="generator" content="Starlight v0.39.2"><link rel="canonical" href="%[1]s/docs/"><link rel="sitemap" href="/sitemap-index.xml"></head><body><div id="starlight__sidebar"><a href="/docs/">Docs</a><a href="/landing/">Landing</a></div><main data-pagefind-body><h1>Docs</h1><div class="sl-markdown-content"><p>Content.</p></div></main></body></html>`, server.URL)
				case test.route:
					fmt.Fprint(writer, test.landing(server.URL))
				case "/sitemap-index.xml":
					fmt.Fprintf(writer, `<sitemapindex><sitemap><loc>%s/sitemap-0.xml</loc></sitemap></sitemapindex>`, server.URL)
				case "/sitemap-0.xml":
					fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/docs/</loc></url><url><loc>%[1]s/landing/</loc></url></urlset>`, server.URL)
				default:
					http.NotFound(writer, request)
				}
			}))
			defer server.Close()

			_, err := ImportHTML(context.Background(), "Invalid Mixed Starlight", "v1", server.URL+"/docs/", htmlRegressionOptions(t, server.Client()))
			if err == nil || !strings.Contains(err.Error(), "invalid non-framework sitemap page") {
				t.Fatalf("expected invalid exclusion rejection, got %v", err)
			}
		})
	}
}

func TestImportHTMLMonolingualStarlightRejectsRedirectingMixedLayoutExclusion(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/":
			fmt.Fprintf(writer, `<!doctype html><html lang="en"><head><meta name="generator" content="Starlight v0.39.2"><link rel="canonical" href="%[1]s/docs/"><link rel="sitemap" href="/sitemap-index.xml"></head><body><div id="starlight__sidebar"><a href="/docs/">Docs</a><a href="/landing/index.html">Landing</a></div><main data-pagefind-body><h1>Docs</h1><div class="sl-markdown-content"><p>Content.</p></div></main></body></html>`, server.URL)
		case "/landing/index.html":
			http.Redirect(writer, request, "/landing/", http.StatusMovedPermanently)
		case "/landing/":
			fmt.Fprintf(writer, `<!doctype html><html><head><link rel="canonical" href="%s/landing/"></head><body><main><h1>Landing</h1><p>Static content.</p></main></body></html>`, server.URL)
		case "/sitemap-index.xml":
			fmt.Fprintf(writer, `<sitemapindex><sitemap><loc>%s/sitemap-0.xml</loc></sitemap></sitemapindex>`, server.URL)
		case "/sitemap-0.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/docs/</loc></url><url><loc>%[1]s/landing/index.html</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	_, err := ImportHTML(context.Background(), "Redirecting Mixed Starlight", "v1", server.URL+"/docs/", htmlRegressionOptions(t, server.Client()))
	if err == nil || !strings.Contains(err.Error(), "non-framework sitemap page redirected") {
		t.Fatalf("expected redirecting exclusion rejection, got %v", err)
	}
}

func TestImportHTMLMultilingualStarlightDoesNotRelaxMixedLayoutInventory(t *testing.T) {
	var server *httptest.Server
	landingRequests := 0
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/en/docs/":
			fmt.Fprintf(writer, `<!doctype html><html lang="en"><head><meta name="generator" content="Starlight v0.39.2"><link rel="canonical" href="%[1]s/en/docs/"><link rel="alternate" hreflang="en" href="%[1]s/en/docs/"><link rel="alternate" hreflang="es" href="%[1]s/es/docs/"><link rel="sitemap" href="/sitemap-index.xml"></head><body><div id="starlight__sidebar"><a href="/en/docs/">Docs</a></div><main data-pagefind-body><h1>Docs</h1><div class="sl-markdown-content"><p>Content.</p></div></main></body></html>`, server.URL)
		case "/landing/":
			landingRequests++
			fmt.Fprintf(writer, `<!doctype html><html><head><link rel="canonical" href="%s/landing/"></head><body><main><h1>Landing</h1><p>Content.</p></main></body></html>`, server.URL)
		case "/sitemap-index.xml":
			fmt.Fprintf(writer, `<sitemapindex><sitemap><loc>%s/sitemap-0.xml</loc></sitemap></sitemapindex>`, server.URL)
		case "/sitemap-0.xml":
			fmt.Fprintf(writer, `<urlset><url><loc>%[1]s/en/docs/</loc><link rel="alternate" hreflang="en" href="%[1]s/en/docs/"/><link rel="alternate" hreflang="es" href="%[1]s/es/docs/"/></url><url><loc>%[1]s/landing/</loc></url></urlset>`, server.URL)
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	_, err := ImportHTML(context.Background(), "Multilingual Mixed Starlight", "v1", server.URL+"/en/docs/", htmlRegressionOptions(t, server.Client()))
	if err == nil || !strings.Contains(err.Error(), "multilingual URL record has no alternates") || landingRequests != 0 {
		t.Fatalf("multilingual mixed-layout inventory = requests %d, err %v", landingRequests, err)
	}
}

func TestImportHTMLDocusaurusDerivesSidebarRootAndDeduplicatesLeafSlashAlias(t *testing.T) {
	for _, startPath := range []string{"/docs/api", "/docs/api/"} {
		t.Run(strings.TrimPrefix(startPath, "/docs/"), func(t *testing.T) {
			routes := []string{"/docs/", "/docs/api"}
			for index := 1; index <= 23; index++ {
				routes = append(routes, fmt.Sprintf("/docs/page-%02d", index))
			}
			requests := make(map[string]int)
			var progress []Progress
			var server *httptest.Server
			page := func(route string) string {
				canonical := route
				if canonical == "/docs/api/" {
					canonical = "/docs/api"
				}
				var sidebar strings.Builder
				for _, linked := range routes {
					fmt.Fprintf(&sidebar, `<li><a href="%s">Page</a></li>`, linked)
				}
				sidebar.WriteString(`<li><a href="/docs/api/">API slash alias</a></li>`)
				return `<!doctype html><html class="plugin-docs docs-doc-page"><head><meta name="generator" content="Docusaurus v3.10.2"><link rel="canonical" href="` + server.URL + canonical + `"></head><body><nav><a href="/blog/">Blog</a></nav><aside><ul class="theme-doc-sidebar-menu">` + sidebar.String() + `</ul></aside><div class="theme-doc-markdown markdown"><header><h1>` + route + `</h1></header><p>Static Docusaurus content.</p></div></body></html>`
			}
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				requests[request.URL.Path]++
				if request.URL.Path == "/docs/api/" || request.URL.Path == "/docs/api" || containsString(routes, request.URL.Path) {
					fmt.Fprint(writer, page(request.URL.Path))
					return
				}
				http.NotFound(writer, request)
			}))
			defer server.Close()

			options := htmlRegressionOptions(t, server.Client())
			options.Progress = func(update Progress) { progress = append(progress, update) }
			result, err := ImportHTML(context.Background(), "Prettier Docusaurus", "v1", server.URL+startPath, options)
			if err != nil || result.Pages != 25 || result.Truncated {
				t.Fatalf("Docusaurus leaf import = %+v, %v", result, err)
			}
			aliasPath := "/docs/api/"
			if startPath == aliasPath {
				aliasPath = "/docs/api"
			}
			if requests[aliasPath] != 0 {
				t.Fatalf("slash alias %s fetched %d times", aliasPath, requests[aliasPath])
			}
			pageRequests := 0
			for route, count := range requests {
				if strings.HasPrefix(route, "/docs/") {
					pageRequests += count
				}
			}
			if pageRequests != 25 {
				t.Fatalf("Docusaurus page requests = %d, want 25: %v", pageRequests, requests)
			}
			entries, readErr := os.ReadDir(result.Destination + "/documentation")
			if readErr != nil || len(entries) != 25 {
				t.Fatalf("generated Docusaurus documents = %d, %v", len(entries), readErr)
			}
			var finalPageProgress Progress
			for _, update := range progress {
				if update.Stage == "page" {
					finalPageProgress = update
				}
			}
			if finalPageProgress.Pages != 25 || finalPageProgress.Queued != 0 {
				t.Fatalf("final Docusaurus page progress does not show 25 pages and zero queued: %+v", progress)
			}
		})
	}
}

func TestImportHTMLDocusaurusDerivesCommonRootAcrossNestedSidebarSections(t *testing.T) {
	blogRequests := 0
	startAliasRequests := 0
	var server *httptest.Server
	page := func(route, title string) string {
		canonical := strings.TrimSuffix(route, "/")
		sidebar := `<li><a href="/docs/guide/start">Start</a></li><li><a href="/docs/guide/second">Guide</a></li><li><a href="/docs/reference/overview">Reference</a></li><li><a href="/blog">Blog</a></li>`
		return `<!doctype html><html class="plugin-docs docs-doc-page"><head><meta name="generator" content="Docusaurus v3.10.2"><link rel="canonical" href="` + server.URL + canonical + `"></head><body><ul class="theme-doc-sidebar-menu">` + sidebar + `</ul><div class="theme-doc-markdown markdown"><header><h1>` + title + `</h1></header><p>Nested Docusaurus content.</p></div></body></html>`
	}
	server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/guide/start/":
			fmt.Fprint(writer, page(request.URL.Path, "Start"))
		case "/docs/guide/start":
			startAliasRequests++
			fmt.Fprint(writer, page(request.URL.Path, "Start alias"))
		case "/docs/guide/second":
			fmt.Fprint(writer, page(request.URL.Path, "Guide"))
		case "/docs/reference/overview":
			fmt.Fprint(writer, page(request.URL.Path, "Reference"))
		case "/blog":
			blogRequests++
			fmt.Fprint(writer, page(request.URL.Path, "Blog"))
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	result, err := ImportHTML(context.Background(), "Nested Docusaurus", "v1", server.URL+"/docs/guide/start/", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 3 || result.Truncated || blogRequests != 0 || startAliasRequests != 0 {
		t.Fatalf("nested Docusaurus import = %+v, blog=%d alias=%d, err=%v", result, blogRequests, startAliasRequests, err)
	}
}

func TestImportHTMLDocusaurusScopesOmittedCurrentPageAndIgnoresUnrelatedSidebarLink(t *testing.T) {
	blogRequests := 0
	page := func(title, sidebar string) string {
		return `<!doctype html><html class="plugin-docs docs-doc-page"><head><meta name="generator" content="Docusaurus v3.10.2"></head><body><ul class="theme-doc-sidebar-menu">` + sidebar + `</ul><div class="theme-doc-markdown markdown"><header><h1>` + title + `</h1></header><p>Static content.</p></div></body></html>`
	}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/docs/start":
			fmt.Fprint(writer, page("Start", `<li><a href="/docs/second">Second</a></li><li><a href="/blog">Blog</a></li>`))
		case "/docs/second":
			fmt.Fprint(writer, page("Second", `<li><a href="/docs/start">Start</a></li>`))
		case "/blog":
			blogRequests++
			fmt.Fprint(writer, page("Blog", ""))
		default:
			http.NotFound(writer, request)
		}
	}))
	defer server.Close()

	result, err := ImportHTML(context.Background(), "Scoped Docusaurus", "v1", server.URL+"/docs/start", htmlRegressionOptions(t, server.Client()))
	if err != nil || result.Pages != 2 || result.Truncated || blogRequests != 0 {
		t.Fatalf("omitted-current Docusaurus import = %+v, blog requests=%d, err=%v", result, blogRequests, err)
	}
}

func TestDocusaurusDocumentationRootUsesSecureFallbackAndRejectsAliasAmbiguity(t *testing.T) {
	pageURL, _ := url.Parse("https://docs.test/docs/api/")
	document, err := parseHTML([]byte(`<!doctype html><html><body><ul class="theme-doc-sidebar-menu"><li><a href="/docs/api">API</a></li><li><a href="/blog/">Blog</a></li></ul></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	if root, rootErr := docusaurusDocumentationRoot(document, pageURL); rootErr != nil || root != "/docs/api/" {
		t.Fatalf("insufficient sidebar evidence did not retain narrow scope: root=%q err=%v", root, rootErr)
	}

	document, err = parseHTML([]byte(`<!doctype html><html><body><ul class="theme-doc-sidebar-menu"><li><a href="/docs/api?one">API one</a></li><li><a href="/docs/api?two">API two</a></li></ul></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	if root, rootErr := docusaurusDocumentationRoot(document, pageURL); rootErr == nil {
		t.Fatalf("ambiguous Docusaurus query aliases accepted as %q", root)
	}

	nestedURL, _ := url.Parse("https://docs.test/docs/guide/start/")
	document, err = parseHTML([]byte(`<!doctype html><html><body><ul class="theme-doc-sidebar-menu"><li><a href="/docs/guide/second">Guide</a></li><li><a href="/docs/reference/overview?one">Reference one</a></li><li><a href="/docs/reference/overview?two">Reference two</a></li></ul></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	if root, rootErr := docusaurusDocumentationRoot(document, nestedURL); rootErr == nil {
		t.Fatalf("ambiguous nested Docusaurus sidebar accepted as %q", root)
	}

	document, err = parseHTML([]byte(`<!doctype html><html><head><link rel="canonical" href="/docs/api"><link rel="canonical" href="/docs/api/"></head><body></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	if _, _, canonicalErr := docusaurusPageIdentity(document, pageURL, "/docs/"); canonicalErr == nil {
		t.Fatal("ambiguous Docusaurus canonical aliases accepted")
	}
}

func htmlRegressionOptions(t *testing.T, client *http.Client) Options {
	t.Helper()
	return Options{
		LibraryRoot: t.TempDir(), Rebuild: func(context.Context) error { return nil }, HTTPClient: client,
		HTMLScope: "path", MaxHTMLPages: -1, MaxHTMLDepth: -1,
	}
}
