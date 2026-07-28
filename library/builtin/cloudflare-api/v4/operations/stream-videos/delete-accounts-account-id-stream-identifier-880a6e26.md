---
title: Delete video
page_id: operation-delete-accounts-account-id-stream-identifier-8f1132d2
path: operations/stream-videos
description: Deletes a video and its copies from Cloudflare Stream.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}
operation_ids:
    - stream-videos-delete-video
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete video

`DELETE /accounts/{account_id}/stream/{identifier}`

Operation ID: `stream-videos-delete-video`

Deletes a video and its copies from Cloudflare Stream.

## Definition

```yaml
{"operationId": "stream-videos-delete-video", "summary": "Delete video", "description": "Deletes a video and its copies from Cloudflare Stream.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "Delete video response."}, "4XX": {"description": "Delete video response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
