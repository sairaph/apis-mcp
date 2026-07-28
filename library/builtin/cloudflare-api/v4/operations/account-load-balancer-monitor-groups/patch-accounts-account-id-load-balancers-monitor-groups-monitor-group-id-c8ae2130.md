---
title: Patch Monitor Group
page_id: operation-patch-accounts-account-id-load-balancers-monitor-groups-monitor-group-id-2e69c6d5
path: operations/account-load-balancer-monitor-groups
description: Apply changes to an existing monitor group, overwriting the supplied properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
operation_ids:
    - account-load-balancer-monitor-groups-patch-monitor-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Monitor Group

`PATCH /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}`

Operation ID: `account-load-balancer-monitor-groups-patch-monitor-group`

Apply changes to an existing monitor group, overwriting the supplied properties.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-patch-monitor-group", "summary": "Patch Monitor Group", "description": "Apply changes to an existing monitor group, overwriting the supplied properties.", "parameters": [{"name": "monitor_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group"}}}}, "responses": {"200": {"description": "Patch Monitor Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}}}}, "412": {"description": "Precondition Failed - Referenced monitor does not exist", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}, "4XX": {"description": "Patch Monitor Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
