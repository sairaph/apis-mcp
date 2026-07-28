---
title: List Regions
page_id: operation-get-accounts-account-id-load-balancers-regions-deb3d671
path: operations/load-balancer-regions
description: List all region mappings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/regions
operation_ids:
    - load-balancer-regions-list-regions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Regions

`GET /accounts/{account_id}/load_balancers/regions`

Operation ID: `load-balancer-regions-list-regions`

List all region mappings.

## Definition

```yaml
{"operationId": "load-balancer-regions-list-regions", "summary": "List Regions", "description": "List all region mappings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}, {"name": "subdivision_code", "in": "query", "schema": {"$ref": "#/components/schemas/load-balancing_subdivision_code_a2"}}, {"name": "subdivision_code_a2", "in": "query", "schema": {"$ref": "#/components/schemas/load-balancing_subdivision_code_a2"}}, {"name": "country_code_a2", "in": "query", "schema": {"description": "Two-letter alpha-2 country code followed in ISO 3166-1.", "type": "string", "example": "US"}}], "responses": {"200": {"description": "List Regions response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_region_components-schemas-response_collection"}}}}, "4XX": {"description": "List Regions response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_region_components-schemas-response_collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Regions"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
