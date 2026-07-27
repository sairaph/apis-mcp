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
| MkDocs Material | https://squidfunk.github.io/mkdocs-material/getting-started/ | https://docs.astral.sh/ruff/ | Static HTML / sitemap | complete |
| MkDocs | https://www.mkdocs.org/getting-started/ | https://python-markdown.github.io/ | Static HTML / sitemap | complete |
| Sphinx | https://www.sphinx-doc.org/en/master/ | https://requests.readthedocs.io/en/latest/ | Static HTML / search index | complete |
| VitePress | https://vitepress.dev/guide/getting-started | https://vite.dev/guide/ | Static HTML / sitemap | complete |
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

### MkDocs Material

Verified on 2026-07-27 against Material 9.7 deployments. The former Pydantic
test deployment was replaced after it migrated to Astro Starlight.

| Check | Material development site | Ruff test site |
| --- | --- | --- |
| Framework version | `9.7.0+insiders-4.53.18` | `9.7.6` |
| Job ID | `019fa19e-ae88-7075-b2f5-28cabd7ad2db` | `019fa19e-ae96-723e-9f82-bb7ee2883f09` |
| Framework detection | `mkdocs-material` | `mkdocs-material` |
| Sitemap entries fetched | 108 | 988 |
| Static redirect aliases | 1 | 0 |
| Canonical Markdown pages | 107 | 988 |
| Material content isolated | Pass | Pass |
| Heading permalinks excluded | Pass | Pass |
| Fenced code languages retained | Pass | Pass |
| SQLite pages for document | 107 | 988 |
| FTS acceptance query | `admonitions` | `pyproject` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

MkDocs Material completeness is based on its generated `sitemap.xml`, not only
the primary sidebar. This covers blog pagination and generated detail pages
that Material intentionally omits from navigation. Every same-origin sitemap
entry must be fetched successfully and remain under the detected documentation
root. Static refresh aliases count as fetched inventory but do not generate
duplicate Markdown pages.

### MkDocs

Verified on 2026-07-27 against an independent built-in-theme deployment and a
custom-theme deployment:

| Check | MkDocs development site | Python-Markdown test site |
| --- | --- | --- |
| Theme | Built-in default | Custom Nature-based theme |
| Job ID | `019fa1be-532f-7495-aaa1-ced9d394986d` | `019fa1c3-db5a-7109-be54-1970f4419652` |
| Framework detection | `mkdocs` | `mkdocs` |
| Finite sitemap entries fetched | 19 | 68 |
| Canonical Markdown pages | 19 | 68 |
| Theme content isolated | Pass | Pass |
| Heading permalinks excluded | Pass | Pass |
| Fenced code retained | Pass | Pass |
| SQLite pages for document | 19 | 68 |
| FTS acceptance query | `configuration` | `extension` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

Plain MkDocs uses a non-empty same-origin sitemap as its authoritative
inventory. An absent or empty sitemap fails closed because primary navigation
can omit generated pages. Every inventory entry must be fetched successfully,
and framework detection must remain consistent across all pages.

### Sphinx

Verified on 2026-07-27 against independent `html` and `dirhtml` deployments:

| Check | Sphinx development site | Requests test site |
| --- | --- | --- |
| Builder | `html` | `dirhtml` |
| Job ID | `019fa1f3-3a48-7e7c-8e6e-9a36b4d45b7b` | `019fa1f3-3a78-79e1-acf1-c8058da2dd48` |
| Framework detection | `sphinx` | `sphinx` |
| Search-index documents fetched | 155 | 15 |
| Canonical Markdown pages | 155 | 15 |
| Theme content isolated | Pass | Pass |
| API definition signatures retained | Pass | Pass |
| Heading permalinks excluded | Pass | Pass |
| Pygments code languages retained | Pass | Pass |
| SQLite pages for document | 155 | 15 |
| FTS acceptance query | `autodoc` | `authentication` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

Sphinx completeness is bounded by the finite `docnames` inventory generated in
`searchindex.js`: every titled source document advertised there must be
fetched, generated, and indexed. Builder-native `FILE_SUFFIX` URLs are used for
fetches while `LINK_SUFFIX`, root, query, and `index` aliases are accounted for
without duplicate pages. Generated utility pages, extension-created pages, and
titleless source documents are outside this inventory boundary. Missing,
legacy, malformed, empty, repeated, or colliding indexes fail closed.

### VitePress

Verified on 2026-07-27 against two independent VitePress 2 deployments:

| Check | VitePress development site | Vite test site |
| --- | --- | --- |
| Job ID | `019fa32a-d41b-7a86-bd66-b6da344c064b` | `019fa345-331d-7793-84a6-59a93c1b0998` |
| Framework detection | `vitepress` | `vitepress` |
| Finite sitemap entries fetched | 272 | 57 |
| Canonical Markdown pages | 272 | 57 |
| Default/custom theme content isolated | Pass | Pass |
| Heading permalinks excluded | Pass | Pass |
| Shiki code languages retained | Pass | Pass |
| SQLite pages for document | 272 | 57 |
| FTS acceptance query | `sitemap` | `plugins` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

VitePress completeness is bounded by the publisher-selected non-empty sitemap,
including transformed, rewritten, and dynamic routes. The importer derives the
deployment base from `vp-icons.css`, accounts for clean and `index` aliases,
and requires every route to expose useful statically rendered content. Vue was
replaced as the test deployment because two sitemap routes are client-only SSR
shells; those now fail closed instead of becoming title-only placeholders.

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
