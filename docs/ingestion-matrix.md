# Ingestion Matrix

This matrix tracks framework detection and format ingestion independently. A
framework is complete only after both public deployments pass detection,
canonical Markdown generation, SQLite rebuilding, and indexed search. Live
sites are acceptance checks; deterministic local fixtures remain the CI source
of truth because external sites change without notice.

## Acceptance Workflow

1. Detect the framework from the supplied URL before choosing an engine.
2. Exercise the development deployment and inspect generated Markdown.
3. Rebuild the SQLite generation and query content from the imported pages.
4. Exercise the independent test deployment and inspect errors and output.
5. Retain the detector, scraper profile, local fixture, tests, and live URLs
   only when both deployments pass.

Status values are `planned`, `active`, and `complete`.

## Frameworks

| Framework | Development deployment | Test deployment | Engine | Status |
| --- | --- | --- | --- | --- |
| Docusaurus | https://docusaurus.io/docs | https://jestjs.io/docs/getting-started | Static HTML | complete |
| MkDocs Material | https://squidfunk.github.io/mkdocs-material/getting-started/ | https://docs.pydantic.dev/latest/ | Static HTML | planned |
| MkDocs | https://www.mkdocs.org/getting-started/ | https://click.palletsprojects.com/ | Static HTML | planned |
| Sphinx | https://www.sphinx-doc.org/en/master/ | https://docs.python.org/3/ | Static HTML / source | planned |
| VitePress | https://vitepress.dev/guide/getting-started | https://vuejs.org/guide/introduction.html | Static HTML | planned |
| Nextra | https://nextra.site/docs | https://swr.vercel.app/docs/getting-started | Static HTML | planned |
| Astro Starlight | https://starlight.astro.build/getting-started/ | https://docs.astro.build/en/getting-started/ | Static HTML | planned |
| Docsify | https://docsify.js.org/#/quickstart | https://docsify-this.net/ | Source Markdown | planned |
| mdBook | https://rust-lang.github.io/mdBook/ | https://doc.rust-lang.org/book/ | Static HTML / TOC | planned |
| VuePress | https://v2.vuepress.vuejs.org/guide/ | https://vuex.vuejs.org/ | Static HTML | planned |
| Rspress | https://rspress.rs/guide/start/getting-started | https://rsbuild.dev/guide/start/index | Static HTML / Markdown | planned |
| GitHub repository Markdown | https://github.com/github/docs/tree/main/content | https://github.com/facebook/docusaurus/tree/main/website/docs | Repository tree / source Markdown | planned |

Public URLs are revalidated when work starts on their row. A changed or
retired deployment is replaced with another independent public deployment.

## Formats And Platforms

| Format or platform | Development deployment | Test deployment | Preferred engine | Status |
| --- | --- | --- | --- | --- |
| OpenAPI 3 JSON/YAML | https://demo.netbox.dev/api/schema/ | https://raw.githubusercontent.com/getsentry/sentry-api-schema/main/openapi-derefed.json | OpenAPI | complete |
| Swagger 2 JSON/YAML | https://petstore.swagger.io/v2/swagger.json | https://generator.swagger.io/api/swagger.json | OpenAPI | complete |
| Swagger UI | https://petstore.swagger.io/ | https://demo.netbox.dev/api/schema/swagger-ui/ | OpenAPI discovery | complete |
| Redoc | https://redocly.github.io/redoc/ | https://demo.netbox.dev/api/schema/redoc/ | OpenAPI discovery | planned |
| Scalar | https://docs.scalar.com/swagger-editor | https://galaxy.scalar.com/ | OpenAPI discovery | planned |
| Stoplight Elements | https://elements-demo.stoplight.io/ | https://docs.stoplight.io/ | OpenAPI discovery | planned |
| RapiDoc | https://eu.api.ovh.com/console/ | https://www.goatcounter.com/api2.html | OpenAPI discovery / finite catalog | complete |
| Fern | https://buildwithfern.com/learn | https://docs.cohere.com/ | OpenAPI endpoints / Markdown | planned |
| GitBook | https://gitbook.com/docs | https://docs.gitbook.com/ | Markdown alternate / llms.txt | planned |
| Mintlify | https://www.mintlify.com/docs | https://docs.anthropic.com/ | Markdown alternate / llms.txt | planned |
| ReadMe | https://docs.readme.com/main/docs | https://docs.nvidia.com/ | Page Markdown / llms.txt | planned |
| Read the Docs | https://docs.readthedocs.com/platform/stable/ | https://requests.readthedocs.io/en/latest/ | Markdown negotiation / HTML | planned |
| llms.txt | https://www.mintlify.com/docs/llms.txt | https://gitbook.com/docs/llms.txt | llms.txt | planned |
| Static HTML | https://docusaurus.io/docs | https://www.mkdocs.org/ | Static HTML | active |
| Source Markdown shell | https://docsify.js.org/#/quickstart | https://docsify-this.net/ | Source Markdown | planned |

## Completed Evidence

### Docusaurus

Verified on 2026-07-26 against Docusaurus 3.10.1 deployments:

