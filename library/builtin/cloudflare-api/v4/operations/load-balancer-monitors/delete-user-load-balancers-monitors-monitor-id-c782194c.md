---
title: Delete Monitor
page_id: operation-delete-user-load-balancers-monitors-monitor-id-274d65d0
path: operations/load-balancer-monitors
description: Delete a configured monitor.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /user/load_balancers/monitors/{monitor_id}
operation_ids:
    - load-balancer-monitors-delete-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Monitor

`DELETE /user/load_balancers/monitors/{monitor_id}`

Operation ID: `load-balancer-monitors-delete-monitor`

Delete a configured monitor.

## Definition

```yaml
{"operationId": "load-balancer-monitors-delete-monitor", "summary": "Delete Monitor", "description": "Delete a configured monitor.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_id_response"}}}}, "4XX": {"description": "Delete Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_id_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
