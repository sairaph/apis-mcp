---
title: List Regions
page_id: operation-get-accounts-account-id-addressing-regional-hostnames-regions-97f45711
path: operations/dls-regional-services
description: List all Regional Services regions available for use by this account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/regional_hostnames/regions
operation_ids:
    - dls-account-regional-hostnames-list-regions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Regions

`GET /accounts/{account_id}/addressing/regional_hostnames/regions`

Operation ID: `dls-account-regional-hostnames-list-regions`

List all Regional Services regions available for use by this account.

## Definition

```yaml
{"operationId": "dls-account-regional-hostnames-list-regions", "summary": "List Regions", "description": "List all Regional Services regions available for use by this account.", "parameters": [{"$ref": "#/components/parameters/dls_account_id"}], "responses": {"200": {"description": "List regions response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"key": {"$ref": "#/components/schemas/dls_region_key"}, "label": {"description": "Human-readable text label for the region", "type": "string", "example": "Canada"}}, "type": "object"}}}}]}}}}, "4XX": {"description": "Failure to list regions", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dls_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}], "tags": ["DLS Regional Services"], "x-api-token-group": ["DNS Read", "DNS Write"]}
```
