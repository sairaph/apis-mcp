---
title: Monitor Group Details
page_id: operation-get-accounts-account-id-load-balancers-monitor-groups-monitor-group-id-b8c11667
path: operations/account-load-balancer-monitor-groups
description: Fetch a single configured monitor group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}
operation_ids:
    - account-load-balancer-monitor-groups-monitor-group-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Monitor Group Details

`GET /accounts/{account_id}/load_balancers/monitor_groups/{monitor_group_id}`

Operation ID: `account-load-balancer-monitor-groups-monitor-group-details`

Fetch a single configured monitor group.

## Definition

```yaml
{"operationId": "account-load-balancer-monitor-groups-monitor-group-details", "summary": "Monitor Group Details", "description": "Fetch a single configured monitor group.", "parameters": [{"name": "monitor_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "responses": {"200": {"description": "Monitor Group Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}}}}, "4XX": {"description": "Monitor Group Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-group-single-response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitor Groups"]}
```
