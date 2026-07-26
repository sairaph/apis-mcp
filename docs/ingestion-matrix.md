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

Public URLs are revalidated when work starts on their row. A changed or
retired deployment is replaced with another independent public deployment.

## Formats And Platforms

| Format or platform | Development deployment | Test deployment | Preferred engine | Status |
| --- | --- | --- | --- | --- |
| OpenAPI 3 JSON/YAML | https://demo.netbox.dev/api/schema/ | https://raw.githubusercontent.com/getsentry/sentry-api-schema/main/openapi-derefed.json | OpenAPI | complete |
| Swagger 2 JSON/YAML | https://petstore.swagger.io/ | https://generator.swagger.io/api/swagger.json | OpenAPI | complete |
| Swagger UI | https://petstore.swagger.io/ | https://demo.netbox.dev/api/schema/swagger-ui/ | OpenAPI discovery | complete |
| Redoc | https://redocly.github.io/redoc/ | https://demo.netbox.dev/api/schema/redoc/ | OpenAPI discovery | planned |
| Scalar | https://docs.scalar.com/swagger-editor | https://galaxy.scalar.com/ | OpenAPI discovery | planned |
| Stoplight Elements | https://elements-demo.stoplight.io/ | https://docs.stoplight.io/ | OpenAPI discovery | planned |
| RapiDoc | https://mrin9.github.io/RapiDoc/ | Public test deployment to be selected | OpenAPI discovery | planned |
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
