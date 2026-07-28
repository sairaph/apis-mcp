---
title: List Monitors
page_id: operation-get-accounts-account-id-load-balancers-monitors-1d54e28a
path: operations/account-load-balancer-monitors
description: List configured monitors for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitors
operation_ids:
    - account-load-balancer-monitors-list-monitors
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Monitors

`GET /accounts/{account_id}/load_balancers/monitors`

Operation ID: `account-load-balancer-monitors-list-monitors`

List configured monitors for an account.

## Definition

```yaml
{"operationId": "account-load-balancer-monitors-list-monitors", "summary": "List Monitors", "description": "List configured monitors for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "List Monitors response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-collection"}}}}, "4XX": {"description": "List Monitors response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
