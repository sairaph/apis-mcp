---
title: List all AI Gateway Dynamic Route Versions.
page_id: operation-get-accounts-account-id-ai-gateway-gateways-gateway-id-routes-id-version-ea93f8be
path: operations/ai-gateway-dynamic-routes
description: List all AI Gateway Dynamic Route Versions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/routes/{id}/versions
operation_ids:
    - aig-config-list-gateway-dynamic-route-versions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all AI Gateway Dynamic Route Versions.

`GET /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/routes/{id}/versions`

Operation ID: `aig-config-list-gateway-dynamic-route-versions`

List all AI Gateway Dynamic Route Versions.

## Definition

```yaml
{"operationId": "aig-config-list-gateway-dynamic-route-versions", "summary": "List all AI Gateway Dynamic Route Versions.", "description": "List all AI Gateway Dynamic Route Versions.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}], "responses": {"200": {"description": "Success", "content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"order_by": {"type": "string"}, "order_by_direction": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "versions": {"type": "array", "items": {"properties": {"active": {"type": "string", "enum": ["true", "false"]}, "created_at": {"type": "string"}, "data": {"type": "string"}, "is_valid": {"type": "boolean"}, "version_id": {"type": "string"}}, "required": ["version_id", "data", "active", "created_at"], "type": "object"}}}, "required": ["versions", "page", "per_page", "order_by", "order_by_direction"]}, "success": {"type": "boolean"}}, "required": ["success", "data"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Dynamic Routes"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.dynamic-routing", "x-fern-sdk-method-name": "list-versions", "x-forge-hidden": true}
```
