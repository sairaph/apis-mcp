---
title: Storage use
page_id: operation-get-accounts-account-id-stream-storage-usage-1c79aa9f
path: operations/stream-videos
description: Returns information about an account's storage use.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/storage-usage
operation_ids:
    - stream-videos-storage-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Storage use

`GET /accounts/{account_id}/stream/storage-usage`

Operation ID: `stream-videos-storage-usage`

Returns information about an account's storage use.

## Definition

```yaml
{"operationId": "stream-videos-storage-usage", "summary": "Storage use", "description": "Returns information about an account's storage use.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "creator", "in": "query", "schema": {"$ref": "#/components/schemas/stream_creator"}}], "responses": {"200": {"description": "Returns information about an account's storage use response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_storage_use_response"}}}}, "4XX": {"description": "Returns information about an account's storage use response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.videos", "x-fern-sdk-method-name": "storage-usage", "x-forge-hidden": true}
```
