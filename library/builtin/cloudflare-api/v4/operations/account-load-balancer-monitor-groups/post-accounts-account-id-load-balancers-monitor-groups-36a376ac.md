---
title: Create Monitor Group
page_id: operation-post-accounts-account-id-load-balancers-monitor-groups-28d6060e
path: operations/account-load-balancer-monitor-groups
description: Create a new monitor group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups
operation_ids:
    - account-load-balancer-monitor-groups-create-monitor-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Monitor Group

`POST /accounts/{account_id}/load_balancers/monitor_groups`

Operation ID: `account-load-balancer-monitor-groups-create-monitor-group`

Create a new monitor group.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-create-monitor-group", "summary": "Create Monitor Group", "description": "Create a new monitor group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group"}}}}, "responses": {"200": {"description": "Create Monitor Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}}}}, "412": {"description": "Precondition Failed - Referenced monitor does not exist", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}, "4XX": {"description": "Create Monitor Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
