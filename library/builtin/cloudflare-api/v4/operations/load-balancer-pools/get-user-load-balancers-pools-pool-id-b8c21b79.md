---
title: Pool Details
page_id: operation-get-user-load-balancers-pools-pool-id-04bae44a
path: operations/load-balancer-pools
description: Fetch a single configured pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/pools/{pool_id}
operation_ids:
    - load-balancer-pools-pool-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Pool Details

`GET /user/load_balancers/pools/{pool_id}`

Operation ID: `load-balancer-pools-pool-details`

Fetch a single configured pool.

## Definition

```yaml
{"operationId": "load-balancer-pools-pool-details", "summary": "Pool Details", "description": "Fetch a single configured pool.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}], "responses": {"200": {"description": "Pool Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-single_response"}}}}, "4XX": {"description": "Pool Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
