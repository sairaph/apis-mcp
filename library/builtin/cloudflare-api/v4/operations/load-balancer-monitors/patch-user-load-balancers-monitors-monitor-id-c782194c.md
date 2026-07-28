---
title: Patch Monitor
page_id: operation-patch-user-load-balancers-monitors-monitor-id-75edb259
path: operations/load-balancer-monitors
description: Apply changes to an existing monitor, overwriting the supplied properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /user/load_balancers/monitors/{monitor_id}
operation_ids:
    - load-balancer-monitors-patch-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Monitor

`PATCH /user/load_balancers/monitors/{monitor_id}`

Operation ID: `load-balancer-monitors-patch-monitor`

Apply changes to an existing monitor, overwriting the supplied properties.

## Definition

```yaml
{"operationId": "load-balancer-monitors-patch-monitor", "summary": "Patch Monitor", "description": "Apply changes to an existing monitor, overwriting the supplied properties.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Patch Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-single"}}}}, "4XX": {"description": "Patch Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-single"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
