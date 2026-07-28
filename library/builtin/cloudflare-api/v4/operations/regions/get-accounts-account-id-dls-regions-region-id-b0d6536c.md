---
title: Get a DLS region
page_id: operation-get-accounts-account-id-dls-regions-region-id-b555ca54
path: operations/regions
description: Retrieve a single DLS region (managed or custom) by ID or region key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dls/regions/{region_id}
operation_ids:
    - publicGetRegion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a DLS region

`GET /accounts/{account_id}/dls/regions/{region_id}`

Operation ID: `publicGetRegion`

Retrieve a single DLS region (managed or custom) by ID or region key.

## Definition

```yaml
{"operationId": "publicGetRegion", "summary": "Get a DLS region", "description": "Retrieve a single DLS region (managed or custom) by ID or region key.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}, {"name": "region_id", "in": "path", "required": true, "schema": {"description": "UUID of the region (custom or managed) or region_key of a managed region.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890", "maxLength": 255}}], "responses": {"200": {"description": "Region found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_FetchPublicRegionResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Regions"], "x-api-token-group": ["DLS: Read", "DLS: Write"]}
```
