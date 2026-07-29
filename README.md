# apis-mcp

[![release](https://img.shields.io/github/v/release/sairaph/apis-mcp?include_prereleases&label=release)](https://github.com/sairaph/apis-mcp/releases)
[![CI](https://github.com/sairaph/apis-mcp/actions/workflows/ci.yml/badge.svg)](https://github.com/sairaph/apis-mcp/actions/workflows/ci.yml)
[![license](https://img.shields.io/github/license/sairaph/apis-mcp)](#license)
[![platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-blue)](#what-it-does)

Give any AI agent a **searchable API reference and HTTP workspace** through
**7 MCP tools**, with an optional catalog of 12,989 documentation pages, a small
single binary, a universal installer, built-in client detection, and a full TUI
for you.

macOS:

```bash
curl -fsSL https://github.com/sairaph/apis-mcp/raw/main/install.sh | sh
```

Linux:

```bash
curl -fsSL https://github.com/sairaph/apis-mcp/raw/main/install.sh | sh
```

Windows (PowerShell):

```powershell
irm https://github.com/sairaph/apis-mcp/raw/main/install.ps1 | iex
```

The installer downloads the right binary for your OS and architecture, puts
`apis-mcp` on your `PATH`, then opens an interactive configurer. It detects your
AI clients and lets you choose both which clients to connect and which API
references to download, with live download, verification, and indexing status.
Run `apis-mcp configure` anytime to change either list.

> **After installing, open a new terminal.** The installer adds itself to your
> `PATH`, but the terminal it ran in keeps the `PATH` it started with.

## What it does

- **Download only the APIs you need** - choose from 12,989 searchable pages for
  Cloudflare, OpenRouter, Stripe, Tailscale, Tavily, Z.ai, and the example API.
- **A live API catalog** - `Manage APIs` refreshes directly from this repository,
  so newly published packs appear without an application update.
- **Seven direct MCP tools** - `apis_collections`, `apis_list`, `apis_pages`,
  `apis_search`, `apis_read`, `apis_call`, and `apis_sessions`.
- **Local and inspectable** - canonical Markdown with YAML frontmatter and a
  SQLite full-text index rebuilt entirely from those source files.
- **Bring your own references** - import OpenAPI 3.x, Swagger 2.x, `llms.txt`,
  and Markdown directories alongside official packs.
- **A capable HTTP workspace** - persistent response caching, UUIDv7 cookie
  sessions, retries, redirects, compressed responses, JSONPath previews, and
  background downloads.
- **A terminal launcher, CLI, and full-screen workspace** - open APIs or
  Configure from the launcher, browse documentation, make requests, and manage
  sessions outside an AI client.
- **Detects 13 AI clients** - configuration is safe to rerun and preserves
  other registered MCP servers.
- **Single static binary** - one CGO-free Go binary for Linux, macOS, and
  Windows on amd64 and arm64.

## Install

**macOS / Linux** - downloads the matching binary to `~/.apis-mcp/bin` and adds
it to your `PATH`:

```bash
curl -fsSL https://github.com/sairaph/apis-mcp/raw/main/install.sh | sh
```

**Windows (PowerShell)** - downloads to `%LOCALAPPDATA%\apis-mcp` and adds it to
your user `PATH`:

```powershell
irm https://github.com/sairaph/apis-mcp/raw/main/install.ps1 | iex
```

Each installer picks the matching asset for your OS and architecture from the
latest [GitHub Release](https://github.com/sairaph/apis-mcp/releases). If no
supported shell profile or interactive terminal is available, it prints the
exact `PATH` or configuration command to run later instead of silently skipping
setup.

Run `apis-mcp configure` whenever you want to change registered AI clients or
installed API packs. The pack selector uses up to three columns; press `space`
to toggle one API, `a` to select or clear all, and `r` to refresh the catalog.
Applying a changed selection shows per-pack and overall progress through cache
checks, downloads, verification, indexing, settings, and client registration.

Run `apis-mcp` in a terminal to open the launcher. Choose `APIs` for the
full-screen workspace or `Configure` for setup. Closing either returns to the
launcher. Large documentation branches load in bounded windows; use `pgup` and
`pgdn` or `[` and `]` to move between them.

## Documentation Library

Official documentation is distributed as content-addressed ZIP packs. Selected
packs are downloaded and verified under `~/.apis-mcp/packs`; changing the
selection atomically rebuilds the local search index. User document sets remain
separate under `~/.apis-mcp/library` and use this structure:

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
apis-mcp
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

Official pack sources remain under `library/builtin`, but are not embedded in
the binary. Regenerate and verify the downloadable artifacts after changing
them:

```sh
go run ./internal/cmd/pack generate
go run ./internal/cmd/pack verify
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
