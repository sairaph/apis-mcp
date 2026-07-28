---
title: Return WebVTT captions for a provided language
page_id: operation-get-accounts-account-id-stream-identifier-captions-language-vtt-dc67e881
path: operations/stream-subtitles-captions
description: Return WebVTT captions for a provided language.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/captions/{language}/vtt
operation_ids:
    - stream-subtitles/-captions-get-vtt-caption-or-subtitle
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Return WebVTT captions for a provided language

`GET /accounts/{account_id}/stream/{identifier}/captions/{language}/vtt`

Operation ID: `stream-subtitles/-captions-get-vtt-caption-or-subtitle`

Return WebVTT captions for a provided language.

## Definition

```yaml
{"operationId": "stream-subtitles/-captions-get-vtt-caption-or-subtitle", "summary": "Return WebVTT captions for a provided language", "description": "Return WebVTT captions for a provided language.", "parameters": [{"name": "language", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_language"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Return WebVTT caption or subtitle response.", "content": {"text/vtt": {"schema": {"type": "string", "example": "'WEBVTT\n 00:00:00.000 --> 00:00:02.480\n This is example response'\n"}}}}, "4XX": {"description": "Return WebVTT caption or subtitle response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Subtitles/Captions"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.captions.language.vtt", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
