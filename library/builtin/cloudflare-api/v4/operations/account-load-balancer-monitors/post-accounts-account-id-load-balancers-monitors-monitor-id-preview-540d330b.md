---
title: Preview Monitor
page_id: operation-post-accounts-account-id-load-balancers-monitors-monitor-id-preview-1762839b
path: operations/account-load-balancer-monitors
description: Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/load_balancers/monitors/{monitor_id}/preview
operation_ids:
    - account-load-balancer-monitors-preview-monitor
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview Monitor

`POST /accounts/{account_id}/load_balancers/monitors/{monitor_id}/preview`

Operation ID: `account-load-balancer-monitors-preview-monitor`

Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.

## Definition

```yaml
{"operationId": "account-load-balancer-monitors-preview-monitor", "summary": "Preview Monitor", "description": "Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.", "parameters": [{"name": "monitor_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Preview Monitor response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_preview_response"}}}}, "4XX": {"description": "Preview Monitor response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_preview_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Monitors"]}
```
