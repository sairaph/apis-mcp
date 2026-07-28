---
title: Web Reader
page_id: operation-post-paas-v4-reader-cb1ade84
path: operations/tools-api
description: Reads and parses the content of the specified URL. Supports selectable return formats, cache control, image retention, and summary options.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/reader
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# Web Reader

`POST /paas/v4/reader`

Reads and parses the content of the specified URL. Supports selectable return formats, cache control, image retention, and summary options.

## Definition

```yaml
{"tags": ["Tools API"], "summary": "Web Reader", "description": "Reads and parses the content of the specified URL. Supports selectable return formats, cache control, image retention, and summary options.", "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/ReaderRequest"}, "examples": {"Basic": {"value": {"url": "https://www.example.com"}}}}}, "required": true}, "responses": {"200": {"description": "Processing successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/ReaderResponse"}}}}, "default": {"description": "The request has failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
