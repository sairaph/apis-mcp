---
title: List Load Balancer Usage
page_id: operation-get-accounts-account-id-load-balancers-usage-ab595d4a
path: operations/account-load-balancers
description: Get current load balancer resource usage counts for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/usage
operation_ids:
    - account-load-balancers-list-load-balancer-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Load Balancer Usage

`GET /accounts/{account_id}/load_balancers/usage`

Operation ID: `account-load-balancers-list-load-balancer-usage`

Get current load balancer resource usage counts for an account.

## Definition

```yaml
{"operationId": "account-load-balancers-list-load-balancer-usage", "summary": "List Load Balancer Usage", "description": "Get current load balancer resource usage counts for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "Load Balancer Usage response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-usage_response"}}}}, "4XX": {"description": "Load Balancer Usage response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-usage_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
