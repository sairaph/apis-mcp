---
title: List Pool References
page_id: operation-get-user-load-balancers-pools-pool-id-references-75726c70
path: operations/load-balancer-pools
description: Get the list of resources that reference the provided pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/pools/{pool_id}/references
operation_ids:
    - load-balancer-pools-list-pool-references
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Pool References

`GET /user/load_balancers/pools/{pool_id}/references`

Operation ID: `load-balancer-pools-list-pool-references`

Get the list of resources that reference the provided pool.

## Definition

```yaml
{"operationId": "load-balancer-pools-list-pool-references", "summary": "List Pool References", "description": "Get the list of resources that reference the provided pool.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}], "responses": {"200": {"description": "List Pool References response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_pools-references-response"}}}}, "4XX": {"description": "List Pool References response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_pools-references-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
