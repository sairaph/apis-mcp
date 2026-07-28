---
title: Update Monitor
page_id: operation-put-accounts-account-id-load-balancers-monitors-monitor-id-1de318ae
path: operations/account-load-balancer-monitors
description: Modify a configured monitor.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitors/{monitor_id}
operation_ids:
    - account-load-balancer-monitors-update-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Monitor

`PUT /accounts/{account_id}/load_balancers/monitors/{monitor_id}`

Operation ID: `account-load-balancer-monitors-update-monitor`

Modify a configured monitor.

## Definition

```yaml
{"operationId": "account-load-balancer-monitors-update-monitor", "summary": "Update Monitor", "description": "Modify a configured monitor.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Update Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-single"}}}}, "4XX": {"description": "Update Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-single"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
