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
| Nextra | https://nextra.site/docs | https://authjs.dev/getting-started | Static HTML / sitemap | complete |
| Astro Starlight | https://starlight.astro.build/getting-started/ | https://docs.astro.build/en/getting-started/ | Static HTML / sitemap index | complete |
| Docsify | https://docsify.js.org/#/quickstart | https://docsify-this.net/?basePath=https%3A%2F%2Fraw.githubusercontent.com%2Fhibbitts-design%2Fdocsify-this-multiple-page-course-site%2Fmain | Source Markdown / Git tree | complete |
| mdBook | https://rust-lang.github.io/mdBook/ | https://doc.rust-lang.org/book/ | Static HTML / TOC | complete |
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
| Source Markdown shell | https://docsify.js.org/#/quickstart | https://docsify-this.net/?basePath=https%3A%2F%2Fraw.githubusercontent.com%2Fhibbitts-design%2Fdocsify-this-multiple-page-course-site%2Fmain | Source Markdown / Git tree | complete |

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
that Material intentionally omits from navigation. Every same-origin HTML page
entry must be fetched successfully and remain under the detected documentation
root; a non-page record invalidates the authoritative sitemap. Static refresh
aliases count as fetched inventory but do not generate duplicate Markdown pages.

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
framework detection must remain consistent across all pages, and non-page
records invalidate the authoritative sitemap.

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

### Nextra

Verified on 2026-07-27 against independent Nextra 4 and classic Nextra 3
deployments. SWR was replaced after it migrated from Nextra to Fumadocs.

| Check | Nextra development site | Auth.js test site |
| --- | --- | --- |
| Framework generation | Nextra 4 | Nextra 3 |
| Job ID | `019fa397-c780-7984-a2f9-7cfab5c064d7` | `019fa397-c780-7aab-a984-52bc013b585e` |
| Framework detection | `nextra` | `nextra` |
| Scoped sitemap entries fetched | 59 | 135 |
| Canonical Markdown pages | 59 | 135 |
| Theme content isolated | Pass | Pass |
| Fenced code languages retained | Pass | Pass |
| SQLite pages for document | 59 | 135 |
| FTS acceptance query | `markdown` | `authentication` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

Nextra completeness is bounded by the non-empty sitemap at the detected Next.js
deployment root, filtered to the documentation section derived from the current
page and its Nextra sidebar. Every route in that scoped inventory must be
fetched, generated, and indexed. Sidebar links outside the current documentation
section do not widen the inventory, and an absent sitemap, ambiguous deployment
root, empty section, or non-static content container fails closed.

### Astro Starlight

Verified on 2026-07-27 against two independent Starlight 0.41 deployments:

| Check | Starlight development site | Astro Docs test site |
| --- | --- | --- |
| Framework version | `0.41.4` | `0.41.0` |
| Astro version | `7.0.2` | `7.0.2` |
| Job ID | `019fa3cc-d45b-7978-b3fa-93566bdc6661` | `019fa3cc-d52b-7783-a0f0-c89db67445df` |
| Framework detection | `astro-starlight` | `astro-starlight` |
| Locale-scoped sitemap entries fetched | 36 | 417 |
| Canonical Markdown pages | 36 | 417 |
| Hero and theme content isolated | Pass | Pass |
| Expressive Code languages retained | Pass | Pass |
| Multi-block Steps retained | Pass | Pass |
| SQLite pages for document | 36 | 417 |
| FTS acceptance query | `sidebar` | `islands` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited crawl completed naturally | Pass | Pass |
| Truncated | No | No |

Astro Starlight completeness is bounded by every URL record in every shard of
the uniquely advertised same-origin sitemap index. For multilingual sites, each
record's `hreflang` alternates identify the selected locale instead of relying
on sidebar coverage or hard-coded locale prefixes. The importer validates every
shard and URL record, requires canonical static Starlight content for every
selected route, and rejects redirects, repeated entries, malformed records,
ambiguous locale metadata, or client-only shells before publication.

### Docsify

Verified on 2026-07-27 against the official Docsify deployment and an
independent Docsify-This deployment backed by a multi-page course repository:

