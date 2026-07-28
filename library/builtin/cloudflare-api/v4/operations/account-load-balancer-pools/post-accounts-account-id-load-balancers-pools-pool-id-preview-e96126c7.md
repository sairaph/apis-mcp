---
title: Preview Pool
page_id: operation-post-accounts-account-id-load-balancers-pools-pool-id-preview-deaa3487
path: operations/account-load-balancer-pools
description: Preview pool health using provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/load_balancers/pools/{pool_id}/preview
operation_ids:
    - account-load-balancer-pools-preview-pool
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview Pool

`POST /accounts/{account_id}/load_balancers/pools/{pool_id}/preview`

Operation ID: `account-load-balancer-pools-preview-pool`

Preview pool health using provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.

## Definition

```yaml
{"operationId": "account-load-balancer-pools-preview-pool", "summary": "Preview Pool", "description": "Preview pool health using provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.", "parameters": [{"name": "pool_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/load-balancing_components-schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_monitor-editable"}]}}}}, "responses": {"200": {"description": "Preview Pool response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/load-balancing_preview_response"}}}}, "4XX": {"description": "Preview Pool response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/load-balancing_preview_response"}, {"$ref": "#/components/schemas/load-balancing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Load Balancer Pools"]}
```
