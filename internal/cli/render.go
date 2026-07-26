package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/sairaph/apis-mcp/internal/cache"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/internal/install"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"github.com/sairaph/apis-mcp/library"
)

// RenderHuman writes concise terminal output. MCP YAML/frontmatter rendering is
// deliberately not used by this presentation adapter.
func RenderHuman(writer io.Writer, value any) error {
	switch result := value.(type) {
	case string:
		_, err := fmt.Fprintln(writer, result)
		return err
	case library.CollectionsResult:
		fmt.Fprintf(writer, "Collections (page %d/%d, %d total)\n", result.Page, result.TotalPages, result.Total)
		for _, item := range result.Collections {
			fmt.Fprintf(writer, "  %-24s %4d APIs  %s\n", item.Collection, item.APICount, item.Name)
		}
		return nil
	case library.ListResult:
		fmt.Fprintf(writer, "API library (page %d/%d, %d total)\n", result.Page, result.TotalPages, result.Total)
		for _, item := range result.APIs {
			versions := make([]string, 0, len(item.Versions))
			for _, version := range item.Versions {
				versions = append(versions, fmt.Sprintf("%s [%s, %d pages]", version.Version, version.DocID, version.Pages))
			}
			fmt.Fprintf(writer, "  %s\n    %s\n", item.Name, strings.Join(versions, ", "))
		}
		return nil
	case library.PagesResult:
		fmt.Fprintf(writer, "%s pages at %q (page %d/%d, %d total)\n", result.DocID, result.Path, result.Page, result.TotalPages, result.Total)
		for _, path := range result.Paths {
			fmt.Fprintf(writer, "  / %-36s %d nested pages\n", path.Path, path.NestedPages)
		}
		for _, page := range result.Pages {
			fmt.Fprintf(writer, "  %-38s %s\n", page.PageID, page.Title)
		}
		return nil
	case library.SearchResult:
		fmt.Fprintf(writer, "Search %q in %s (page %d/%d, %d hits)\n", result.Query, result.DocID, result.Page, result.TotalPages, result.Total)
		for _, hit := range result.Hits {
			fmt.Fprintf(writer, "  %s:%d  %s\n    %s\n", hit.PageID, hit.Line, hit.Title, hit.Snippet)
		}
		return nil
	case library.ReadResult:
		fmt.Fprintf(writer, "%s / %s: %s (lines %d-%d of %d)\n\n", result.DocID, result.PageID, result.Title, result.Lines[0], result.Lines[1], result.TotalLines)
		_, err := io.WriteString(writer, result.Markdown)
		if err == nil && result.Markdown != "" && !strings.HasSuffix(result.Markdown, "\n") {
			_, err = io.WriteString(writer, "\n")
		}
		return err
	case httpcall.Result:
		fmt.Fprintf(writer, "%s %s\n", result.Request.Method, result.Request.Endpoint)
		fmt.Fprintf(writer, "  %s; HTTP %d; %d bytes; session %s\n", result.Response.State, result.Response.Status, result.Response.DecodedBytes, result.Request.SessionID)
		bodyPath := result.Cache.BodyPath
		if bodyPath == "" {
			bodyPath = result.Cache.FinalPath
		}
		if bodyPath != "" {
			fmt.Fprintf(writer, "  body: %s\n", bodyPath)
		}
		if result.Preview != nil {
			fmt.Fprintf(writer, "\n%s\n", result.Preview.Content)
		}
		return nil
	case []sessions.SessionInfo:
		fmt.Fprintf(writer, "Sessions (%d)\n", len(result))
		for _, item := range result {
			fmt.Fprintf(writer, "  %s  %d cookies  last used %s\n", item.ID, item.CookieCount, item.LastUsedAt.Format("2006-01-02 15:04:05Z"))
		}
		return nil
	case sessions.Inspection:
		fmt.Fprintf(writer, "Session %s (%d cookies)\n", result.Session.ID, result.Session.CookieCount)
		for _, cookie := range result.Cookies {
			fmt.Fprintf(writer, "  %-24s %s%s\n", cookie.Name, cookie.Domain, cookie.Path)
		}
		return nil
	case sessions.CleanupResult:
		_, err := fmt.Fprintf(writer, "Removed %d expired sessions.\n", result.Removed)
		return err
	case cache.CleanupResult:
		_, err := fmt.Fprintf(writer, "Removed %d expired cache entries and %d orphans.\n", result.RemovedEntries, result.RemovedOrphans)
		return err
	case []install.Result:
		if len(result) == 0 {
			_, err := fmt.Fprintln(writer, "No detected clients. Use --client ID or --all to create configuration.")
			return err
		}
		for _, item := range result {
			action := "unchanged"
			if item.Changed && item.Removed {
				action = "removed"
			} else if item.Changed {
				action = "configured"
			}
			if item.DryRun && item.Changed {
				action = "would be " + action
			}
			fmt.Fprintf(writer, "%-20s %-20s %s\n", item.Client.Name, action, item.Path)
			if item.Backup != "" {
				fmt.Fprintf(writer, "  backup: %s\n", item.Backup)
			}
		}
		return nil
	case []Diagnostic:
		for _, item := range result {
			state := "OK"
			if !item.OK {
				state = "WARN"
			}
			fmt.Fprintf(writer, "%-5s %-18s %s\n", state, item.Name, item.Detail)
		}
		return nil
	case importer.Result:
		_, err := fmt.Fprintf(writer, "Imported %s %s (%s): %d pages\n  source: %s\n  library: %s\n", result.Name, result.Version, result.Kind, result.Pages, result.Source, result.Destination)
		return err
	default:
		raw, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			return fmt.Errorf("render human output: %w", err)
		}
		_, err = fmt.Fprintln(writer, string(raw))
		return err
	}
}
