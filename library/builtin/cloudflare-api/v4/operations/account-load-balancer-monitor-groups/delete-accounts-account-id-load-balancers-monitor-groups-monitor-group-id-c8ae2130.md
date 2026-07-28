---
title: Delete Monitor Group
page_id: operation-delete-accounts-account-id-load-balancers-monitor-groups-monitor-group-i-5de49366
path: operations/account-load-balancer-monitor-groups
description: Delete a configured monitor group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
operation_ids:
    - account-load-balancer-monitor-groups-delete-monitor-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Monitor Group

`DELETE /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}`

Operation ID: `account-load-balancer-monitor-groups-delete-monitor-group`

Delete a configured monitor group.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-delete-monitor-group", "summary": "Delete Monitor Group", "description": "Delete a configured monitor group.", "parameters": [{"name": "monitor_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "Delete Monitor Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}}}}, "412": {"description": "Precondition Failed - Monitor group is in use by one or more pools", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}, "4XX": {"description": "Delete Monitor Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
