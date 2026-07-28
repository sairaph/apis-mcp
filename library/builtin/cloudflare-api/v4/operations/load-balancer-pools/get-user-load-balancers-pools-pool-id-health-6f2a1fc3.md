---
title: Pool Health Details
page_id: operation-get-user-load-balancers-pools-pool-id-health-312cb5cc
path: operations/load-balancer-pools
description: Fetch the latest pool health status for a single pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/pools/{pool_id}/health
operation_ids:
    - load-balancer-pools-pool-health-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Pool Health Details

`GET /user/load_balancers/pools/{pool_id}/health`

Operation ID: `load-balancer-pools-pool-health-details`

Fetch the latest pool health status for a single pool.

## Definition

```yaml
{"operationId": "load-balancer-pools-pool-health-details", "summary": "Pool Health Details", "description": "Fetch the latest pool health status for a single pool.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}], "responses": {"200": {"description": "Pool Health Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_health_details"}}}}, "4XX": {"description": "Pool Health Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_health_details"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
