---
title: POST /paas/v4/web_search
page_id: operation-post-paas-v4-web-search-4e6d2599
path: operations/untagged
description: The [Web Search](/guides/tools/web-search) is a specialized search engine for large language models. Building upon traditional search engine capabilities like web crawling and ranking, it enhances intent recognition to return results better suited for LLM processing (including webpage titles, URLs, summaries, site names, favicons etc.).
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/web_search
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# POST /paas/v4/web_search

`POST /paas/v4/web_search`

The [Web Search](/guides/tools/web-search) is a specialized search engine for large language models. Building upon traditional search engine capabilities like web crawling and ranking, it enhances intent recognition to return results better suited for LLM processing (including webpage titles, URLs, summaries, site names, favicons etc.).

## Definition

```yaml
{"description": "The [Web Search](/guides/tools/web-search) is a specialized search engine for large language models. Building upon traditional search engine capabilities like web crawling and ranking, it enhances intent recognition to return results better suited for LLM processing (including webpage titles, URLs, summaries, site names, favicons etc.).", "parameters": [{"$ref": "#/components/parameters/AcceptLanguage"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/WebSearchRequest"}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/WebSearchResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