| Check | Docsify development site | Docsify-This test site |
| --- | --- | --- |
| Job ID | `019fa42f-b214-754a-a1ef-6e4beeff83f2` | `019fa42f-b262-7675-b1f1-21bdb3b6e78a` |
| Detection / engine | `docsify` / `docsify` | `docsify` / `docsify` |
| Statically configured GitHub roots | 4 | 1 |
| Git tree `.md` sources exhausted | 49 | 26 |
| Canonical Markdown pages | 49 | 26 |
| Immutable commit sources | Pass | Pass |
| Markdown directory layout retained | Pass | Pass |
| Frontmatter fields and Docsify syntax retained | Pass | Pass |
| SQLite pages for document | 49 | 26 |
| FTS acceptance query | `quickstart` | `schedule` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited inventory completed naturally | Pass | Pass |
| Truncated | No | No |

Docsify completeness is bounded by GitHub-backed source roots that are
statically advertised in the shell configuration or supplied through an
explicit `basePath` query. A standard GitHub Pages deployment is accepted only
when public deployment metadata identifies its successful commit and URL and
exactly one immutable repository shell at the root or under `docs/` matches the
deployed shell. Each ref, including refs containing slashes, is resolved to an
immutable commit; its recursive Git tree must be untruncated, and every selected
`.md` blob must be fetched and published at its mirrored repository path. An
upstream `_index.md` receives a deterministic page filename to avoid collision
with the canonical manifest convention. Generic HTTP, ambiguous Pages roots,
and dynamic `basePath` sources fail closed when they cannot provide a finite
enumerable inventory. The importer preserves source Markdown and non-conflicting
frontmatter fields instead of executing Docsify or scraping its client-rendered
shell. Non-`.md` resources are outside the canonical page inventory and are not
mirrored; each page retains its immutable upstream source for resolving such
references. `GITHUB_TOKEN` is used only for explicit HTTPS `api.github.com`
inventory requests when present, allowing large or repeated inventories to
avoid anonymous rate limits.

### mdBook

Verified on 2026-07-27 against two independent books generated by current
mdBook 0.5 releases:

| Check | mdBook development guide | Rust Book test site |
| --- | --- | --- |
| Framework version | `0.5.4` | `0.5.x` (`0.5.1` project pin) |
| Job ID | `019fa478-7ea6-7971-a476-2f86cdb7560f` | `019fa478-7f1e-7cfc-a6f9-fb3e77e90cae` |
| Detection / engine | `mdbook` / `html` | `mdbook` / `html` |
| Fetchable TOC chapters exhausted | 31 | 111 |
| Root-to-first-chapter alias | `index.html` | `title-page.html` |
| Canonical Markdown pages | 31 | 111 |
| Main content isolated from book chrome | Pass | Pass |
| Heading permalinks excluded | Pass | Pass |
| Rust code languages and hidden lines retained | Pass | Pass |
| SQLite pages for document | 31 | 111 |
| FTS acceptance query | `preprocessors` | `ownership` |
| Detached terminal state | `succeeded` | `succeeded` |
| Unlimited inventory completed naturally | Pass | Pass |
| Truncated | No | No |

mdBook chapter completeness is bounded by the single generated `ol.chapter`
inventory in the book root's `toc.html`. Draft entries without chapter anchors
are intentionally excluded. Every fetchable chapter URL must be unique,
query-free, same-origin, under the derived book root, and successfully rendered
with a consistent mdBook TOC asset and static main content. The root and
`index.html` aliases map to the first TOC chapter without generating duplicate
pages. Print output, copied assets, 404 pages, and unlinked redirect aliases are
not chapter pages and are outside this completeness claim because mdBook does
not publish a finite manifest for them.

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

## Extended Compatibility Sweep

Verified on 2026-07-27 against 30 additional public API and developer
documentation deployments, three per completed framework or framework-backed
format. Every final job completed with an empty queue and no page or depth
truncation. Generated Markdown and independently reopened SQLite page counts
matched for all 2,656 canonical pages, and every document returned relevant FTS
results.

