---
title: Preview Monitor
page_id: operation-post-user-load-balancers-monitors-monitor-id-preview-ca7bafe7
path: operations/load-balancer-monitors
description: Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /user/load_balancers/monitors/{monitor_id}/preview
operation_ids:
    - load-balancer-monitors-preview-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview Monitor

`POST /user/load_balancers/monitors/{monitor_id}/preview`

Operation ID: `load-balancer-monitors-preview-monitor`

Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.

## Definition

```yaml
{"operationId": "load-balancer-monitors-preview-monitor", "summary": "Preview Monitor", "description": "Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Preview Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_preview_response"}}}}, "4XX": {"description": "Preview Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_preview_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"]}
```
