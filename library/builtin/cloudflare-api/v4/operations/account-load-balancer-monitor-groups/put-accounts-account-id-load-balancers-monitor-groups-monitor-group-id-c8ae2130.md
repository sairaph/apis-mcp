---
title: Update Monitor Group
page_id: operation-put-accounts-account-id-load-balancers-monitor-groups-monitor-group-id-587a9444
path: operations/account-load-balancer-monitor-groups
description: Modify a configured monitor group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
operation_ids:
    - account-load-balancer-monitor-groups-update-monitor-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Monitor Group

`PUT /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}`

Operation ID: `account-load-balancer-monitor-groups-update-monitor-group`

Modify a configured monitor group.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-update-monitor-group", "summary": "Update Monitor Group", "description": "Modify a configured monitor group.", "parameters": [{"name": "monitor_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group"}}}}, "responses": {"200": {"description": "Update Monitor Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}}}}, "412": {"description": "Precondition Failed - Referenced monitor does not exist", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}, "4XX": {"description": "Update Monitor Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
