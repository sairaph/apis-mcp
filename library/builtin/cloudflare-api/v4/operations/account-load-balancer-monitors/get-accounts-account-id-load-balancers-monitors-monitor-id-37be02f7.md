---
title: Monitor Details
page_id: operation-get-accounts-account-id-load-balancers-monitors-monitor-id-d4e4f13f
path: operations/account-load-balancer-monitors
description: List a single configured monitor for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitors/{monitor_id}
operation_ids:
    - account-load-balancer-monitors-monitor-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Monitor Details

`GET /accounts/{account_id}/load_balancers/monitors/{monitor_id}`

Operation ID: `account-load-balancer-monitors-monitor-details`

List a single configured monitor for an account.

## Definition

```yaml
{"operationId": "account-load-balancer-monitors-monitor-details", "summary": "Monitor Details", "description": "List a single configured monitor for an account.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "Monitor Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-single"}}}}, "4XX": {"description": "Monitor Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-single"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
