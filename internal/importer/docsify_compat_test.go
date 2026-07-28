package importer

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sairaph/apis-mcp/library"
)

func docsifyCompatShell(config string) string {
	return `<!doctype html><html><body><div id="app"></div><script>window.$docsify = ` + config + `;</script><script src="docsify.min.js"></script></body></html>`
}

func docsifyCompatDocument(t *testing.T, config string) *htmlNode {
	t.Helper()
	document, err := parseHTML([]byte(docsifyCompatShell(config)))
	if err != nil {
		t.Fatal(err)
	}
	return document
}

func TestDocsifyCompatibilityDerivesStandardGitHubPagesSource(t *testing.T) {
	const config = `{repo: 'https://github.com/theoephraim/node-google-spreadsheet', loadSidebar: true}`
	document := docsifyCompatDocument(t, config)
	shellURL, _ := url.Parse("https://theoephraim.github.io/node-google-spreadsheet/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil || len(selections) != 1 {
		t.Fatalf("GitHub Pages selections: %+v, %v", selections, err)
	}
	selection := selections[0]
	if selection.owner != "theoephraim" || selection.repo != "node-google-spreadsheet" || selection.ref != "" || selection.path != "" || selection.exact || !selection.pagesDeployment || selection.shellIdentity == "" || selection.shellURL != shellURL.String() {
		t.Fatalf("GitHub Pages selection = %+v", selection)
	}

	const deploymentID = "12345"
	const commitSHA = "1111111111111111111111111111111111111111"
	const treeSHA = "2222222222222222222222222222222222222222"
	t.Setenv("GITHUB_TOKEN", "pages-token")
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.URL.Host == "api.github.com" && request.Header.Get("Authorization") != "Bearer pages-token" {
			t.Errorf("GitHub API authorization = %q", request.Header.Get("Authorization"))
		}
		if request.URL.Host != "api.github.com" && request.Header.Get("Authorization") != "" {
			t.Errorf("GitHub token leaked to %s", request.URL.Host)
		}
		switch {
		case request.URL.Host == "theoephraim.github.io":
			return docsifyTestResponse(request, http.StatusOK, docsifyCompatShell(config)), nil
		case strings.HasSuffix(request.URL.Path, "/deployments"):
			return docsifyTestResponse(request, http.StatusOK, `[{"id":`+deploymentID+`,"sha":"`+commitSHA+`","ref":"main"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/deployments/"+deploymentID+"/statuses"):
			return docsifyTestResponse(request, http.StatusOK, `[{"state":"success","environment_url":"https://theoephraim.github.io/node-google-spreadsheet/"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/commits/"+commitSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commitSHA+`","commit":{"tree":{"sha":"`+treeSHA+`"}}}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+commitSHA+"/index.html"):
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+commitSHA+"/docs/index.html"):
			return docsifyTestResponse(request, http.StatusOK, docsifyCompatShell(config)), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+commitSHA+"/docs/guide.md"):
			return docsifyTestResponse(request, http.StatusOK, "# Guide\n\nDeployment source proof.\n"), nil
		case strings.HasSuffix(request.URL.Path, "/git/trees/"+treeSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+treeSHA+`","truncated":false,"tree":[{"path":"docs/guide.md","type":"blob","size":10}]}`), nil
		default:
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		}
	})}
	root := t.TempDir()
	index := filepath.Join(t.TempDir(), "library.sqlite")
	result, err := ImportDocsify(context.Background(), "Google Spreadsheet", "v1", shellURL.String(), Options{
		LibraryRoot: root, HTTPClient: client, Rebuild: func(ctx context.Context) error {
			return library.Rebuild(ctx, library.Options{UserRoot: root, IndexPath: index, ExcludeBuiltin: true})
		},
	})
	if err != nil || result.Pages != 1 || result.Sources != 1 {
		t.Fatalf("verified GitHub Pages import: %+v, %v", result, err)
	}
	snapshot, err := library.Open(context.Background(), library.Options{UserRoot: root, IndexPath: index, ExcludeBuiltin: true})
	if err != nil {
		t.Fatal(err)
	}
	defer snapshot.Close()
	search, err := snapshot.Search(context.Background(), library.SearchRequest{DocID: "google-spreadsheet-v1", Query: "source proof"})
	if err != nil || search.Total == 0 {
		t.Fatalf("deployment-derived Docsify source was not indexed: %+v, %v", search, err)
	}
}

func TestDocsifyCompatibilityIgnoresEllipsisInsideString(t *testing.T) {
	document := docsifyCompatDocument(t, `{repo: 'https://github.com/chiru-labs/ERC721A', placeholder: 'Search...'}`)
	shellURL, _ := url.Parse("https://chiru-labs.github.io/ERC721A/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil || len(selections) != 1 || !selections[0].pagesDeployment {
		t.Fatalf("literal ellipsis selections: %+v, %v", selections, err)
	}
}

