# Agent Tool Contract

Status: first-release contract.

Project, repository, binary, and MCP server name: `apis-mcp`.

The first release supports MCP over stdio only. The `apis-mcp mcp` command
starts the stdio server used by installed AI clients. Streamable HTTP, SSE, and
other network transports are not exposed.

The server uses the official Go MCP SDK at
`github.com/modelcontextprotocol/go-sdk`. Tool handlers decode into typed Go
input structures, while their JSON Schemas explicitly declare validation,
defaults, and configuration-dependent limits. MCP framing and lifecycle
behavior are delegated to the SDK rather than implemented as a custom JSON-RPC
transport.

The binary has the same terminal-sensitive entrypoint model as `favro-mcp`.
Bare invocation on a real stdin/stdout terminal opens the full interactive CLI
application, including its human-oriented renderer and configuration screens.
Bare invocation without a terminal runs the stdio MCP server for client
compatibility. `apis-mcp mcp` always starts the stdio server explicitly.

## Installation And Runtime Model

The first release follows the single-binary installation and process model used
by `favro-mcp`. Platform bootstrap installers download the matching release
binary into a per-user application directory, make the command available on
`PATH`, and immediately launch the interactive setup flow when a terminal is
available. Setup can also be reopened later from the full CLI application or by
an explicit configuration command.

Setup refreshes the official documentation-pack catalog and lets the user choose
which API families to keep locally. Opening setup again fetches newly published
catalog entries without requiring a new application release. An unavailable
catalog never removes or disables already installed packs.

The setup flow detects supported MCP clients and registers selected clients
with an absolute path to the installed binary plus the `mcp` subcommand. Client
configuration never relies on the client's working directory or on `apis-mcp`
being discoverable through that client's `PATH`. Reconfiguration is idempotent
and preserves unrelated client configuration.

Each AI client starts its own `apis-mcp mcp` stdio process. There is no daemon,
system service, shared network listener, or separately managed MCP process in
the first release. Persisted documentation, settings, call cache, and cookie
sessions are shared through the per-user application data directory, while
in-memory task ownership belongs to the specific MCP process that started the
work.

The full-screen CLI application and the inline setup flow are separate terminal
programs. When setup is opened from inside the application, the application
releases the terminal, runs setup, reloads affected state, and then resumes the
full-screen application.

## Application Architecture

The interactive CLI and MCP server are presentation adapters over one shared,
typed application-service layer. Catalog navigation, search, page reads, HTTP
calls, sessions, cache operations, configuration, and lifecycle behavior are
implemented once in transport-neutral services.

The CLI and MCP server do not call each other. The CLI uses a human-oriented
full-screen renderer, while MCP handlers use the agent-oriented YAML-frontmatter
and Markdown renderer defined by this contract. Neither renderer owns business
logic, persistence, network policy, pagination, or validation. Service results
and domain errors are typed so both adapters preserve the same behavior while
presenting it appropriately for their audience.

### Confirmed Implementation Defaults

The first release uses the following local-first implementation defaults:

- One CGO-free Go binary is built for Linux, macOS, and Windows on amd64 and
  arm64.
- Project-maintained canonical Markdown is published as independently
  downloadable, content-addressed ZIP packs. The binary embeds no official API
  catalog. Verified archives and their active state live under
  `~/.apis-mcp/packs`.
- The setup flow refreshes the remote catalog on entry. Pack selection supports
  one, two, or three columns according to terminal width, and stages all changes
  until the final apply step. Failed downloads or index builds leave the prior
  active set unchanged.
- User document sets live under `~/.apis-mcp/library`. A user-defined API family
  replaces the matching official family as one unit; unrelated official packs
  remain available. Duplicate identities at the same precedence fail validation
  with all source locations reported.
- Each version directory is self-contained. Its `_index.md` repeats API-family
  fields, and all versions of one family must agree on family name,
  description, and collections.
- Page frontmatter requires `title`. `path` defaults to the page's relative
  parent directory, and `page_id` defaults to a title slug. Stable short hashes
  of relative file paths resolve collisions; explicit `path` and `page_id`
  values are accepted after validation.
- The server fingerprints installed packs and user canonical files and rebuilds
  changed sources at startup. The interactive app and CLI also expose Markdown,
  OpenAPI, `llms.txt`, and static HTML imports plus explicit rebuild operations.
  The first release does not add MCP tools that modify the library.
- Initial source adapters support raw Markdown trees, OpenAPI documents,
  Swagger/Redoc pages exposing an OpenAPI document, `llms.txt`, and ordinary
  static HTML. JavaScript browser automation is not bundled. OpenAPI imports
  produce an overview plus operation-oriented pages and reusable schema pages.
- Canonical Markdown remains authoritative. A pure-Go SQLite database stores
  discardable metadata and FTS search indexes. A complete new generation is
  built beside the active database and atomically published.
- Process configuration is loaded once at startup from
  `~/.apis-mcp/config.toml`. Schema-affecting changes require restarting the
  affected AI client.
- Each MCP process pins one published library generation. Rebuilds become
  visible to new MCP processes and newly opened CLI operations, preserving
  pagination within a process.
- MCP results use text content containing YAML frontmatter and Markdown. The
  first release does not duplicate results through MCP `structuredContent`.
- Cookie sessions use server-generated IDs without a separate naming layer.
- The full CLI application is accompanied by one-shot commands for library
  management, documentation navigation/search/read, HTTP calls, sessions,
  cache management, configuration, installation, diagnostics, and version
  reporting.
