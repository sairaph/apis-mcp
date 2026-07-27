package importer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type docsifyRoundTripFunc func(*http.Request) (*http.Response, error)

func (function docsifyRoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func docsifyTestResponse(request *http.Request, status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Status: fmt.Sprintf("%d %s", status, http.StatusText(status)), Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body)), Request: request}
}

func TestDocsifyDetectionRequiresRuntimeAndConfiguration(t *testing.T) {
	for _, test := range []struct {
		name string
		html string
		want bool
	}{
		{name: "official", html: `<div id="app"></div><script>window.$docsify = {loadSidebar: true}</script><script src="//cdn.jsdelivr.net/npm/docsify@5/dist/docsify.min.js"></script>`, want: true},
		{name: "package root", html: `<div id="app"></div><script>window.$docsify = {}</script><script src="//cdn.jsdelivr.net/npm/docsify@5"></script>`, want: true},
		{name: "fork", html: `<div id="app"></div><script>window.$docsify = {name: 'Docs'}</script><script src="https://cdn.jsdelivr.net/gh/example/docsify@5/docsify.js"></script>`, want: true},
		{name: "plugin only", html: `<div id="app"></div><script>window.$docsify = {}</script><script src="docsify-plugin-search.min.js"></script>`},
		{name: "runtime only", html: `<div id="app"></div><script src="docsify.min.js"></script>`},
	} {
		t.Run(test.name, func(t *testing.T) {
			document, err := parseHTML([]byte(`<!doctype html><html><body>` + test.html + `</body></html>`))
			if err != nil {
				t.Fatal(err)
			}
			if got := looksLikeDocsify(document); got != test.want {
				t.Fatalf("Docsify detection = %v, want %v", got, test.want)
			}
		})
	}
}

func TestDocsifyDiscoversStaticGitHubSources(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><body><div id="app"></div><script>
window.$docsify = { alias: {
  '/translated/(.*)': 'https://cdn.jsdelivr.net/gh/example/translated@main/$1',
  '/awesome': 'https://raw.githubusercontent.com/example/awesome/master/README.md',
  '/changes': 'https://raw.githubusercontent.com/example/docs/main/CHANGELOG.md'
}};
// const unused = 'https://github.com/example/dead/tree/main/docs/';
const edit = 'https://github.com/example/docs/blob/develop/docs/' + route;
</script><script src="docsify.min.js"></script></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	shellURL, _ := url.Parse("https://docs.test/#/guide")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil {
		t.Fatal(err)
	}
	if len(selections) != 4 {
		t.Fatalf("Docsify selections = %+v", selections)
	}
	wanted := map[string]bool{
		"example/awesome/master/README.md/true": true,
		"example/docs/develop/docs/false":       true,
		"example/docs/main/CHANGELOG.md/true":   true,
		"example/translated/main//false":        true,
	}
	for _, selection := range selections {
		key := selection.owner + "/" + selection.repo + "/" + selection.ref + "/" + selection.path + "/"
		if selection.exact {
			key += "true"
		} else {
			key += "false"
		}
		if !wanted[key] {
			t.Errorf("unexpected Docsify selection %q", key)
		}
	}
}

func TestDocsifyGitHubBasePathRejectsGenericHTTP(t *testing.T) {
	if _, err := docsifyGitHubBasePath("https://example.com/docs"); err == nil {
		t.Fatal("generic HTTP basePath accepted without a finite inventory")
	}
}

func TestDocsifyRejectsDynamicBasePathWithUnrelatedGitHubURL(t *testing.T) {
	for _, config := range []string{
		`{basePath: chooseSource(), repo: 'https://github.com/example/docs/tree/main/docs/'}`,
		`{basePath, repo: 'https://github.com/example/docs/tree/main/docs/'}`,
		`makeConfig('https://github.com/example/docs/tree/main/docs/')`,
		`{...makeConfig(), repo: 'https://github.com/example/docs/tree/main/docs/'}`,
		`{}; window.$docsify = makeConfig('https://github.com/example/docs/tree/main/docs/')`,
		`{}; window.$docsify.basePath = 'https://raw.githubusercontent.com/example/docs/main/docs'`,
	} {
		document, err := parseHTML([]byte(`<!doctype html><html><body><div id="app"></div><script>window.$docsify = ` + config + `;</script><script src="docsify.min.js"></script></body></html>`))
		if err != nil {
			t.Fatal(err)
		}
		shellURL, _ := url.Parse("https://docs.test/")
		if _, err := docsifyGitHubSources(document, shellURL); err == nil || !strings.Contains(err.Error(), "dynamic Docsify") {
			t.Errorf("dynamic config %q accepted: %v", config, err)
		}
	}
}