| Framework | Deployment | Final job | Finite inventory accounting | Markdown / SQLite pages |
| --- | --- | --- | --- | ---: |
| Docusaurus | [Playwright Node.js API](https://playwright.dev/docs/api/class-playwright) | `019fa501-9c05-7f28-a0f8-dcece5a53dc6` | 70 unique scoped sidebar routes | 70 / 70 |
| Docusaurus | [Babel documentation](https://babeljs.io/docs/config-files) | `019fa503-1530-780e-9a92-6841b7b097e9` | 102 unique scoped sidebar routes | 102 / 102 |
| Docusaurus | [Prettier API](https://prettier.io/docs/api/) | `019fa504-6a36-7b29-aee6-be8317bd3f6d` | 24 canonical sidebar pages; slash alias deduplicated | 24 / 24 |
| RapiDoc | [Pi-hole FTL API](https://ftl.pi-hole.net/master/docs/) | `019fa567-3e08-79c1-928a-4aa263fe44c6` | 22-file graph; 92 operations and 109 reachable schemas | 202 / 202 |
| RapiDoc | [Mailpit API v1](https://mailpit.axllent.org/docs/api-v1/view.html) | `019fa569-a4cb-7362-b21a-d0b0dfa6ef10` | 1 Swagger source; 25 operations and 19 definitions | 45 / 45 |
| RapiDoc | [OpenAPI Generator Online](https://api.openapi-generator.tech/index.html) | `019fa56a-28b7-7eb5-bbe2-b2681c512407` | 1 Swagger source; 7 operations and 5 definitions | 13 / 13 |
| MkDocs Material | [PyIceberg](https://py.iceberg.apache.org/) | `019fa4b3-7f5b-7590-9238-5429f012a7ba` | 91 sitemap records | 91 / 91 |
| MkDocs Material | [Hatch 1.17](https://hatch.pypa.io/1.17/) | `019fa4b5-5e07-779e-86c8-e367901b9194` | 85 sitemap records | 85 / 85 |
| MkDocs Material | [HTTPX](https://www.python-httpx.org/) | `019fa496-16a7-7087-81b4-3a869d88b834` | 23 sitemap records | 23 / 23 |
| MkDocs | [CCTools](https://cctools.readthedocs.io/en/latest/) | `019fa494-854d-7d12-b7cf-4e1c79212011` | 119 sitemap records | 119 / 119 |
| MkDocs | [LORIS](https://acesloris.readthedocs.io/en/latest/) | `019fa4b4-3cfa-7547-acea-e7f1bb283c04` | 240 sitemap records; root/index alias accounted | 240 / 240 |
| MkDocs | [Sense HAT](https://sense-hat.readthedocs.io/en/latest/) | `019fa497-377e-7d70-a071-c6449d457520` | 4 sitemap records | 4 / 4 |
| Sphinx | [Flask](https://flask.palletsprojects.com/en/stable/) | `019fa493-dc5e-76f8-8726-7621f97aaac3` | 76 titled search-index documents | 76 / 76 |
| Sphinx | [aiohttp](https://docs.aiohttp.org/en/stable/) | `019fa496-2f91-77d0-aaf0-0cc94ad915f1` | 39 titled search-index documents | 39 / 39 |
| Sphinx | [SQLAlchemy 2.0](https://docs.sqlalchemy.org/en/20/) | `019fa4b3-8c63-7738-b0ff-06613b6b30fc` | 180 titled search-index documents | 180 / 180 |
| VitePress | [OpenAPI TypeScript](https://openapi-ts.dev/introduction) | `019fa4b4-5ca0-7d10-bc99-26c41afefbc7` | 55 sitemap records; 2 redirect aliases | 55 / 55 |
| VitePress | [Vue Router API](https://router.vuejs.org/api/) | `019fa4b5-6de2-7734-8160-958ebb656090` | 229 sitemap records; 144 case-normalizing aliases | 229 / 229 |
| VitePress | [Element Plus](https://element-plus.org/en-US/component/overview) | `019fa4b9-7f2a-77b9-a2ab-eb2ff0ac3a48` | 307 multilingual sitemap records | 307 / 307 |
| Nextra | [React Flow API](https://reactflow.dev/api-reference) | `019fa501-bab8-73c6-bce9-0b49e1437ff6` | 330 records: 119 API pages, 205 out of scope, 6 non-pages | 119 / 119 |
| Nextra | [Typia](https://typia.io/docs/) | `019fa4b8-497b-7dbf-b8f3-9483d6599b1f` | 34 docs routes selected from a 71-record sitemap shard | 34 / 34 |
| Nextra | [imgix Rendering API](https://docs.imgix.com/apis/rendering/overview) | `019fa4cc-643a-7244-830a-d13f7a73a938` | 209 English rendering routes selected from 577 records | 209 / 209 |
| Astro Starlight | [sharp](https://sharp.pixelplumbing.com/) | `019fa4b3-e61a-73f6-9656-3484f826efd0` | 119 sitemap records | 119 / 119 |
| Astro Starlight | [Web Monetization](https://webmonetization.org/docs/) | `019fa4cc-65fc-7568-af76-cda1eff79a61` | 39 fetched: 32 Starlight pages and 7 validated non-framework pages | 32 / 32 |
| Astro Starlight | [ScreenshotOne](https://screenshotone.com/docs/getting-started/) | `019fa4b8-762e-713b-8b63-d6b5c68aa72d` | 81 docs routes selected from 530 records | 81 / 81 |
| Docsify | [node-google-spreadsheet](https://theoephraim.github.io/node-google-spreadsheet/#/) | `019fa519-a0c1-72cc-9d81-f9835d5f1098` | 9 immutable deployed-commit `docs/` Markdown sources | 9 / 9 |
| Docsify | [ERC721A](https://chiru-labs.github.io/ERC721A/#/) | `019fa519-a112-766c-8245-8612f1bb951c` | 11 immutable deployed-commit `docs/` Markdown sources | 11 / 11 |
| Docsify | [Polly.JS](https://netflix.github.io/pollyjs/#/) | `019fa519-a112-75a2-a7b3-dbf04db1f109` | 31 deployed `gh-pages` root and 3 advertised `master` Markdown sources | 34 / 34 |
| mdBook | [Servo Book](https://book.servo.org/) | `019fa4b3-b341-78a3-8f4a-34b454088a02` | 58 TOC records: 57 pages and 1 empty-group alias | 57 / 57 |
| mdBook | [Rust on ESP](https://docs.espressif.com/projects/rust/book/) | `019fa494-dd26-7a71-bd05-57215047dae3` | 23 TOC records | 23 / 23 |
| mdBook | [Aya Book](https://aya-rs.dev/book/) | `019fa4b5-6003-760b-a575-5dc249d5d76c` | 24 root, nested, and clean-route TOC records | 24 / 24 |

The unrepaired candidate runs exposed 18 failures that correctly published
nothing and two dangerous successful-but-incomplete imports. Pi-hole initially
published only its root OpenAPI file, and Prettier initially narrowed a
trailing-slash leaf URL to two duplicate aliases. The fixes established the
following additional compatibility contracts:

- external OpenAPI references form a bounded, same-scope graph that is fully
  loaded, validated, bundled, counted, and generated; OpenAPI 3.1
  pointer-based graphs are supported, while external JSON Schema resource
  semantics such as `$id`, `$anchor`, and `$dynamicRef` fail closed;
- transient source GETs receive bounded, cancellation-aware retries while
  permanent failures still fail the atomic import;
- dense generated API pages remain bounded but may contain up to 250,000 HTML
  nodes;
- sitemap indexes, deterministic route aliases, locale alternates, and initial
  locale redirects are explicitly accounted; non-page records are either
  rejected or narrowly validated and excluded according to the framework;
- monolingual Starlight deployments may use a site-wide mixed-layout sitemap
  only when every excluded page is fetched and proves a non-redirecting,
  self-canonical, non-Starlight static page;
- standard GitHub Pages Docsify deployments pin the successful deployment
  commit and accept exactly one matching immutable root or `docs/` shell, while
  retaining explicitly advertised Markdown roots;
- current mdBook root and clean routes are valid TOC records, and an empty
  index-group record is an alias only when it has one unique in-inventory next
  chapter; and
- Docusaurus leaf starts derive a non-root documentation section from static
  sidebar evidence and collapse slash aliases without crossing into unrelated
  site sections.

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