- The installer begins with the client registry and config formats supported by
  `favro-mcp`, while using atomic writes, backups, and explicit partial-failure
  status. Re-running the installer is the binary update mechanism; there is no
  automatic background updater. A complete replacement binary is staged before
  the installer sends a bounded, best-effort stop request to any legacy daemon.
- Official pack sources record source and retrieval metadata and are published
  only when the project can redistribute them. Users are responsible for custom
  imports.
- Shutdown stops new work, cancels foreground calls, gives process-owned
  background downloads a short bounded drain, publishes cancellation errors,
  and exits without detached workers.
- Hard implementation ceilings are 1 MiB of rendered MCP output, 1 MiB for a
  header file, 50 MiB for a request payload, and 20 redirects. Lower
  configurable response and token budgets still apply.

## MCP Surface

The server exposes seven direct, typed MCP tools:

- `apis_collections`: list documentation collections.
- `apis_list`: list and filter API documentation sets.
- `apis_pages`: navigate pages and category paths in one documentation set.
- `apis_search`: search one documentation set.
- `apis_read`: read a documentation page or line range.
- `apis_call`: execute an HTTP request.
- `apis_sessions`: discover and manage persistent HTTP cookie sessions.

Tool results use YAML frontmatter for machine-readable metadata followed by a
Markdown body. Tool failures set MCP `isError` and use error metadata in YAML
frontmatter followed by an actionable Markdown explanation.

Each tool exposes its own Go-backed input schema. There is no shared action
discriminator, nested operation envelope, or loose `args` object.

### Input Shape

```json
{"page": 1}
```

The preceding input calls `apis_collections`.

```json
{
  "name": "linkedin",
  "version": "202505",
  "collection": "social_media",
  "page": 1
}
```

The preceding input calls `apis_list`.

```json
{
  "doc_id": "linkedin-marketing-api-202505",
  "path": "posts/creation",
  "page": 1
}
```

The preceding input calls `apis_pages`.

```json
{
  "doc_id": "linkedin-marketing-api-202505",
  "query": "create an organization post",
  "path": "posts",
  "page": 1
}
```

The preceding input calls `apis_search`.

```json
{
  "doc_id": "linkedin-marketing-api-202505",
  "page_id": "create-a-post",
  "lines": [120, 180]
}
```

The preceding input calls `apis_read`.

No tool accepts a caller-controlled page-size or result-limit argument.
Response budgets are configured by the application. Pagination uses an optional
one-based `page` argument.

### `apis_collections`

Arguments, in schema order:

| Argument | Required | Meaning |
|---|---:|---|
| `page` | no | One-based result page; defaults to `1`. |

Each result contains a stable collection identifier, display name, optional
description, and number of API families in the collection. When another result
page exists, the Markdown body includes an exact `apis_collections` follow-up
call.

```yaml
page: 1
total: 1
total_pages: 1
collections:
  - collection: social_media
    name: Social Media
    description: Social network, publishing, and advertising APIs.
    api_count: 4
```

`collection` is the stable slug passed to `apis_list`. `description` is included
only when supplied by collection metadata; it is not generated from member APIs.

### `apis_list`

Arguments, in schema order:

| Argument | Required | Meaning |
|---|---:|---|
| `name` | no | Case-insensitive substring match against API names. |
| `version` | no | Case-insensitive exact match against an available version. |
| `collection` | no | Case-insensitive exact match against a collection slug returned by `apis_collections`. |
| `page` | no | One-based result page; defaults to `1`. |

The result lists API families in the same general style that Sana lists
meetings. Versions are grouped inline under each API family. Every version has
its own `doc_id`, generated from the API name and version.

Example frontmatter:

```yaml
page: 1
total: 1
total_pages: 1
apis:
  - name: LinkedIn Marketing API
    description: Create and manage LinkedIn marketing content.
    collections:
      - social_media
    versions:
      - version: "202505"
        doc_id: linkedin-marketing-api-202505
        pages: 84
      - version: "202401"
        doc_id: linkedin-marketing-api-202401
        pages: 77
```

When another result page exists, the Markdown body includes an exact follow-up
call that preserves `name` and `version`.

### `apis_pages`

Arguments, in schema order:

| Argument | Required | Meaning |
|---|---:|---|
| `doc_id` | yes | Unique API documentation version returned by `apis_list`. |
| `path` | no | Exact documentation category path to browse. |
| `page` | no | One-based result page; defaults to `1`. |

Documentation paths are relative to the configured root of one documentation
version. They are not API endpoint paths and do not repeat the API version.

For example, with documentation root:

```text
https://docs.example.com/api/v2/
```

the page URL:

```text
https://docs.example.com/api/v2/level1/level2/page
```

has path `level1/level2`. The `api/v2` portion belongs to source and version
discovery, not page navigation.

Without `path`, the result contains pages at the documentation root and direct
top-level child paths. With `path`, it contains pages directly in that exact
path and direct child paths only. It never returns sibling paths. Every child
path includes the total number of pages nested recursively beneath it.

Results are ordered with direct child paths first, followed by pages directly in
the selected path. Pagination applies to this combined ordered sequence, so
navigation categories are exposed before leaf pages.

Frontmatter represents the two record types in separate `paths` and `pages`
arrays. A response may contain one or both arrays depending on where the current
pagination window falls, but the logical pagination order remains paths first.

