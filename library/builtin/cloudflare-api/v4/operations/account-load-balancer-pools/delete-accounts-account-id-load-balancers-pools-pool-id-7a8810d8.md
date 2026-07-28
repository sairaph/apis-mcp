---
title: Delete Pool
page_id: operation-delete-accounts-account-id-load-balancers-pools-pool-id-40d80fa1
path: operations/account-load-balancer-pools
description: Delete a configured pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/load_balancers/pools/{pool_id}
operation_ids:
    - account-load-balancer-pools-delete-pool
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Pool

`DELETE /accounts/{account_id}/load_balancers/pools/{pool_id}`

Operation ID: `account-load-balancer-pools-delete-pool`

Delete a configured pool.

## Definition

```yaml
{"operationId": "account-load-balancer-pools-delete-pool", "summary": "Delete Pool", "description": "Delete a configured pool.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Pool response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_schemas-id_response"}}}}, "4XX": {"description": "Delete Pool response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_schemas-id_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Pools"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
