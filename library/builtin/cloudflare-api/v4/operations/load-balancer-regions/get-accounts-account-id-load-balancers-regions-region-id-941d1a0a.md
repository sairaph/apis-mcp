---
title: Get Region
page_id: operation-get-accounts-account-id-load-balancers-regions-region-id-178b0d2e
path: operations/load-balancer-regions
description: Get a single region mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/regions/{region_id}
operation_ids:
    - load-balancer-regions-get-region
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Region

`GET /accounts/{account_id}/load_balancers/regions/{region_id}`

Operation ID: `load-balancer-regions-get-region`

Get a single region mapping.

## Definition

```yaml
{"operationId": "load-balancer-regions-get-region", "summary": "Get Region", "description": "Get a single region mapping.", "parameters": [{"name": "region_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_region_code"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "Get Region response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-single_response"}}}}, "4XX": {"description": "Get Region response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Regions"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
