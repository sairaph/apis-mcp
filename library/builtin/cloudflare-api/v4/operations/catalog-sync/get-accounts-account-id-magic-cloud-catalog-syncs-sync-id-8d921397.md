---
title: Read Catalog Sync
page_id: operation-get-accounts-account-id-magic-cloud-catalog-syncs-sync-id-5305ac86
path: operations/catalog-sync
description: Read a Catalog Sync (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}
operation_ids:
    - catalog-syncs-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read Catalog Sync

`GET /accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}`

Operation ID: `catalog-syncs-read`

Read a Catalog Sync (Closed Beta).

## Definition

```yaml
{"operationId": "catalog-syncs-read", "summary": "Read Catalog Sync", "description": "Read a Catalog Sync (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "sync_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_catalog_sync_id"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_read_account_catalog_sync_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Catalog Sync"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
