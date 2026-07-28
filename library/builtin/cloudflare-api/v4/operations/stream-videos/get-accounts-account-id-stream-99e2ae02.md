---
title: List videos
page_id: operation-get-accounts-account-id-stream-d9761303
path: operations/stream-videos
description: Lists up to 1000 videos from a single request. For a specific range, refer to the optional parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream
operation_ids:
    - stream-videos-list-videos
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List videos

`GET /accounts/{account_id}/stream`

Operation ID: `stream-videos-list-videos`

Lists up to 1000 videos from a single request. For a specific range, refer to the optional parameters.

## Definition

```yaml
{"operationId": "stream-videos-list-videos", "summary": "List videos", "description": "Lists up to 1000 videos from a single request. For a specific range, refer to the optional parameters.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "status", "in": "query", "schema": {"$ref": "#/components/schemas/stream_media_state"}}, {"name": "creator", "in": "query", "schema": {"$ref": "#/components/schemas/stream_creator"}}, {"name": "type", "in": "query", "schema": {"$ref": "#/components/schemas/stream_type"}}, {"name": "asc", "in": "query", "schema": {"$ref": "#/components/schemas/stream_asc"}}, {"name": "video_name", "in": "query", "schema": {"$ref": "#/components/schemas/stream_video_name"}}, {"name": "search", "in": "query", "schema": {"$ref": "#/components/schemas/stream_search"}}, {"name": "start", "in": "query", "schema": {"$ref": "#/components/schemas/stream_start"}}, {"name": "end", "in": "query", "schema": {"$ref": "#/components/schemas/stream_end"}}, {"name": "include_counts", "in": "query", "schema": {"$ref": "#/components/schemas/stream_include_counts"}}, {"name": "id", "in": "query", "description": "Filter by video ID(s). Can be a single ID or a comma-separated list of IDs.", "schema": {"type": "string", "example": "ea95132c15732412d22c1476fa83f27a"}}, {"name": "name", "in": "query", "description": "Filter by video name/UID(s). Can be a single name or a comma-separated list.", "schema": {"type": "string"}}, {"name": "live_input_id", "in": "query", "description": "Filter by live input ID to find videos associated with a specific live stream.", "schema": {"type": "string"}}, {"name": "before", "in": "query", "description": "Alias for 'end'. Returns videos created before this date/time (RFC 3339 format).", "schema": {"type": "string", "format": "date-time"}}, {"name": "after", "in": "query", "description": "Alias for 'start'. Returns videos created after this date/time (RFC 3339 format).", "schema": {"type": "string", "format": "date-time"}}, {"name": "limit", "in": "query", "description": "Maximum number of videos to return (default 1000, max 1000).", "schema": {"type": "integer", "maximum": 1000, "minimum": 1}}], "responses": {"200": {"description": "List videos response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_video_response_collection"}}}}, "4XX": {"description": "List videos response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
