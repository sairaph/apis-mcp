---
title: List Pools
page_id: operation-get-user-load-balancers-pools-a2b65a41
path: operations/load-balancer-pools
description: List configured pools.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/pools
operation_ids:
    - load-balancer-pools-list-pools
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Pools

`GET /user/load_balancers/pools`

Operation ID: `load-balancer-pools-list-pools`

List configured pools.

## Definition

```yaml
{"operationId": "load-balancer-pools-list-pools", "summary": "List Pools", "description": "List configured pools.", "parameters": [{"name": "monitor", "in": "query", "schema": {"description": "The ID of the Monitor to use for checking the health of origins within this pool.", "type": "string"}}], "responses": {"200": {"description": "List Pools response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-response_collection"}}}}, "4XX": {"description": "List Pools response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-response_collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
