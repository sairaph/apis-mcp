---
title: Retrieve video details
page_id: operation-get-accounts-account-id-stream-identifier-0abc1cce
path: operations/stream-videos
description: Fetches details for a single video.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}
operation_ids:
    - stream-videos-retrieve-video-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve video details

`GET /accounts/{account_id}/stream/{identifier}`

Operation ID: `stream-videos-retrieve-video-details`

Fetches details for a single video.

## Definition

```yaml
{"operationId": "stream-videos-retrieve-video-details", "summary": "Retrieve video details", "description": "Fetches details for a single video.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "Retrieve video details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_video_response_single"}}}}, "4XX": {"description": "Retrieve video details response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