```yaml
doc_id: linkedin-marketing-api-202505
path: posts
page: 1
total: 14
total_pages: 2
paths:
  - path: posts/creation
    nested_pages: 8
  - path: posts/permissions
    nested_pages: 3
pages:
  - page_id: posts-overview
    title: Posts overview
    description: Concepts and requirements for publishing posts.
```

Each page record contains `page_id` and `title`. It includes `description` only
when explicit source metadata supplies one, such as OpenAPI summary/description,
Markdown frontmatter, an `llms.txt` link description, or explicit HTML metadata.
The scraper does not generate descriptions or infer them from the first body
paragraph.

When child paths are present, the Markdown body explains how to call
`apis_pages` with a returned path. When another result page exists, it also
includes an exact follow-up call preserving `doc_id` and `path`.

### `apis_search`

Arguments, in schema order:

| Argument | Required | Meaning |
|---|---:|---|
| `doc_id` | yes | API documentation version to search. |
| `query` | yes | Text to find in that documentation. |
| `path` | no | Restrict results to this documentation path and its descendants. |
| `page` | no | One-based result page; defaults to `1`. |

The required `doc_id` and `query` properties appear first in the schema.

The query language accepts ordinary terms and quoted phrases. It does not expose
backend-specific Boolean operators, field queries, exclusions, fuzzy matching,
wildcards, or prefix syntax. Search terms use case-insensitive whole-token text
matching before relevance ranking.

A result is eligible when it matches at least one query term or quoted phrase.
Ranking uses text relevance plus a coordination boost: matching every clause
receives the strongest boost, and matching more distinct clauses outranks an
otherwise similar partial match. Backend numeric scores remain internal and are
not included in search output.

Exact matches in page title, `page_id`, extracted API endpoint, and operation ID
receive the strongest field boost. Heading and documentation-path matches
receive a medium boost. Body prose, parameter descriptions, and examples use
normal text weight.

Search results follow Sana's search/read workflow. A body hit represents one
matching source line and includes:

- `page_id`
- exact matching line number
- page title
- documentation path
- a short query-centered snippet from the matching line

A metadata-only hit is returned when page title, `page_id`, documentation path,
extracted API endpoint, or operation ID matches even though that value is absent
from the readable page body. It includes `page_id`, title, path, the matched
metadata field, and a query-centered metadata snippet; it omits `line`.
Metadata-only hits remain eligible for the same relevance ranking and
pagination as body hits.

The snippet collapses whitespace. It contains the first matching query clause
in input order with up to 80 characters of context on each side, adding `...`
when text was omitted. If no clause can be located after normalization, it uses
the first 160 characters of the matching line. Neighboring source lines are not
included in search output; the caller uses `apis_read` when it needs them.

Search relevance is used internally but its backend-specific numeric score is
not exposed. Equal-ranked hits use stable page and line identifiers as
tie-breakers so pagination remains deterministic.

The Markdown body starts with Sana-style result accounting:

- `Showing <total> matching lines for "<query>" (keyword, ranked by relevance).`
  when every hit fits on the current page
- `Showing <shown> out of <total> matching lines for "<query>" (keyword, ranked by relevance).`
  when the result set is paginated
- `No documentation lines match "<query>".` when there are no hits
- `No results on page <page> (<total> match(es); <pages> page(s)).` when the
  query has hits but the requested page is beyond them

Hits are rendered as a table with columns `page_id`, `line`, `title`, `path`,
`match`, and `snippet`. The `line` cell is empty for metadata-only hits;
`match` is `body` for a source-line hit or the metadata field name for a
metadata-only hit. When more hits exist, the body includes an exact next-page
call that preserves `doc_id`, `query`, and `path`.

When a result page contains body hits, it includes this guidance:

```text
Read around a hit with apis_read({"doc_id":"<doc_id>","page_id":"<page_id>","lines":[<line>-2,<line>+2]}).
```

When it contains metadata-only hits, it also includes this guidance:

```text
Read a metadata-only hit with apis_read({"doc_id":"<doc_id>","page_id":"<page_id>"}).
```

### `apis_read`

Arguments, in schema order:

| Argument | Required | Meaning |
|---|---:|---|
| `doc_id` | yes | API documentation version containing the page. |
| `page_id` | yes | Page identifier returned by `apis_pages` or `apis_search`. |
| `lines` | no | Inclusive one-based `[start, end]` line range. |

`path` is not accepted by `apis_read`. A `page_id` is unique within its `doc_id`;
deterministic suffixes resolve title-slug collisions.

Without `lines`, `apis_read` starts at line 1 and returns as much of the page as fits
the configured response budget. With `lines`, it starts at the requested line
and returns through the requested end or the response budget, whichever comes
first. If cropped, the result includes the exact continuation range.

The page fragment is returned verbatim inside a tilde-based Markdown fence. The
fence uses at least three tildes and grows when necessary so fences inside the
stored page cannot close it. Guidance is placed after the closing fence.

Example:

````markdown
---
doc_id: linkedin-marketing-api-202505
title: Create a post
page_id: create-a-post
path: posts/creation
lines: [1, 200]
total_lines: 438
truncated: true
---

~~~markdown
## Create a post

```json
{"author":"urn:li:organization:123"}
```
~~~

Output was cropped. Continue with `apis_read({"doc_id":"linkedin-marketing-api-202505","page_id":"create-a-post","lines":[201,400]})`.
````

Line numbers are represented by frontmatter and are not inserted into the
verbatim page body.

## Documentation Data Model

### Canonical Library And Rebuilds

Canonical Markdown files are the durable source of truth for the documentation
library. Search indexes, navigation trees, line maps, page counts, token
estimates, and any database representation are derived artifacts that can be
deleted and rebuilt without losing library content.

