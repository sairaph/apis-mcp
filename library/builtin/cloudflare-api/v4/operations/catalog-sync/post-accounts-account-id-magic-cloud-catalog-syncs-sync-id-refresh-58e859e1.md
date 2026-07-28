---
title: Run Catalog Sync
page_id: operation-post-accounts-account-id-magic-cloud-catalog-syncs-sync-id-refresh-6b7d5911
path: operations/catalog-sync
description: Refresh a Catalog Sync's destination by running the sync policy against latest resource catalog (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}/refresh
operation_ids:
    - catalog-syncs-refresh
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Run Catalog Sync

`POST /accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}/refresh`

Operation ID: `catalog-syncs-refresh`

Refresh a Catalog Sync's destination by running the sync policy against latest resource catalog (Closed Beta).

## Definition

```yaml
{"operationId": "catalog-syncs-refresh", "summary": "Run Catalog Sync", "description": "Refresh a Catalog Sync's destination by running the sync policy against latest resource catalog (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "sync_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_catalog_sync_id"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_refresh_catalog_sync_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "422": {"description": "Unprocessable Entity.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Catalog Sync"], "x-api-token-group": ["Magic WAN Write"]}
```
