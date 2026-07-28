---
title: Pool Details
page_id: operation-get-accounts-account-id-load-balancers-pools-pool-id-9a22495d
path: operations/account-load-balancer-pools
description: Fetch a single configured pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/pools/{pool_id}
operation_ids:
    - account-load-balancer-pools-pool-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Pool Details

`GET /accounts/{account_id}/load_balancers/pools/{pool_id}`

Operation ID: `account-load-balancer-pools-pool-details`

Fetch a single configured pool.

## Definition

```yaml
{"operationId": "account-load-balancer-pools-pool-details", "summary": "Pool Details", "description": "Fetch a single configured pool.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "Pool Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-single_response"}}}}, "4XX": {"description": "Pool Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
