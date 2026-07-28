---
title: Create signed URL tokens for videos
page_id: operation-post-accounts-account-id-stream-identifier-token-e3b4b148
path: operations/stream-videos
description: Creates a signed URL token for a video. If a body is not provided in the request, a token is created with default values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/token
operation_ids:
    - stream-videos-create-signed-url-tokens-for-videos
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create signed URL tokens for videos

`POST /accounts/{account_id}/stream/{identifier}/token`

Operation ID: `stream-videos-create-signed-url-tokens-for-videos`

Creates a signed URL token for a video. If a body is not provided in the request, a token is created with default values.

## Definition

```yaml
{"operationId": "stream-videos-create-signed-url-tokens-for-videos", "summary": "Create signed URL tokens for videos", "description": "Creates a signed URL token for a video. If a body is not provided in the request, a token is created with default values.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_signed_token_request"}}}}, "responses": {"200": {"description": "Create signed URL tokens for videos response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_signed_token_response"}}}}, "4XX": {"description": "Create signed URL tokens for videos response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.token", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
