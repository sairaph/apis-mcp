---
title: List Monitor Group References
page_id: operation-get-accounts-account-id-load-balancers-monitor-groups-monitor-group-id-r-5db628d4
path: operations/account-load-balancer-monitor-groups
description: Get the list of resources that reference the provided monitor group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}/references
operation_ids:
    - account-load-balancer-monitor-groups-list-monitor-group-references
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Monitor Group References

`GET /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}/references`

Operation ID: `account-load-balancer-monitor-groups-list-monitor-group-references`

Get the list of resources that reference the provided monitor group.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-list-monitor-group-references", "summary": "List Monitor Group References", "description": "Get the list of resources that reference the provided monitor group.", "parameters": [{"name": "monitor_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "List Monitor Group References response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-references-response"}}}}, "4XX": {"description": "List Monitor Group References response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-references-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
