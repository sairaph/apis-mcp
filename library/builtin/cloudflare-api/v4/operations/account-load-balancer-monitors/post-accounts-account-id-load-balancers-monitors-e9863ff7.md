---
title: Create Monitor
page_id: operation-post-accounts-account-id-load-balancers-monitors-ebde0d7b
path: operations/account-load-balancer-monitors
description: Create a configured monitor.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitors
operation_ids:
    - account-load-balancer-monitors-create-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Monitor

`POST /accounts/{account_id}/load_balancers/monitors`

Operation ID: `account-load-balancer-monitors-create-monitor`

Create a configured monitor.

## Definition

```yaml
{"operationId": "account-load-balancer-monitors-create-monitor", "summary": "Create Monitor", "description": "Create a configured monitor.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Create Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-single"}}}}, "4XX": {"description": "Create Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-single"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
