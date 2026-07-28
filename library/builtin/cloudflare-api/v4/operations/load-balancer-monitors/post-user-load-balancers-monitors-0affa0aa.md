---
title: Create Monitor
page_id: operation-post-user-load-balancers-monitors-e632948e
path: operations/load-balancer-monitors
description: Create a configured monitor.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /user/load_balancers/monitors
operation_ids:
    - load-balancer-monitors-create-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Monitor

`POST /user/load_balancers/monitors`

Operation ID: `load-balancer-monitors-create-monitor`

Create a configured monitor.

## Definition

```yaml
{"operationId": "load-balancer-monitors-create-monitor", "summary": "Create Monitor", "description": "Create a configured monitor.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Create Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-single"}}}}, "4XX": {"description": "Create Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-single"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write"]}
```
