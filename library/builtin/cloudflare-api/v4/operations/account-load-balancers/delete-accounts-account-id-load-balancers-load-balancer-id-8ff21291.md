---
title: Delete Account Load Balancer
page_id: operation-delete-accounts-account-id-load-balancers-load-balancer-id-2121ae78
path: operations/account-load-balancers
description: Delete a configured account-scoped load balancer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/load_balancers/{load_balancer_id}
operation_ids:
    - account-load-balancers-delete-account-load-balancer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Account Load Balancer

`DELETE /accounts/{account_id}/load_balancers/{load_balancer_id}`

Operation ID: `account-load-balancers-delete-account-load-balancer`

Delete a configured account-scoped load balancer.

## Definition

```yaml
{"operationId": "account-load-balancers-delete-account-load-balancer", "summary": "Delete Account Load Balancer", "description": "Delete a configured account-scoped load balancer.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}, {"name": "load_balancer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-identifier"}}], "responses": {"200": {"description": "Delete Account Load Balancer response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-id_response"}}}}, "4XX": {"description": "Delete Account Load Balancer response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_components-schemas-id_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancers Account Write"]}
```