func TestDocsifyCompatibilityResolvesPollySources(t *testing.T) {
	const config = `{
		repo: 'https://github.com/Netflix/pollyjs',
		alias: {
			'.*?/CHANGELOG': 'https://raw.githubusercontent.com/Netflix/pollyjs/master/CHANGELOG',
			'.*?/CONTRIBUTING': 'https://raw.githubusercontent.com/Netflix/pollyjs/master/CONTRIBUTING'
		},
		homepage: 'https://raw.githubusercontent.com/Netflix/pollyjs/master/README.md'
	}`
	document := docsifyCompatDocument(t, config)
	shellURL, _ := url.Parse("https://netflix.github.io/pollyjs/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil || len(selections) != 4 {
		t.Fatalf("Polly selections: %+v, %v", selections, err)
	}

	const deploymentID = "54321"
	const pagesCommitSHA = "5555555555555555555555555555555555555555"
	const pagesTreeSHA = "6666666666666666666666666666666666666666"
	const masterCommitSHA = "7777777777777777777777777777777777777777"
	const masterTreeSHA = "8888888888888888888888888888888888888888"
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		switch {
		case strings.HasSuffix(request.URL.Path, "/deployments"):
			return docsifyTestResponse(request, http.StatusOK, `[{"id":`+deploymentID+`,"sha":"`+pagesCommitSHA+`","ref":"gh-pages"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/deployments/"+deploymentID+"/statuses"):
			return docsifyTestResponse(request, http.StatusOK, `[{"state":"success","environment_url":"https://netflix.github.io/pollyjs/"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/commits/"+pagesCommitSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+pagesCommitSHA+`","commit":{"tree":{"sha":"`+pagesTreeSHA+`"}}}`), nil
		case strings.HasSuffix(request.URL.Path, "/commits/master"):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+masterCommitSHA+`","commit":{"tree":{"sha":"`+masterTreeSHA+`"}}}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+pagesCommitSHA+"/index.html"):
			return docsifyTestResponse(request, http.StatusOK, docsifyCompatShell(config)), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+pagesCommitSHA+"/docs/index.html"):
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		case strings.HasSuffix(request.URL.Path, "/git/trees/"+pagesTreeSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+pagesTreeSHA+`","truncated":false,"tree":[{"path":"_sidebar.md","type":"blob","size":10},{"path":"getting-started.md","type":"blob","size":10}]}`), nil
		case strings.HasSuffix(request.URL.Path, "/git/trees/"+masterTreeSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+masterTreeSHA+`","truncated":false,"tree":[{"path":"README.md","type":"blob","size":10},{"path":"CHANGELOG.md","type":"blob","size":10},{"path":"CONTRIBUTING.md","type":"blob","size":10},{"path":"packages/internal.md","type":"blob","size":10}]}`), nil
		default:
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		}
	})}
	reader := &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	files, err := docsifyGitHubInventory(context.Background(), selections, reader)
	if err != nil {
		t.Fatal(err)
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, file.path)
		if file.ref != pagesCommitSHA && file.ref != "master" || file.ref == pagesCommitSHA && file.commit != pagesCommitSHA || file.ref == "master" && file.commit != masterCommitSHA {
			t.Errorf("file was not pinned to its advertised source: %+v", file)
		}
	}
	want := []string{"_sidebar.md", "getting-started.md", "CHANGELOG.md", "CONTRIBUTING.md", "README.md"}
	if strings.Join(paths, "\n") != strings.Join(want, "\n") {
		t.Fatalf("Polly inventory = %q, want %q", paths, want)
	}
}

