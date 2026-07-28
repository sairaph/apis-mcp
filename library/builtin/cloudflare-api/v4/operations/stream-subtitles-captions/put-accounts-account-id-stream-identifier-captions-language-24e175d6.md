---
title: Upload captions or subtitles
page_id: operation-put-accounts-account-id-stream-identifier-captions-language-cd4be6d9
path: operations/stream-subtitles-captions
description: Uploads the caption or subtitle file to the endpoint for a specific BCP47 language. One caption or subtitle file per language is allowed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/captions/{language}
operation_ids:
    - stream-subtitles/-captions-upload-captions-or-subtitles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload captions or subtitles

`PUT /accounts/{account_id}/stream/{identifier}/captions/{language}`

Operation ID: `stream-subtitles/-captions-upload-captions-or-subtitles`

Uploads the caption or subtitle file to the endpoint for a specific BCP47 language. One caption or subtitle file per language is allowed.

## Definition

```yaml
{"operationId": "stream-subtitles/-captions-upload-captions-or-subtitles", "summary": "Upload captions or subtitles", "description": "Uploads the caption or subtitle file to the endpoint for a specific BCP47 language. One caption or subtitle file per language is allowed.", "parameters": [{"name": "language", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_language"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"$ref": "#/components/schemas/stream_caption_basic_upload"}}}}, "responses": {"200": {"description": "Upload captions or subtitles response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_language_response_single"}}}}, "4XX": {"description": "Upload captions or subtitles response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Subtitles/Captions"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.captions.language", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
