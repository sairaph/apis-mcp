---
title: Create watermark profiles via basic upload
page_id: operation-post-accounts-account-id-stream-watermarks-52d7fd3f
path: operations/stream-watermark-profile
description: Creates watermark profiles using a single `HTTP POST multipart/form-data` request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/watermarks
operation_ids:
    - stream-watermark-profile-create-watermark-profiles-via-basic-upload
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create watermark profiles via basic upload

`POST /accounts/{account_id}/stream/watermarks`

Operation ID: `stream-watermark-profile-create-watermark-profiles-via-basic-upload`

Creates watermark profiles using a single `HTTP POST multipart/form-data` request.

## Definition

```yaml
{"operationId": "stream-watermark-profile-create-watermark-profiles-via-basic-upload", "summary": "Create watermark profiles via basic upload", "description": "Creates watermark profiles using a single `HTTP POST multipart/form-data` request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"$ref": "#/components/schemas/stream_name"}, "opacity": {"$ref": "#/components/schemas/stream_opacity"}, "padding": {"$ref": "#/components/schemas/stream_padding"}, "position": {"$ref": "#/components/schemas/stream_position"}, "scale": {"$ref": "#/components/schemas/stream_scale"}, "url": {"description": "URL of the watermark image to copy.", "type": "string", "format": "uri"}}}}, "multipart/form-data": {"schema": {"$ref": "#/components/schemas/stream_watermark_basic_upload"}}}}, "responses": {"200": {"description": "Create watermark profiles via basic upload response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_watermark_response_single"}}}}, "4XX": {"description": "Create watermark profiles via basic upload response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Watermark Profile"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.watermarks", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