A documentation version is fully rebuildable from its API name, version, and
Markdown page files with YAML frontmatter. The rebuild derives `doc_id`, page
identities, hierarchy, navigation counts, searchable fields, stable line
numbers, and all indexes. API descriptions, collection membership, source-root
metadata, and extracted endpoint metadata enrich the result but are not
required for a functional document set.

Each document-version root contains a reserved `_index.md` manifest. Its YAML
frontmatter requires `name` and `version` and may contain `description`,
`collections`, `source_root`, and source/provenance metadata. Its body is
optional package documentation and is not exposed as an API documentation page.
Because the manifest is Markdown, a complete document set remains a portable,
Markdown-only directory tree.

Built-in documentation and user-added documentation use the same canonical
format and ingestion pipeline. Project-maintained source adapters may exploit
known Swagger/OpenAPI generators, documentation frameworks, `llms.txt`, or
stable site layouts, but their output is ordinary canonical Markdown rather
than a private database format. User-facing import and scraping commands call
the same adapters and validators.

Project-maintained canonical files live in the repository. User and future-tool
additions are accepted through a per-user library root. Adding a complete
document-set directory and requesting a rebuild is sufficient to expand the
installed library; users do not edit a database or search index directly.

Rebuilds validate every document set before publishing derived state. A failed
or interrupted rebuild leaves the previously published library usable. A
successful rebuild publishes the new derived state atomically.

### API Family

An API family groups versions for `apis_list` presentation.
Collection membership is many-to-many at the API-family level. Every version of
an API family inherits the family's collections; membership is not repeated or
allowed to diverge per documentation version.

```yaml
name: LinkedIn Marketing API
description: Create and manage LinkedIn marketing content.
collections:
  - social_media
versions:
  - "202505"
  - "202401"
```

### Documentation Version

Each independently searchable version is a document set.

```yaml
doc_id: linkedin-marketing-api-202505
name: LinkedIn Marketing API
version: "202505"
description: Create and manage LinkedIn marketing content.
collections:
  - social_media
source_root: https://docs.example.com/api/v2/
```

`doc_id` is the deterministic slug of name and version. Different API versions
are separate document sets, not paths within one document set.

### Documentation Page

Pages are persisted as Markdown with YAML frontmatter.

```markdown
---
title: Create a post
page_id: create-a-post
path: posts/creation
source: https://docs.example.com/api/v2/posts/creation/create
---

## Create a post

Page content...
```

Fields:

| Field | Meaning |
|---|---|
| `title` | Human-readable page title. |
| `page_id` | Slug of the title, unique within `doc_id`. |
| `path` | Arbitrarily deep documentation category path below `source_root`. |
| `description` | Optional description copied from explicit source metadata. |
| `source` | Canonical source URL or source repository location. |
| `http_methods` | Optional methods extracted from structured API documentation. |
| `api_endpoints` | Optional endpoint paths extracted for exact search boosting. |
| `operation_ids` | Optional operation identifiers extracted for exact search boosting. |
| body | Normalized Markdown used for search and verbatim reads. |

### Search Hit

```yaml
page_id: create-a-post
line: 142
title: Create a post
path: posts/creation
match: body
snippet: Creates a post authored by an organization.
```

Metadata-only example:

```yaml
page_id: create-a-post
title: Create a post
path: posts/creation
match: operation_id
snippet: createOrganizationPost
```

## Pagination And Response Budgets

Numeric pagination is exposed through `page`; callers do not control page size.
The application packs complete result records until the configured approximate
token budget is reached. Consequently, different result pages may contain
different numbers of records.

The default raw-data budgets are:

| Output class | Default approximate budget |
|---|---:|
| `apis_collections`, `apis_list`, `apis_pages`, `apis_search`, and `apis_sessions` | 2,000 tokens |
| `apis_read` page fragments and `apis_call` response previews | 4,000 tokens |

Token estimates use OpenAI's `o200k_base` encoding internally. This
implementation detail is not shown in the installer or other user-facing text.
The configured budget applies only to retrieved information emitted by the
tool. Frontmatter, fences, continuation instructions, and other guidance are
outside the configured budget.

Read output uses the same budget but splits Markdown at a sensible boundary near
the target. Preferred boundaries are headings, blocks, paragraphs, list items,
table rows, code-block lines, and finally physical lines. The returned line range
always describes the actual verbatim source fragment.

An independent byte ceiling still protects the MCP transport and process from
pathological input. Its value remains an implementation decision and is not an
agent-facing argument.

## MCP Tool Configuration UX

Installation does not ask separate questions for pagination, response size,
cache retention, redirect limits, or agent override policy. It shows one summary
of the current MCP tool configuration:

```text
MCP tool configuration
Recommended defaults

Read output                 ~4000 tokens
List output                 ~2000 tokens
Response size limit          50 MB
Agent size-limit override    Enabled
Free disk reserve            20 GB
Saved call retention         1 day
Maximum redirects            5

> Continue
  Change settings
```

`Recommended defaults` is shown only when every displayed value equals its
default. `Continue` is selected by default and advances without changing the
configuration. Pressing Down selects `Change settings`; Enter opens the MCP tool
configuration editor. The same editor is available later from the `apis-mcp`
CLI without rerunning installation.

The editor provides fields or choices for all displayed settings. Changes are
validated and saved together, then the UI returns to the configuration summary.
Invalid values keep the editor open with an inline validation message and do not
partially save the configuration.