func TestDocsifyCompatibilityRejectsUnrelatedPagesShellWithoutPublication(t *testing.T) {
	const deploymentID = "24680"
	const commitSHA = "7777777777777777777777777777777777777777"
	const treeSHA = "8888888888888888888888888888888888888888"
	deployed := docsifyCompatShell(`{repo: 'https://github.com/example/project', loadSidebar: true}`)
	rebuilds := 0
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		switch {
		case request.URL.Host == "example.github.io":
			return docsifyTestResponse(request, http.StatusOK, deployed), nil
		case strings.HasSuffix(request.URL.Path, "/deployments"):
			return docsifyTestResponse(request, http.StatusOK, `[{"id":`+deploymentID+`,"sha":"`+commitSHA+`","ref":"main"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/deployments/"+deploymentID+"/statuses"):
			return docsifyTestResponse(request, http.StatusOK, `[{"state":"success","environment_url":"https://example.github.io/project/"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/commits/"+commitSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commitSHA+`","commit":{"tree":{"sha":"`+treeSHA+`"}}}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+commitSHA+"/index.html"):
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/"+commitSHA+"/docs/index.html"):
			return docsifyTestResponse(request, http.StatusOK, docsifyCompatShell(`{repo: 'https://github.com/example/project', loadSidebar: false}`)), nil
		default:
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		}
	})}
	root := t.TempDir()
	_, err := ImportDocsify(context.Background(), "Unrelated Pages", "v1", "https://example.github.io/project/", Options{
		LibraryRoot: root,
		Rebuild: func(context.Context) error {
			rebuilds++
			return nil
		},
		HTTPClient: client,
	})
	if err == nil || !strings.Contains(err.Error(), "has 0 matching repository source roots") {
		t.Fatalf("unrelated docs/index.html accepted: %v", err)
	}
	entries, readErr := os.ReadDir(root)
	if readErr != nil || len(entries) != 0 || rebuilds != 0 {
		t.Fatalf("failed Pages proof published output: entries=%v rebuilds=%d err=%v", entries, rebuilds, readErr)
	}
}

func TestDocsifyCompatibilityDoesNotGuessDocsForRootOrPagesBranch(t *testing.T) {
	document := docsifyCompatDocument(t, `{repo: 'https://github.com/example/project'}`)
	shellURL, _ := url.Parse("https://example.github.io/project/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		switch {
		case strings.HasSuffix(request.URL.Path, "/deployments"):
			return docsifyTestResponse(request, http.StatusOK, `[]`), nil
		default:
			t.Fatalf("missing deployment metadata made request to %s", request.URL)
			return nil, nil
		}
	})}
	reader := &sourceReader{client: client, perSource: DefaultMaxSourceBytes, total: DefaultMaxTotalBytes}
	if _, err := docsifyGitHubInventory(context.Background(), selections, reader); err == nil || !strings.Contains(err.Error(), "has no valid Pages deployment") {
		t.Fatalf("ambiguous Pages source accepted: %v", err)
	}
}

func TestDocsifyCompatibilityRejectsDuplicateDeploymentShellRoots(t *testing.T) {
	const deploymentID = "13579"
	const commitSHA = "9999999999999999999999999999999999999999"
	const treeSHA = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	deployed := docsifyCompatShell(`{repo: 'https://github.com/example/project', loadSidebar: true}`)
	document := docsifyCompatDocument(t, `{repo: 'https://github.com/example/project', loadSidebar: true}`)
	shellURL, _ := url.Parse("https://example.github.io/project/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{Transport: docsifyRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		switch {
		case strings.HasSuffix(request.URL.Path, "/deployments"):
			return docsifyTestResponse(request, http.StatusOK, `[{"id":`+deploymentID+`,"sha":"`+commitSHA+`","ref":"main"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/deployments/"+deploymentID+"/statuses"):
			return docsifyTestResponse(request, http.StatusOK, `[{"state":"success","environment_url":"https://example.github.io/project/"}]`), nil
		case strings.HasSuffix(request.URL.Path, "/commits/"+commitSHA):
			return docsifyTestResponse(request, http.StatusOK, `{"sha":"`+commitSHA+`","commit":{"tree":{"sha":"`+treeSHA+`"}}}`), nil
		case request.URL.Host == "raw.githubusercontent.com" && strings.HasSuffix(request.URL.Path, "/index.html"):
			return docsifyTestResponse(request, http.StatusOK, deployed), nil
		default:
			return docsifyTestResponse(request, http.StatusNotFound, `{}`), nil
		}
	})}
	reader := newSourceReader(Options{HTTPClient: client, MaxSourceBytes: DefaultMaxSourceBytes, MaxTotalBytes: DefaultMaxTotalBytes})
	if _, err := docsifyGitHubInventory(context.Background(), selections, reader); err == nil || !strings.Contains(err.Error(), "has 2 matching repository source roots") {
		t.Fatalf("duplicate Pages roots accepted: %v", err)
	}
}

func TestDocsifyCompatibilityExplicitRootSkipsPagesInference(t *testing.T) {
	document := docsifyCompatDocument(t, `{repo: 'https://github.com/example/project', basePath: 'https://raw.githubusercontent.com/example/project/main/docs'}`)
	shellURL, _ := url.Parse("https://example.github.io/project/")
	selections, err := docsifyGitHubSources(document, shellURL)
	if err != nil || len(selections) != 1 || selections[0].pagesDeployment || selections[0].shellIdentity != "" || selections[0].ref != "main" || selections[0].path != "docs" {
		t.Fatalf("explicit root selections: %+v, %v", selections, err)
	}
}
