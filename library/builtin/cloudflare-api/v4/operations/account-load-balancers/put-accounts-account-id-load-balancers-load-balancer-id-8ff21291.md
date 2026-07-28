---
title: Update Account Load Balancer
page_id: operation-put-accounts-account-id-load-balancers-load-balancer-id-302fff6d
path: operations/account-load-balancers
description: Update a configured account-scoped load balancer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/load_balancers/{load_balancer_id}
operation_ids:
    - account-load-balancers-update-account-load-balancer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Account Load Balancer

`PUT /accounts/{account_id}/load_balancers/{load_balancer_id}`

Operation ID: `account-load-balancers-update-account-load-balancer`

Update a configured account-scoped load balancer.

## Definition

```yaml
{"operationId": "account-load-balancers-update-account-load-balancer", "summary": "Update Account Load Balancer", "description": "Update a configured account-scoped load balancer.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer-editable"}, {"required": ["name", "default_pools", "fallback_pool"], "type": "object"}]}}}}, "responses": {"200": {"description": "Update Account Load Balancer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}}}}, "4XX": {"description": "Update Account Load Balancer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancers Account Write"]}
```
