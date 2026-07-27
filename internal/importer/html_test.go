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
