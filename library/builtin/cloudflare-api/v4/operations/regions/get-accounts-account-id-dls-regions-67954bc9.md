---
title: List DLS regions for an account
page_id: operation-get-accounts-account-id-dls-regions-12090080
path: operations/regions
description: List the DLS regions (managed and custom) available to an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dls/regions
operation_ids:
    - publicListRegions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DLS regions for an account

`GET /accounts/{account_id}/dls/regions`

Operation ID: `publicListRegions`

List the DLS regions (managed and custom) available to an account.

## Definition

```yaml
{"operationId": "publicListRegions", "summary": "List DLS regions for an account", "description": "List the DLS regions (managed and custom) available to an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}, {"name": "cursor", "in": "query", "description": "Opaque token for cursor-based pagination. Omit for the first page. Pass the value from a previous response to fetch the next page.", "schema": {"type": "string"}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 25, "maximum": 100, "minimum": 1}}, {"name": "type", "in": "query", "description": "Filter regions by type. Omit to return all regions.", "schema": {"type": "string", "enum": ["managed", "custom"]}}], "responses": {"200": {"description": "List of regions.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_RegionPublicPaginatedListResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Regions"], "x-api-token-group": ["DLS: Read", "DLS: Write"]}
```