| Check | Docusaurus development site | Jest test site |
| --- | --- | --- |
| Framework detection | `docusaurus` | `docusaurus` |
| Pages ingested for acceptance | 55 | 29 |
| Ordinary documentation pages | Pass | Pass |
| Generated category pages | Pass | Not encountered in sample |
| Active tab selection | Not encountered in sample | Pass |
| Multiline fenced code | Pass | Pass |
| Framework chrome excluded | Pass | Pass |
| Collections persisted | `documentation_frameworks` | `documentation_frameworks`, `testing` |
| SQLite generation published | Pass | Pass |
| FTS query returned imported pages | Pass | Pass |
| Detached persisted job | Pass | Pass |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

A separate live Docusaurus job was canceled after 26 fetched pages. It reached
the durable `canceled` state and published no destination or SQLite update.

### OpenAPI 3 JSON/YAML

Verified on 2026-07-27 with direct OpenAPI discovery and generation:

| Check | NetBox development schema | Sentry test schema |
| --- | --- | --- |
| Job ID | `019fa070-3cd4-7bf8-87b2-b39835619a7e` | `019fa072-7f27-7dfb-8651-282a16995e27` |
| Detection | `openapi` / `openapi` | `openapi` / `openapi` |
| Canonical Markdown pages | 2,240 | 409 |
| SQLite pages for document | 2,240 | 409 |
| FTS acceptance query | `virtualization` | `issue` |
| Detached terminal state | `succeeded` | `succeeded` |

### Swagger 2 JSON/YAML

Verified on 2026-07-27 with two independent raw Swagger 2 documents:

| Check | Swagger Petstore | Swagger Generator |
| --- | --- | --- |
| Job ID | `019fa074-c240-7041-ad20-cbbb66fbf718` | `019fa074-c2a4-727e-a99f-e7a7b5a8df2d` |
| Framework detection | `swagger` | `swagger` |
| Canonical Markdown pages | 27 | 14 |
| SQLite pages for document | 27 | 14 |
| FTS acceptance query | `inventory` | `clients` |
| Detached terminal state | `succeeded` | `succeeded` |

### Swagger UI

Verified on 2026-07-27 by starting from each rendered Swagger UI shell:

| Check | Swagger Petstore UI | NetBox Swagger UI |
| --- | --- | --- |
| Job ID | `019fa075-7895-780f-8e84-3799c3d34157` | `019fa075-7896-77fe-847d-2a50079c2bff` |
| Detection | `swagger-ui` / `html` | `swagger-ui` / `html` |
| Resolved schema | `https://petstore.swagger.io/v2/swagger.json` | `https://demo.netbox.dev/api/schema/` |
| Resolved schema kind | Swagger 2 | OpenAPI 3 |
| Canonical Markdown pages | 27 | 2,240 |
| SQLite pages for document | 27 | 2,240 |
| FTS acceptance query | `inventory` | `virtualization` |
| Detached terminal state | `succeeded` | `succeeded` |

### RapiDoc

Verified on 2026-07-27 against a customized multi-schema RapiDoc deployment
and an independent stock single-schema deployment:

| Check | OVHcloud API console | GoatCounter API |
| --- | --- | --- |
| Job ID | `019fa116-ed44-7698-82d9-6d241e17bb87` | `019fa116-ed38-78f8-bac7-d88308764fa7` |
| Detection | `rapidoc` / `html` | `rapidoc` / `html` |
| Discovery | Two `spec-roots` catalogs | One standard `spec-url` |
| Specifications exhausted | 88 | 1 |
| Canonical Markdown pages | 15,043 | 44 |
| SQLite pages for document | 15,043 | 44 |
| FTS acceptance query | `vrack` | `hits` |
| Detached terminal state | `succeeded` | `succeeded` |

Stock RapiDoc ingestion resolves every statically advertised `spec-url` from
`rapi-doc` and `rapi-doc-mini` components. OVHcloud's customized `spec-roots`
attribute is handled as a finite vendor catalog: every catalog and every JSON
schema template is resolved to OpenAPI 3, namespaced, fetched, generated, and
indexed in one atomic publication. Empty-path catalog members are retained as
overview/schema documentation. A missing, malformed, repeated, cross-origin,
or unsupported catalog entry fails the complete import.

RapiDoc can also receive an in-memory specification through JavaScript
`loadSpec(object)`. Such a page is detected as RapiDoc but is incomplete unless
it exposes a stable source URL; the importer fails it explicitly instead of
falling back to a generic HTML crawl or claiming completeness.

## Completeness Rules

Framework-specific crawls are complete only when their finite navigation queue
reaches zero without fetch, parse, scope, or indexing errors. Generic HTML is
best effort unless a finite inventory such as a sitemap, `llms.txt`, or known
framework navigation model is discovered and exhausted. A generic crawl must
not claim completeness merely because no more anchor links were found.

GitHub repository Markdown is treated as a repository source rather than as
rendered HTML. Detection must resolve the owner, repository, ref, and optional
subtree from a GitHub URL. Ingestion must enumerate the selected Git tree,
reject GitHub's `truncated` tree response, fetch every selected `.md` and `.mdx`
blob, preserve repository-relative paths and links, and verify that discovered,
fetched, generated, and indexed file counts agree. Authentication, rate-limit,
and missing-blob failures make the job incomplete rather than silently reducing
the document set.
