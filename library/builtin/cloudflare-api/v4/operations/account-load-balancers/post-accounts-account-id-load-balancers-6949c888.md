---
title: Create Account Load Balancer
page_id: operation-post-accounts-account-id-load-balancers-0a596769
path: operations/account-load-balancers
description: Create a new account-scoped load balancer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/load_balancers
operation_ids:
    - account-load-balancers-create-account-load-balancer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Account Load Balancer

`POST /accounts/{account_id}/load_balancers`

Operation ID: `account-load-balancers-create-account-load-balancer`

Create a new account-scoped load balancer.

## Definition

```yaml
{"operationId": "account-load-balancers-create-account-load-balancer", "summary": "Create Account Load Balancer", "description": "Create a new account-scoped load balancer.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer-editable"}, {"required": ["name", "default_pools", "fallback_pool"], "type": "object"}]}}}}, "responses": {"200": {"description": "Create Account Load Balancer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}}}}, "4XX": {"description": "Create Account Load Balancer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancers Account Write"]}
```
