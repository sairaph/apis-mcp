---
title: List Prebuilt Policies
page_id: operation-get-accounts-account-id-magic-cloud-catalog-syncs-prebuilt-policies-fe20f07a
path: operations/catalog-sync
description: List prebuilt catalog sync policies (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/catalog-syncs/prebuilt-policies
operation_ids:
    - catalog-syncs-prebuilt-policies-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Prebuilt Policies

`GET /accounts/{account_id}/magic/cloud/catalog-syncs/prebuilt-policies`

Operation ID: `catalog-syncs-prebuilt-policies-list`

List prebuilt catalog sync policies (Closed Beta).

## Definition

```yaml
{"operationId": "catalog-syncs-prebuilt-policies-list", "summary": "List Prebuilt Policies", "description": "List prebuilt catalog sync policies (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "destination_type", "in": "query", "description": "Specify type of destination, omit to return all.", "schema": {"$ref": "#/components/schemas/mcn_catalog_sync_destination_type"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_catalog_syncs_prebuilt_policies_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Catalog Sync"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