func TestDocsifyRejectsQuerySourceOutsideBasePath(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><body><div id="app"></div><script>window.$docsify = {}</script><script src="docsify.min.js"></script></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	for _, source := range []string{"../README.md", `\\example.com\README.md`} {
		shellURL, _ := url.Parse("https://docs.test/?basePath=https%3A%2F%2Fraw.githubusercontent.com%2Fexample%2Fdocs%2Fmain%2Fdocs&homepage=" + url.QueryEscape(source))
		if _, err := docsifyGitHubSources(document, shellURL); err == nil {
			t.Errorf("escaping query source %q accepted", source)
		}
	}
}

func TestDocsifyGitHubRefValidation(t *testing.T) {
	for _, ref := range []string{"main", "release/v2+docs", "feature/user@example"} {
		if !validGitHubRef(ref) {
			t.Errorf("valid ref %q rejected", ref)
		}
	}
	for _, ref := range []string{"", "../main", "bad ref", "topic.lock", "topic@{one"} {
		if validGitHubRef(ref) {
			t.Errorf("invalid ref %q accepted", ref)
		}
	}
}

func TestSourceReaderScopesGitHubToken(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "fixture-token")
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.URL.Scheme == "https" && request.URL.Host == "api.github.com" {
			if request.Header.Get("Authorization") != "Bearer fixture-token" {
				t.Errorf("GitHub authorization = %q", request.Header.Get("Authorization"))
			}
		} else if request.Header.Get("Authorization") != "" {
			t.Errorf("authorization leaked to %s", request.URL.Host)
		}
		return docsifyTestResponse(request, http.StatusOK, `{}`), nil
	})}
	reader := newSourceReader(Options{HTTPClient: client, MaxSourceBytes: DefaultMaxSourceBytes, MaxTotalBytes: DefaultMaxTotalBytes})
	if _, _, err := reader.readFromOrigin(context.Background(), "https://api.github.com/rate_limit", nil, "https://api.github.com"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := reader.readFromOrigin(context.Background(), "http://api.github.com/rate_limit", nil, ""); err != nil {
		t.Fatal(err)
	}
	if _, _, err := reader.readFromOrigin(context.Background(), "https://example.test/source", nil, "https://example.test"); err != nil {
		t.Fatal(err)
	}
}

func TestDocsifyExtractsEncodedGitHubRef(t *testing.T) {
	document, err := parseHTML([]byte(`<!doctype html><html><body><div id="app"></div><script>window.$docsify = {alias: {'/guide': 'https://raw.githubusercontent.com/example/docs/release%23one/guide.md'}, repo: 'https://github.com/example/docs/tree/release%23one/docs/'}</script><script src="docsify.min.js"></script></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	shellURL, _ := url.Parse("https://docs.test/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil || len(selections) != 2 {
		t.Fatalf("encoded ref sources: %+v, %v", selections, err)
	}
	for _, selection := range selections {
		if selection.ref != "release#one" {
			t.Errorf("decoded ref = %q", selection.ref)
		}
	}
}

func TestDocsifyRejectsOutputPathCollisions(t *testing.T) {
	err := validateDocsifyOutputPaths([]docsifyDocument{
		{source: "one", output: "documentation/Example/README.md"},
		{source: "two", output: "documentation/example/readme.md"},
	})
	if err == nil || !strings.Contains(err.Error(), "collide") {
		t.Fatalf("output collision accepted: %v", err)
	}
}

func TestDocsifyRejectsNonPortableRepositoryPaths(t *testing.T) {
	for _, value := range []string{"docs./README.md", "docs/CON.md", "docs/trailing /README.md", "docs/a:b.md"} {
		if safeRepositoryPath(value) {
			t.Errorf("non-portable path %q accepted", value)
		}
	}
	for _, value := range []string{"docs/README.md", "docs/concept.md", "docs/nested/guide.md"} {
		if !safeRepositoryPath(value) {
			t.Errorf("portable path %q rejected", value)
		}
	}
}

func TestDocsifyInventoryRejectsSelectedNonPortableMarkdown(t *testing.T) {
	const commitSHA = "7777777777777777777777777777777777777777"
	const treeSHA = "8888888888888888888888888888888888888888"
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		if strings.Contains(request.URL.Path, "/commits/") {
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commitSHA+`","commit":{"tree":{"sha":"`+treeSHA+`"}}}`), nil
		}
		if strings.Contains(request.URL.Path, "/git/trees/") {
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+treeSHA+`","truncated":false,"tree":[{"path":"README.md","type":"blob","size":10},{"path":"CON.md","type":"blob","size":10}]}`), nil
		}
		return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
	})}
	reader := &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	_, err := docsifyGitHubInventory(context.Background(), []docsifyGitHubSelection{{owner: "example", repo: "docs", ref: "main"}}, reader)
	if err == nil || !strings.Contains(err.Error(), "non-portable Markdown path") {
		t.Fatalf("non-portable selected Markdown accepted: %v", err)
	}
}

