---
title: List Monitor Groups
page_id: operation-get-accounts-account-id-load-balancers-monitor-groups-cc673ac4
path: operations/account-load-balancer-monitor-groups
description: List configured monitor groups.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups
operation_ids:
    - account-load-balancer-monitor-groups-list-monitor-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Monitor Groups

`GET /accounts/{account_id}/load_balancers/monitor_groups`

Operation ID: `account-load-balancer-monitor-groups-list-monitor-groups`

List configured monitor groups.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-list-monitor-groups", "summary": "List Monitor Groups", "description": "List configured monitor groups.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "List Monitor Groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-response-collection"}}}}, "4XX": {"description": "List Monitor Groups response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-response-collection"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
