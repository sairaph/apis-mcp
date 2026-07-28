# apis-mcp

`apis-mcp` is a local API documentation library, HTTP workspace, CLI, and MCP
server in one Go binary. It gives AI clients searchable API references and a
general HTTP caller without requiring a hosted service.

## Features

- Seven direct MCP tools: `apis_collections`, `apis_list`, `apis_pages`,
  `apis_search`, `apis_read`, `apis_call`, and `apis_sessions`.
- Canonical Markdown documentation with YAML frontmatter.
- OpenAPI 3.x, Swagger 2.x, `llms.txt`, and Markdown-directory imports.
- SQLite FTS search rebuilt entirely from canonical Markdown.
- Persistent HTTP response cache and UUIDv7 cookie sessions.
- Retries, redirects, compressed responses, JSONPath previews, and background
  downloads.
- Full-screen terminal app and human-oriented one-shot CLI commands.
- Idempotent configuration for 13 MCP clients.
- One CGO-free binary for Linux, macOS, and Windows on amd64 and arm64.

## Install

Linux or macOS:

```sh
curl -fsSL https://raw.githubusercontent.com/sairaph/apis-mcp/main/install.sh | sh
```

Windows PowerShell:

```powershell
irm https://raw.githubusercontent.com/sairaph/apis-mcp/main/install.ps1 | iex
```

Run `apis-mcp configure` again whenever you want to register newly installed AI
clients. Run `apis-mcp` in a terminal to open the full-screen application.
The installers resolve one concrete GitHub release, verify the selected binary
against that release's `SHA256SUMS.txt`, and only then replace an existing
installation. The Unix installer updates existing startup files for Bash, Zsh,
Fish, and POSIX shells without creating a profile. If no supported shell
profile or interactive terminal is available, the installers print the exact
PATH or configuration command to run later instead of silently skipping setup.

## Documentation Library

Built-in documentation is embedded in the binary. User document sets live in
`~/.apis-mcp/library` and use this structure:

```text
my-api/
  v1/
    _index.md
    overview.md
    resources/
      create.md
```

`_index.md` contains API metadata:

```markdown
---
name: My API
version: v1
description: API description.
collections: [examples]
source_root: https://docs.example.com/
---
```

Each page requires a title:

```markdown
---
title: Create a resource
api_endpoints: [/resources]
http_methods: [POST]
operation_ids: [createResource]
---

# Create a resource
```

Import existing sources:

```sh
apis-mcp import markdown ./canonical-docs
apis-mcp import openapi "Pet API" v1 ./openapi.yaml
apis-mcp import llms "Example API" 2026 https://docs.example.com/llms.txt
apis-mcp import html "Example API" 2026 https://docs.example.com/api/
apis-mcp rebuild
```

## CLI

```text
apis-mcp collections
apis-mcp list --name stripe
apis-mcp pages DOC_ID --path payments
apis-mcp search DOC_ID "create payment"
apis-mcp read DOC_ID PAGE_ID --lines 20:80
apis-mcp call GET https://api.example.com/items
apis-mcp sessions list
apis-mcp cache cleanup
apis-mcp doctor
apis-mcp configure
```

Run `apis-mcp help` for the complete command reference.

## Local Trust Model

This is a local tool acting with the authority of the user who runs it. HTTP
calls may access arbitrary HTTP/HTTPS destinations and header or payload inputs
may reference local JSON files. The user is responsible for the endpoints,
files, credentials, request effects, and downloaded content selected through
the CLI or their MCP client.

TLS verification remains enabled by default. Response-size, disk-reserve,
redirect, retry, timeout, and private-file protections prevent accidental local
resource exhaustion and cache corruption.

## Build

Go 1.26 or newer is required.

```sh
go test ./...
go vet ./...
CGO_ENABLED=0 go build -trimpath -o apis-mcp .
```

The repository also includes a development-only URL ingestion command. It
detects the source framework or format, dispatches to the matching importer,
writes canonical Markdown, and rebuilds a searchable SQLite generation without
adding another shipped CLI:

```sh
go run -tags=dev ./internal/cmd/ingest start -out /tmp/apis-ingest \
  -name "Petstore" -version v2 -collections examples,petstore \
  https://petstore.swagger.io/
```

`-name` defaults to the URL host and `-version` defaults to `latest`. The
repeatable `-collections` flag accepts comma-separated collection IDs. The
detached job defaults to unlimited pages and depth while staying under the
starting documentation path. Use `-scope domain` to crawl the entire origin,
or set optional `-max-pages` and `-max-depth` safeguards. Large finite API
catalogs can raise the 16 MiB per-source and 64 MiB aggregate download limits
with `-max-source-bytes` and `-max-total-bytes`.

Control and subscribe to the persisted job from separate processes:

```sh
go run -tags=dev ./internal/cmd/ingest status -out /tmp/apis-ingest JOB_ID
go run -tags=dev ./internal/cmd/ingest watch -out /tmp/apis-ingest JOB_ID
go run -tags=dev ./internal/cmd/ingest cancel -out /tmp/apis-ingest JOB_ID
go run -tags=dev ./internal/cmd/ingest list -out /tmp/apis-ingest
```

`watch` streams durable JSON events for detection, each page, publication,
indexing, cancellation, failure, and completion. Successful jobs write
canonical Markdown and a searchable SQLite generation below `-out`; canceled
jobs publish neither. The command's `dev` build constraint and separate `main`
package keep it out of release binaries.
Framework and format coverage is tracked in
[`docs/ingestion-matrix.md`](docs/ingestion-matrix.md).

The full agent-facing behavior is specified in
[`docs/tool-contract.md`](docs/tool-contract.md).

## License

MIT
