---
title: Create Catalog Sync
page_id: operation-post-accounts-account-id-magic-cloud-catalog-syncs-5ff68e5a
path: operations/catalog-sync
description: Create a new Catalog Sync (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/cloud/catalog-syncs
operation_ids:
    - catalog-syncs-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Catalog Sync

`POST /accounts/{account_id}/magic/cloud/catalog-syncs`

Operation ID: `catalog-syncs-create`

Create a new Catalog Sync (Closed Beta).

## Definition

```yaml
{"operationId": "catalog-syncs-create", "summary": "Create Catalog Sync", "description": "Create a new Catalog Sync (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "forwarded", "in": "header", "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_create_catalog_sync_request"}}}}, "responses": {"201": {"description": "Created.", "headers": {"location": {"description": "The path to the newly created list.", "schema": {"type": "string"}}}, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_create_catalog_sync_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "422": {"description": "Unprocessable Entity.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Catalog Sync"], "x-api-token-group": ["Magic WAN Write"]}
```
