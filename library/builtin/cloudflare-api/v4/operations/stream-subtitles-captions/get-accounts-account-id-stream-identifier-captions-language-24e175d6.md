---
title: List captions or subtitles for a provided language
page_id: operation-get-accounts-account-id-stream-identifier-captions-language-1daac962
path: operations/stream-subtitles-captions
description: Lists the captions or subtitles for provided language.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/captions/{language}
operation_ids:
    - stream-subtitles/-captions-get-caption-or-subtitle-for-language
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List captions or subtitles for a provided language

`GET /accounts/{account_id}/stream/{identifier}/captions/{language}`

Operation ID: `stream-subtitles/-captions-get-caption-or-subtitle-for-language`

Lists the captions or subtitles for provided language.

## Definition

```yaml
{"operationId": "stream-subtitles/-captions-get-caption-or-subtitle-for-language", "summary": "List captions or subtitles for a provided language", "description": "Lists the captions or subtitles for provided language.", "parameters": [{"name": "language", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_language"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "List captions or subtitles response for a provided language.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_language_response_single"}}}}, "4XX": {"description": "List captions or subtitles response for a provided language.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Subtitles/Captions"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.captions.language", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
