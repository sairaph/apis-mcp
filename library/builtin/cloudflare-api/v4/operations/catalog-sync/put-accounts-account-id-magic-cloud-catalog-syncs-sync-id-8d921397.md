---
title: Update Catalog Sync
page_id: operation-put-accounts-account-id-magic-cloud-catalog-syncs-sync-id-1a452b44
path: operations/catalog-sync
description: Update a Catalog Sync (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}
operation_ids:
    - catalog-syncs-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Catalog Sync

`PUT /accounts/{account_id}/magic/cloud/catalog-syncs/{sync_id}`

Operation ID: `catalog-syncs-update`

Update a Catalog Sync (Closed Beta).

## Definition

```yaml
{"operationId": "catalog-syncs-update", "summary": "Update Catalog Sync", "description": "Update a Catalog Sync (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "sync_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_catalog_sync_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_update_catalog_sync_request"}}}}, "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_update_catalog_sync_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "422": {"description": "Unprocessable Entity.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Catalog Sync"], "x-api-token-group": ["Magic WAN Write"]}
```
