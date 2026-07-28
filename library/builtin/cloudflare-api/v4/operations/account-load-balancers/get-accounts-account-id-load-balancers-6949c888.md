---
title: List Account Load Balancers
page_id: operation-get-accounts-account-id-load-balancers-f5f0c4eb
path: operations/account-load-balancers
description: List configured account-scoped load balancers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers
operation_ids:
    - account-load-balancers-list-account-load-balancers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Account Load Balancers

`GET /accounts/{account_id}/load_balancers`

Operation ID: `account-load-balancers-list-account-load-balancers`

List configured account-scoped load balancers.

## Definition

```yaml
{"operationId": "account-load-balancers-list-account-load-balancers", "summary": "List Account Load Balancers", "description": "List configured account-scoped load balancers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "List Account Load Balancers response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-response_collection"}}}}, "4XX": {"description": "List Account Load Balancers response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_load-balancer_components-schemas-response_collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancers"], "x-api-token-group": ["Load Balancers Account Write", "Load Balancers Account Read"]}
```
