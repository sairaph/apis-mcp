---
title: List watermark profiles
page_id: operation-get-accounts-account-id-stream-watermarks-1efc6843
path: operations/stream-watermark-profile
description: Lists all watermark profiles for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/watermarks
operation_ids:
    - stream-watermark-profile-list-watermark-profiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List watermark profiles

`GET /accounts/{account_id}/stream/watermarks`

Operation ID: `stream-watermark-profile-list-watermark-profiles`

Lists all watermark profiles for an account.

## Definition

```yaml
{"operationId": "stream-watermark-profile-list-watermark-profiles", "summary": "List watermark profiles", "description": "Lists all watermark profiles for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "List watermark profiles response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_watermark_response_collection"}}}}, "4XX": {"description": "List watermark profiles response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Watermark Profile"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.watermarks", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
