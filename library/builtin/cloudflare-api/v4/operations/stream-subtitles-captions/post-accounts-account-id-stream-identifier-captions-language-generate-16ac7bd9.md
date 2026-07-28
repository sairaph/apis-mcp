---
title: Generate captions or subtitles for a provided language via AI
page_id: operation-post-accounts-account-id-stream-identifier-captions-language-generate-d86102c3
path: operations/stream-subtitles-captions
description: Generate captions or subtitles for provided language via AI.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/captions/{language}/generate
operation_ids:
    - stream-subtitles/-captions-generate-caption-or-subtitle-for-language
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate captions or subtitles for a provided language via AI

`POST /accounts/{account_id}/stream/{identifier}/captions/{language}/generate`

Operation ID: `stream-subtitles/-captions-generate-caption-or-subtitle-for-language`

Generate captions or subtitles for provided language via AI.

## Definition

```yaml
{"operationId": "stream-subtitles/-captions-generate-caption-or-subtitle-for-language", "summary": "Generate captions or subtitles for a provided language via AI", "description": "Generate captions or subtitles for provided language via AI.", "parameters": [{"name": "language", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_language"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Generate captions or subtitles response for a provided language.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_language_response_single"}}}}, "4XX": {"description": "Generate captions or subtitles response for a provided language.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Subtitles/Captions"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.captions.language", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