The summary and editor show this sentence at the bottom in muted gray:

```text
Limits are approximate and apply only to retrieved information.
```

User-facing text does not name the tokenizer and does not enumerate content
outside the estimate.

### Restore Defaults

Whenever the effective MCP tool configuration differs from the recommended
defaults, configurer screens also show `r Restore defaults`. This includes
settings loaded from a prior run and unsaved changes made during the current
configuration session.

Pressing `r` opens a confirmation screen that lists only values that would
change, showing each current value and its recommended default:

```text
Restore recommended defaults?

Response size limit         100 MB -> 50 MB
Saved call retention         1 week -> 1 day

> No
  Yes
```

`No` is selected by default. Confirming `Yes` restores all MCP tool settings
atomically and returns to the summary, which again displays
`Recommended defaults`. Selecting `No` returns without changing anything.
Restore defaults affects only MCP tool configuration; it does not alter
credentials, installed-client selection, downloaded documentation, or cache
contents.

## `apis_call`

The arguments are:

| Argument | Required | Meaning |
|---|---:|---|
| `method` | yes | Case-sensitive HTTP method token. Every valid token except `CONNECT` is accepted. |
| `endpoint` | yes | Complete absolute HTTP or HTTPS URL, including scheme and host. |
| `headers` | no | Inline JSON object, or a local JSON file path string. |
| `payload` | no | Inline JSON object/array, or a local JSON file path string. |
| `timeout` | no | Positive response-header timeout in seconds; defaults to `30`. |
| `retries` | no | Non-negative retry count. The inferred default is `3` for retry-safe requests and `0` otherwise. |
| `json_path` | no | RFC 9535 JSONPath selecting the JSON value shown in the response preview. |
| `session` | no | Server-generated session ID from a prior call. Omit to create a new session. |
| `allow_large_download` | conditional | Bypass the configured response-size cap for this call. Registered only when installer policy permits agent overrides. |

`headers` and `payload` overload their source representation:

- A `headers` object supplies inline header names and values.
- A `headers` string is always interpreted as a local JSON file path; the file
  must contain a JSON object.
- Each header object value must be either a string or an array of strings. Arrays
  preserve repeated values for one field name; top-level header arrays and
  `{name, value}` records are not accepted.
- A `payload` object or array supplies the inline JSON request body.
- A `payload` string is always interpreted as a local JSON file path; the file
  must contain a JSON object or array.
- Inline string, number, boolean, and null payloads are not accepted.

File paths may be absolute or relative. Inputs are distinguished by JSON type,
not filename syntax or file existence. When `payload` is present and the entire
`headers` argument is absent, the client adds `Content-Type: application/json`.
If any headers object or file is supplied, the client does not infer or add a
missing content type. A response note explicitly reports when the content type
was added automatically.

`endpoint` is never resolved from documentation `path` metadata. Documentation
paths describe the hierarchy of scraped pages and are unrelated to REST API URL
paths. Relative URLs, hostless endpoints, and inference from previous tool calls
are not accepted.

HTTP URL user information is accepted and passed to Go's standard HTTP
transport without custom conversion or conflict handling. An explicitly
supplied `Authorization` header follows the transport's standard precedence.
Complete URLs, including user information and query strings, are preserved in
tool output and cache metadata.

### Local Trust Model

`apis-mcp` is a local tool acting with the authority of the user who runs it.
The user is responsible for endpoints, local input files, credentials, request
effects, and downloaded content selected through the CLI or an installed MCP
client. The server does not implement a remote-service trust boundary or try to
constrain a trusted local agent with SSRF-style policy machinery.

`apis_call` accepts HTTP and HTTPS URLs to arbitrary public, loopback, private,
link-local, and metadata destinations on any valid TCP port. Origins do not
need to appear in the documentation catalog. Hostnames and redirects use Go's
ordinary HTTP client behavior; there are no destination allowlists, address-
range blocks, DNS-rebinding checks, or separate port policies.

Header and payload file arguments may use absolute or relative paths and follow
normal operating-system path and symlink behavior. Relative paths resolve from
the MCP process working directory. The application does not maintain approved
filesystem roots.

### Proxy Behavior

The client honors `HTTP_PROXY`, `HTTPS_PROXY`, and `NO_PROXY` using Go's
standard proxy environment behavior. Proxy selection is not controlled by an
`apis_call` argument.

### TLS Verification

TLS certificate verification is enabled by default and cannot be changed by an
`apis_call` argument. The user may disable verification through standalone
Advanced settings.

`method` is validated using HTTP's method-token syntax and forwarded without
case normalization. The tool accepts registered, future, and API-specific
extension methods. `CONNECT` is rejected because it establishes a bidirectional
tunnel rather than completing an ordinary HTTP request/response exchange, and
its authority-form target is incompatible with the required absolute endpoint.

### Response Header Timeout

`timeout` controls how long each request attempt waits for response headers. It
is a positive number of seconds and defaults to `30` when omitted. It does not
limit a body that is actively downloading; the separately configured background
transition starts only after headers have arrived and response caching has been
prepared.

The recommended user-configured maximum is 600 seconds (10 minutes). Agents may
request any whole-second value from `1` through that configured ceiling. The
registered MCP input schema exposes the current ceiling as the `maximum` for
`timeout`. A larger request is rejected with guidance that the user can raise
the maximum through the standalone `apis-mcp` configurer. The configured maximum
cannot be lower than the 30-second omission default.

