---
title: Delete an AI Gateway Dynamic Route.
page_id: operation-delete-accounts-account-id-ai-gateway-gateways-gateway-id-routes-id-6e05d697
path: operations/ai-gateway-dynamic-routes
description: Delete an AI Gateway Dynamic Route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/routes/{id}
operation_ids:
    - aig-config-delete-gateway-dynamic-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an AI Gateway Dynamic Route.

`DELETE /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/routes/{id}`

Operation ID: `aig-config-delete-gateway-dynamic-route`

Delete an AI Gateway Dynamic Route.

## Definition

```yaml
{"operationId": "aig-config-delete-gateway-dynamic-route", "summary": "Delete an AI Gateway Dynamic Route.", "description": "Delete an AI Gateway Dynamic Route.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}], "responses": {"200": {"description": "Success", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "elements": {"type": "array", "items": {"oneOf": [{"properties": {"id": {"type": "string"}, "outputs": {"type": "object", "properties": {"next": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}}, "required": ["next"]}, "type": {"type": "string", "enum": ["start"]}}, "required": ["id", "outputs", "type"], "type": "object"}, {"properties": {"id": {"type": "string"}, "outputs": {"type": "object", "properties": {"false": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}, "true": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}}, "required": ["true", "false"]}, "properties": {"type": "object", "properties": {"conditions": {"type": "object"}}}, "type": {"type": "string", "enum": ["conditional"]}}, "required": ["id", "outputs", "type", "properties"], "type": "object"}, {"properties": {"id": {"type": "string"}, "outputs": {"type": "object", "additionalProperties": {"properties": {"elementId": {"type": "string"}}, "required": ["elementId"], "type": "object"}}, "type": {"type": "string", "enum": ["percentage"]}}, "required": ["id", "outputs", "type"], "type": "object"}, {"properties": {"id": {"type": "string"}, "outputs": {"type": "object", "properties": {"fallback": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}, "success": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}}, "required": ["success", "fallback"]}, "properties": {"type": "object", "properties": {"key": {"type": "string"}, "limit": {"type": "number"}, "limitType": {"enum": ["count", "cost"], "type": "string"}, "window": {"type": "number"}}, "required": ["limitType", "key", "limit", "window"]}, "type": {"type": "string", "enum": ["rate"]}}, "required": ["id", "outputs", "type", "properties"], "type": "object"}, {"properties": {"id": {"type": "string"}, "outputs": {"type": "object", "properties": {"fallback": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}, "success": {"type": "object", "properties": {"elementId": {"type": "string"}}, "required": ["elementId"]}}, "required": ["success", "fallback"]}, "properties": {"type": "object", "properties": {"model": {"type": "string"}, "provider": {"type": "string"}, "retries": {"type": "number"}, "timeout": {"type": "number"}}, "required": ["provider", "model", "timeout", "retries"]}, "type": {"type": "string", "enum": ["model"]}}, "required": ["id", "outputs", "type", "properties"], "type": "object"}, {"properties": {"id": {"type": "string"}, "outputs": {"type": "object", "additionalProperties": {"properties": {"elementId": {"type": "string"}}, "required": ["elementId"], "type": "object"}}, "type": {"type": "string", "enum": ["end"]}}, "required": ["id", "outputs", "type"], "type": "object"}]}}, "gateway_id": {"type": "string"}, "id": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string"}}, "required": ["id", "name", "elements", "created_at", "modified_at", "gateway_id"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Dynamic Routes"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.dynamic-routing", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
