---
title: Patch Account Load Balancer
page_id: operation-patch-accounts-account-id-load-balancers-load-balancer-id-25a02193
path: operations/account-load-balancers
description: Apply changes to an existing account-scoped load balancer, overwriting the supplied properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/load_balancers/{load_balancer_id}
operation_ids:
    - account-load-balancers-patch-account-load-balancer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Account Load Balancer

`PATCH /accounts/{account_id}/load_balancers/{load_balancer_id}`

Operation ID: `account-load-balancers-patch-account-load-balancer`

Apply changes to an existing account-scoped load balancer, overwriting the supplied properties.

## Definition

```yaml
{"operationId": "account-load-balancers-patch-account-load-balancer", "summary": "Patch Account Load Balancer", "description": "Apply changes to an existing account-scoped load balancer, overwriting the supplied properties.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer-editable"}]}}}}, "responses": {"200": {"description": "Patch Account Load Balancer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}}}}, "4XX": {"description": "Patch Account Load Balancer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-single_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancers Account Write"]}
```