func TestDocsifyBoundsRefCandidates(t *testing.T) {
	selection := docsifyGitHubSelection{owner: "example", repo: "docs", ref: "main", path: strings.Repeat("segment/", docsifyMaxRefCandidates) + "docs"}
	reader := &sourceReader{client: &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		t.Fatal("bounded ref selection made a request")
		return nil, nil
	})}, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	if _, _, err := docsifyResolveSelectionRefs(context.Background(), []docsifyGitHubSelection{selection}, reader); err == nil || !strings.Contains(err.Error(), "ref candidates") {
		t.Fatalf("unbounded ref selection accepted: %v", err)
	}
}

func TestDocsifyInventoryResolvesSlashRef(t *testing.T) {
	const commitSHA = "1111111111111111111111111111111111111111"
	const treeSHA = "2222222222222222222222222222222222222222"
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		switch {
		case strings.HasSuffix(request.URL.Path, "/commits/release/v2/docs"):
			return docsifyTestResponse(request, http.StatusUnprocessableEntity, `{}`), nil
		case strings.HasSuffix(request.URL.Path, "/commits/release/v2"):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commitSHA+`","commit":{"tree":{"sha":"`+treeSHA+`"}}}`), nil
		case strings.HasSuffix(request.URL.Path, "/git/trees/"+treeSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+treeSHA+`","truncated":false,"tree":[{"path":"docs/guide.md","type":"blob","size":10}]}`), nil
		default:
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		}
	})}
	reader := &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	files, err := docsifyGitHubInventory(context.Background(), []docsifyGitHubSelection{{owner: "example", repo: "docs", ref: "release", path: "v2/docs"}}, reader)
	if err != nil || len(files) != 1 || files[0].ref != "release/v2" || files[0].path != "docs/guide.md" {
		t.Fatalf("slash ref inventory: %+v, %v", files, err)
	}
}

func TestDocsifyInventoryKeepsSamePathAcrossRefs(t *testing.T) {
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		if strings.Contains(request.URL.Path, "/commits/") {
			ref := pathBase(request.URL.Path)
			commit, tree := strings.Repeat("1", 40), strings.Repeat("a", 40)
			if ref == "release" {
				commit, tree = strings.Repeat("2", 40), strings.Repeat("b", 40)
			}
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commit+`","commit":{"tree":{"sha":"`+tree+`"}}}`), nil
		}
		if strings.Contains(request.URL.Path, "/git/trees/") {
			tree := pathBase(request.URL.Path)
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+tree+`","truncated":false,"tree":[{"path":"README.md","type":"blob","size":10}]}`), nil
		}
		return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
	})}
	reader := &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	files, err := docsifyGitHubInventory(context.Background(), []docsifyGitHubSelection{
		{owner: "example", repo: "docs", ref: "main", path: "README.md", exact: true},
		{owner: "example", repo: "docs", ref: "release", path: "README.md", exact: true},
	}, reader)
	if err != nil || len(files) != 2 {
		t.Fatalf("multi-ref inventory: %+v, %v", files, err)
	}
}

func TestDocsifyInventoryDeduplicatesRefAliases(t *testing.T) {
	const commitSHA = "3333333333333333333333333333333333333333"
	const treeSHA = "4444444444444444444444444444444444444444"
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		if strings.Contains(request.URL.Path, "/commits/") {
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commitSHA+`","commit":{"tree":{"sha":"`+treeSHA+`"}}}`), nil
		}
		if strings.Contains(request.URL.Path, "/git/trees/") {
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+treeSHA+`","truncated":false,"tree":[{"path":"README.md","type":"blob","size":10}]}`), nil
		}
		return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
	})}
	reader := &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	files, err := docsifyGitHubInventory(context.Background(), []docsifyGitHubSelection{
		{owner: "example", repo: "docs", ref: "main", path: "README.md", exact: true},
		{owner: "example", repo: "docs", ref: "stable", path: "README.md", exact: true},
	}, reader)
	if err != nil || len(files) != 1 {
		t.Fatalf("ref alias inventory: %+v, %v", files, err)
	}
}

func pathBase(value string) string {
	parts := strings.Split(strings.Trim(value, "/"), "/")
	return parts[len(parts)-1]
}
