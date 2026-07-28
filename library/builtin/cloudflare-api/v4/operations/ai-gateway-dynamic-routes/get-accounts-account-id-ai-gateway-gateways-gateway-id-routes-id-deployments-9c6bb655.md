---
title: List all AI Gateway Dynamic Route Deployments.
page_id: operation-get-accounts-account-id-ai-gateway-gateways-gateway-id-routes-id-deploym-74475bd7
path: operations/ai-gateway-dynamic-routes
description: List all AI Gateway Dynamic Route Deployments.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/routes/{id}/deployments
operation_ids:
    - aig-config-list-gateway-dynamic-route-deployments
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all AI Gateway Dynamic Route Deployments.

`GET /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/routes/{id}/deployments`

Operation ID: `aig-config-list-gateway-dynamic-route-deployments`

List all AI Gateway Dynamic Route Deployments.

## Definition

```yaml
{"operationId": "aig-config-list-gateway-dynamic-route-deployments", "summary": "List all AI Gateway Dynamic Route Deployments.", "description": "List all AI Gateway Dynamic Route Deployments.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}], "responses": {"200": {"description": "Success", "content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"deployments": {"type": "array", "items": {"properties": {"created_at": {"type": "string"}, "deployment_id": {"type": "string"}, "version_id": {"type": "string"}}, "required": ["deployment_id", "version_id", "created_at"], "type": "object"}}, "order_by": {"type": "string"}, "order_by_direction": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}}, "required": ["deployments", "page", "per_page", "order_by", "order_by_direction"]}, "success": {"type": "boolean"}}, "required": ["success", "data"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Dynamic Routes"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.dynamic-routing", "x-fern-sdk-method-name": "list-deployments", "x-forge-hidden": true}
```
