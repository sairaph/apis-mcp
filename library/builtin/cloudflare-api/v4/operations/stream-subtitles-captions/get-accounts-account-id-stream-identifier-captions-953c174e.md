---
title: List captions or subtitles
page_id: operation-get-accounts-account-id-stream-identifier-captions-cb694c32
path: operations/stream-subtitles-captions
description: Lists the available captions or subtitles for a specific video.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/captions
operation_ids:
    - stream-subtitles/-captions-list-captions-or-subtitles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List captions or subtitles

`GET /accounts/{account_id}/stream/{identifier}/captions`

Operation ID: `stream-subtitles/-captions-list-captions-or-subtitles`

Lists the available captions or subtitles for a specific video.

## Definition

```yaml
{"operationId": "stream-subtitles/-captions-list-captions-or-subtitles", "summary": "List captions or subtitles", "description": "Lists the available captions or subtitles for a specific video.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "List captions or subtitles response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_language_response_collection"}}}}, "4XX": {"description": "List captions or subtitles response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Subtitles/Captions"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.captions", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
