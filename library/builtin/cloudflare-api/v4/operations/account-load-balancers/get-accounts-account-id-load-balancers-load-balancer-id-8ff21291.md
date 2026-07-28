---
title: Account Load Balancer Details
page_id: operation-get-accounts-account-id-load-balancers-load-balancer-id-a1ad54f9
path: operations/account-load-balancers
description: Fetch a single configured account-scoped load balancer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/{load_balancer_id}
operation_ids:
    - account-load-balancers-account-load-balancer-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Account Load Balancer Details

`GET /accounts/{account_id}/load_balancers/{load_balancer_id}`

Operation ID: `account-load-balancers-account-load-balancer-details`

Fetch a single configured account-scoped load balancer.

## Definition

```yaml
{"operationId": "account-load-balancers-account-load-balancer-details", "summary": "Account Load Balancer Details", "description": "Fetch a single configured account-scoped load balancer.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "responses": {"200": {"description": "Account Load Balancer Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}}}}, "4XX": {"description": "Account Load Balancer Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancers Account Write", "Load Balancers Account Read"]}
```
