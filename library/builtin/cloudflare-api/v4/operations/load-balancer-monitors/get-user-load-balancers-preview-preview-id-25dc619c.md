---
title: Preview Result
page_id: operation-get-user-load-balancers-preview-preview-id-3ff3086f
path: operations/load-balancer-monitors
description: Get the result of a previous preview operation using the provided preview_id.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/load_balancers/preview/{preview_id}
operation_ids:
    - load-balancer-monitors-preview-result
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview Result

`GET /user/load_balancers/preview/{preview_id}`

Operation ID: `load-balancer-monitors-preview-result`

Get the result of a previous preview operation using the provided preview_id.

## Definition

```yaml
{"operationId": "load-balancer-monitors-preview-result", "summary": "Preview Result", "description": "Get the result of a previous preview operation using the provided preview_id.", "parameters": [{"name": "preview_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_preview_id"}}], "responses": {"200": {"description": "Preview Result response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_preview_result_response"}}}}, "4XX": {"description": "Preview Result response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_preview_result_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Load Balancer Monitors"], "x-api-token-group": ["Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read"]}
```