If an attempt reaches its response-header timeout, no response body or final
cache path is claimed. When the request ultimately fails for this reason, the
tool returns an actionable timeout error that shows the effective timeout and
suggests retrying `apis_call` with a larger `timeout` value.

### Retries

When `retries` is omitted, the client uses three retries for methods marked
idempotent in its pinned IANA HTTP method registry and for requests containing
an `Idempotency-Key` header. Non-idempotent methods and unknown/custom methods
default to zero retries. An explicit non-negative `retries` value overrides the
inferred default and is treated as caller authorization to retry a request even
when duplicate side effects are possible.

The three default retry delays are 5, 15, and 30 seconds. Retries apply only to
classified transient transport failures, response-header timeouts, and selected
transient HTTP statuses such as 408, 425, 429, 500, 502, 503, and 504. Permanent
DNS failures, TLS verification failures, policy denials, malformed requests,
and local file errors are not retried.

Attempt metadata records the reason and delay for each retry. If all attempts
fail, the tool returns an actionable error showing the attempts made and notes
that retry count can be changed with `retries`.

The recommended maximum retry count is `30`. `retries` accepts integers from
zero through the configured maximum: zero explicitly disables retries, while
positive values request that many attempts after the initial request. Delays are
5, 15, 30, and 60 seconds, followed by 120 seconds for every remaining retry.
The tool description warns that large values can keep one call active for a long
time.

The maximum retry count is configurable through the standalone `apis-mcp`
configurer but is not shown or editable during installation. The registered MCP
input schema sets the configured value as the `maximum` for `retries`. Requests
above it are rejected with guidance that the user can change the ceiling through
the CLI configurer.

When a transient response contains a valid `Retry-After`, that value replaces
the normal scheduled delay when it is 120 seconds or less. If the server asks
for a longer wait, automatic retries stop and the tool returns the response with
guidance that reports the requested wait time. The client does not silently cap
the delay and retry earlier than the server requested.

Every received HTTP response is returned as a normal MCP tool result, including
`4xx` and `5xx` statuses after retries are exhausted. Status, headers, body
preview, cache paths, and retry history remain available for inspection. MCP
`isError` is reserved for failures that prevent a valid HTTP response, including
input validation, response-header timeout exhaustion, transport failure, local
input-file failure, and cache publication failure.

### Redirects

The HTTP client follows at most five redirects by default. Every hop repeats the
HTTP/HTTPS URL validation. The result metadata records the redirect chain and
final URL.

Redirect header forwarding uses Go's standard HTTP client behavior.

The redirect limit is an application configuration setting managed through the
`apis-mcp` CLI. It is not an `apis_call` argument and the installer does not ask
the user to adjust it. A configured value of zero disables redirect following.
The exact CLI command and hard upper bound remain implementation decisions.

Redirect method handling follows conventional HTTP client behavior. `301`,
`302`, and `303` can rewrite a non-GET/HEAD request to `GET` and remove its body.
`307` and `308` preserve the original method and body. Each redirect metadata
record includes the source URL, destination URL, status, method before the hop,
method after the hop, and whether the body was retained.

### Saved Responses

Every completed HTTP response is saved to the private tool-call cache. There is
no `save` argument and no uncached mode. The implementation streams response
bytes to disk rather than buffering an unbounded body in memory.

Supported HTTP `Content-Encoding` values are decoded while streaming. The
decoded representation is the cached body used for previews, JSONPath, and
other local tools. Metadata preserves the original encoding and records both
wire bytes and decoded bytes. The configured response-size limit applies to
decoded bytes so compression cannot bypass it. Stacked encodings are decoded in
reverse application order.

If any declared content encoding is unsupported, the body is cached exactly as
received, metadata records `decoded: false`, and no embedded body preview or
JSONPath selection is attempted. The final filename uses an `.encoded` suffix so
agents do not mistake it for decoded content.

The tool result contains a token-capped response preview and identifies the
complete cached body with an absolute local path. YAML frontmatter also reports
the HTTP status, content type, byte count, cache path, and whether the displayed
preview was truncated. Binary responses may omit the embedded preview but still
return their cache path and metadata.

Each cache entry stores the response body separately from a metadata sidecar so
other tools can consume the body directly. HTTP error responses are cached like
successful responses. A transport failure with no response body records only
bounded diagnostic metadata.

Cached files and directories use restrictive permissions and are created only
under the application-owned call-cache root.

### JSONPath Selection

When `json_path` is present, the complete response body is still cached. If the
body is valid JSON, the token-capped embedded preview contains the selected
value instead of the complete response root. Selection uses RFC 9535 JSONPath
semantics only; JavaScript expressions, script evaluation, and implementation
extension functions are not accepted.

If the response is still downloading when the tool returns at the configured
background threshold, background completion caches the full body but does not evaluate `json_path`.
The early result records selection state `skipped` and explains that selection
was not performed because the download continued after the tool call returned.
It does not create or promise a later selection artifact.

The selector does not alter the cached body. Frontmatter records the requested
JSONPath and whether selection succeeded. If the selector is invalid, the body
is not JSON, or no value matches, the HTTP result remains successful at the MCP
level. Frontmatter includes `selection_error` and the full cached response path;
the tool does not silently fall back to an unfiltered preview.

### Call Result Frontmatter

Every `apis_call` result includes YAML frontmatter with the metadata needed to
understand and reuse the response. The stable groups are:

