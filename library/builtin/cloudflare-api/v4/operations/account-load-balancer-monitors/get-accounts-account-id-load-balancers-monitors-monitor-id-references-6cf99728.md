---
title: List Monitor References
page_id: operation-get-accounts-account-id-load-balancers-monitors-monitor-id-references-8cce7075
path: operations/account-load-balancer-monitors
description: Get the list of resources that reference the provided monitor.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitors/{monitor_id}/references
operation_ids:
    - account-load-balancer-monitors-list-monitor-references
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Monitor References

`GET /accounts/{account_id}/load_balancers/monitors/{monitor_id}/references`

Operation ID: `account-load-balancer-monitors-list-monitor-references`

Get the list of resources that reference the provided monitor.

## Definition

```yaml
{"operationId": "account-load-balancer-monitors-list-monitor-references", "summary": "List Monitor References", "description": "Get the list of resources that reference the provided monitor.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "List Monitor References response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-references-response"}}}}, "4XX": {"description": "List Monitor References response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-references-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
