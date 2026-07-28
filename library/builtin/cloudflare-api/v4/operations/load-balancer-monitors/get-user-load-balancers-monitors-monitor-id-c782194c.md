---
title: Monitor Details
page_id: operation-get-user-load-balancers-monitors-monitor-id-76fffe77
path: operations/load-balancer-monitors
description: List a single configured monitor for a user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/monitors/{monitor_id}
operation_ids:
    - load-balancer-monitors-monitor-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Monitor Details

`GET /user/load_balancers/monitors/{monitor_id}`

Operation ID: `load-balancer-monitors-monitor-details`

List a single configured monitor for a user.

## Definition

```yaml
{"operationId": "load-balancer-monitors-monitor-details", "summary": "Monitor Details", "description": "List a single configured monitor for a user.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}], "responses": {"200": {"description": "Monitor Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_monitor-response-single"}}}}, "4XX": {"description": "Monitor Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-response-single"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