```yaml
request:
  id: 019c8f64-7b31-7b2e-9a61-5f8b987fd542
  session_id: 019c8f64-b550-7bd5-8af2-a1d71e6491be
  method: GET
  endpoint: https://api.example.com/items
  automatic_headers:
    content-type: application/json
response:
  state: complete
  status: 200
  content_type: application/json
  bytes_received: 18243
  duration_ms: 241
cache:
  body_path: /home/user/.apis-mcp/cache/calls/019c8f64-7b31-7b2e-9a61-5f8b987fd542/body.json
  headers_path: /home/user/.apis-mcp/cache/calls/019c8f64-7b31-7b2e-9a61-5f8b987fd542/headers.json
  metadata_path: /home/user/.apis-mcp/cache/calls/019c8f64-7b31-7b2e-9a61-5f8b987fd542/metadata.yaml
  completed_at: 2026-07-25T12:00:00Z
preview:
  truncated: true
  approximate_tokens: 4000
selection:
  json_path: $.items[*].id
  matched: true
```

Fields that do not apply are omitted. A background download uses
`response.state: downloading`, reports progress and ETA fields under `response`,
and replaces `cache.body_path` with `cache.temp_path` plus
`cache.final_path`. A failed JSONPath selection includes
`selection.selection_error` while retaining the cached body path.
When selection is skipped for a background download, selection metadata uses
`state: skipped`.

The Markdown body contains the token-capped tilde-fenced preview when one is
available. Guidance and the cache path appear after the closing fence when the
preview is truncated, binary, still downloading, or selection failed.

Preview rendering depends on the response body:

- Valid JSON is pretty-printed and fenced as `json`.
- Other valid UTF-8 textual content is preserved verbatim and fenced using the
  response media type when a useful Markdown language label is known.
- Binary content has no embedded body preview; the result returns metadata and
  the cached body path.

The preview transformation never modifies the complete cached response body.
The outer tilde fence grows when necessary so response content cannot close it.

Frontmatter includes only useful response headers whose values fit a small
display cap. Candidates include content metadata, redirect and pagination
links, retry and rate-limit information, request/trace identifiers, validators,
and modification timestamps. Long values and other response headers are omitted
from the displayed map rather than truncated ambiguously.

The complete response header map is saved to `headers.json`, with each field
represented as an array of strings so repeated values are preserved. The
absolute file path is returned as `cache.headers_path`. The cached header file
uses the same restrictive permissions as the response body. Display filtering
does not remove fields from this cached file.

### Cookies

Every valid `apis_call` belongs to a persistent cookie session. When `session`
is omitted, the server creates a UUIDv7 session before starting the HTTP request
and returns that ID in result frontmatter. When an existing session ID is
provided, the call reuses it, returns the same ID, and contributes response
cookies to the same jar. Agents select sessions returned by the server; they do
not create or name them.

Matching cookies are sent according to domain, path, expiry, and secure-cookie
rules; response cookies update the jar. Sessions remain available across tool
calls and MCP server restarts. Redirects consult the jar again for each
destination rather than forwarding a prebuilt cookie header.

If the request also supplies an explicit `Cookie` header, jar cookies are loaded
first and valid explicit cookie pairs are merged over them. An explicit pair
replaces jar cookies with the same name for that request; non-conflicting pairs
are appended. Response `Set-Cookie` fields still update the persistent session.

Cookie jar files use restrictive permissions outside the response cache and use
the saved-call retention period as their inactivity TTL, which defaults to
one day. Accepting a call with an existing session refreshes `last_used_at`.
Cleanup skips a session while an `apis_call` holds its lock and removes it only
after the full inactivity period has elapsed. The standalone `apis-mcp` CLI can
list, inspect, and delete cookie sessions. Installation does not ask about or
display cookie sessions. Cookie values are treated as sensitive in ordinary CLI
summaries.

Session IDs must be canonical server-generated UUIDv7 values. A supplied ID that
does not identify an existing session is not treated as a request to create one.

## `apis_sessions`

`apis_sessions` lists, inspects, and removes persistent cookie-jar sessions with three
optional arguments:

| Argument | Required | Meaning |
|---|---:|---|
| `id` | no | Server-generated session identifier to inspect or remove. |
| `delete` | no | When `true`, remove the session identified by `id`; defaults to `false`. |
| `page` | no | One-based page of session or cookie results; defaults to `1`. |

Behavior is selected without a separate action field:

- No arguments lists sessions using token-budgeted pagination.
- `id` lists metadata and a token-budgeted page of cookies for that session.
- `id` with `delete: true` removes the complete session and its persisted jar.
- `delete: true` without `id` returns an actionable validation error.
- `delete: false` behaves the same as omitting `delete`.
- `page` is invalid when `delete: true` because deletion has no result set.

Paginated results include page, total count, total pages, and an exact next-page
call when more entries exist. Session and cookie pages use the configurable list
output budget, defaulting to approximately 2,000 tokens.

Deleting a session does not delete cached API responses created while that
session was active. If an in-flight `apis_call` currently holds the session jar
lock, deletion returns a busy error rather than racing the request. Individual
cookie deletion remains available only through the human CLI.

List output includes session ID, creation/update times, cookie count, and known
domains. Inspect output includes cookie name, value, domain, path, expiry,
secure, HTTP-only, and same-site metadata. Cookie values are shown because the
local user and their MCP client already have access to the session.

### Long Downloads

The handler waits for a response body for up to the configured background
transition threshold, which defaults to 60 seconds. If the body is still
downloading at that point, the download continues as a bounded background task
owned by the MCP server rather than the completed tool-call context. The early
tool result reports:

- state `downloading`
- the partial temporary path
- the expected final path
- bytes received so far
- elapsed time
- measured average download speed
- declared total bytes, when a trustworthy `Content-Length` is available
- approximate remaining minutes, only when total length makes that calculable

The Markdown guidance tells the agent to wait approximately the calculated
duration and check whether the final path exists. When total length is unknown,
the result says that completion time cannot be estimated rather than inventing
an ETA. The temporary file is explicitly marked partial and must not be treated
as a complete response.

The background transition threshold is configurable through the standalone
`apis-mcp` configurer. It is not an `apis_call` argument, is not shown or edited
during installation, and does not cancel an active download.

There is no short body idle timeout. An advanced stalled-download fallback
defaults to one hour without receiving any body bytes. Each successful read
resets that timer, so a slow but progressing download can continue indefinitely.
If the fallback expires, the background task removes the temporary body and
publishes the timeout details at `error_path`.

The stalled-download fallback is user-configurable only in the standalone
`apis-mcp` configurer's advanced settings. It is not an `apis_call` argument and
is not shown or edited during installation.

The final path does not exist while the download is active. Completion flushes
and closes the temporary body, atomically renames it to the final path, and then
records `completed_at`. Cache retention starts when the final path is created,
never when the temporary file was created or when the HTTP request began.

A downloader holds an application lock for its temporary entry. Cleanup skips
locked entries. A process crash releases the lock, allowing the next startup or
periodic cleanup to identify and remove the orphaned temporary body and lock.
Temporary entries do not receive a retention timestamp.

If a background download fails after the tool has returned, the final body path
remains absent. Error metadata is written atomically at the final body path plus
the `.error` extension, for example `<uuid>/body.json.error`. The error artifact records
a bounded error code and message, failure time, and bytes received before
failure. The partial temporary body is removed.

An early background-download result identifies all three relevant paths and
explains their states:

- `temp_path` exists only while the partial download is active.
- `final_path` appears when the complete body is published successfully.
- `error_path` appears instead if the background download fails.

Temporary bodies append `.temp` to the expected final path, for example
`<uuid>/body.json.temp` or `<uuid>/body.encoded.temp`. The background error artifact appends
`.error` to the expected final path. The implementation does not use `.part`.

Each call uses a UUIDv7 as its cache-directory name. The response body inside
that directory uses a safe extension selected from trusted media-type handling;
unknown binary content uses `.bin`, and an unsupported content encoding uses
`.encoded`. `Content-Disposition` and URL filename suggestions are stored only
in metadata and never become local paths.

The agent is instructed to check for `final_path` after the estimated duration
and, if it is absent, check `error_path`. Error-artifact retention starts when
the error path is created.

### Response Size Policy

The default maximum saved response body is 50 MB. Per-call agent overrides are
enabled by default, so recommended configurations expose
`allow_large_download`. These values appear in the MCP tool configuration
summary and are changed only through `Change settings` or the standalone CLI
configurer, not through mandatory installer questions.

When agent overrides are enabled, the MCP tool schema includes optional
`allow_large_download`, defaulting to `false`. Setting it to `true` bypasses the
configured soft cap for that call. When overrides are disabled, the property is
not present in the registered tool schema and runtime input containing it is
rejected.

If a declared `Content-Length` exceeds the cap, the request is stopped before
downloading the body. If total size was not declared and streamed bytes cross
the cap, the download is aborted. When overrides are available, the tool error
explains that the agent may retry with `allow_large_download: true`. When they
are disabled, the error explains that the user can increase the limit or enable
agent overrides through the `apis-mcp` configurer.

If the cap is crossed after the call has already returned into background mode,
the configured-size error is published at `error_path` using the same background
failure contract.

### Disk Space Reserve

The cache filesystem must retain at least 20 GB of free space by default after
an `apis_call` download. This reserve is non-bypassable by
`allow_large_download`. It is shown and configurable in both the installer MCP
tool settings editor and the standalone `apis-mcp` configurer.

Before opening a response body, the server removes cache entries that are
already eligible for retention cleanup and checks available space again. If the
filesystem is already at or below the reserve, the body download does not start.
When a trustworthy body size is known and would cross the reserve, the server
also rejects it before writing.

For responses without a trustworthy final decoded size, the streaming writer
checks available space before bounded writes and stops before it knowingly
crosses the configured reserve. A foreground failure returns an actionable tool
error. A background failure removes the temporary body and publishes the reason
at `error_path`.

The reserve protects against space consumed by `apis-mcp`; unrelated processes
can change free space concurrently, so diagnostics must not claim the tool can
guarantee a system-wide amount against external writers.

### Cache Retention

Available retention values are `1 hour`, `1 day`, `1 week`, and `1 month`.
`1 day` is the recommended default. The value is shown and edited through the
MCP tool configuration flow rather than a mandatory installer prompt. Cleanup
runs at startup and periodically while the MCP server remains active. Final
entries are eligible based on `completed_at`; failures to delete cache entries
are reported to diagnostics without failing an unrelated tool call.

## Confirmed Naming Rules

- Tool names are `apis_collections`, `apis_list`, `apis_pages`, `apis_search`,
  `apis_read`, `apis_call`, and `apis_sessions`.
- Filters are named `name`, `version`, and `path`, not `*_filter`.
- Pagination is `page`, with no agent-facing `limit`.
- API versions are represented by distinct `doc_id` values.
- Documentation hierarchy is represented by `path`.
- A page is selected for reading by `doc_id` and `page_id`.
