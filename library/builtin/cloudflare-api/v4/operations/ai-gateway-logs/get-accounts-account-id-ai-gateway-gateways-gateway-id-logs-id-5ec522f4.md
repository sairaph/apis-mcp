---
title: Get Gateway Log Detail
page_id: operation-get-accounts-account-id-ai-gateway-gateways-gateway-id-logs-id-c8d8d984
path: operations/ai-gateway-logs
description: Retrieves detailed information for a specific AI Gateway log entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}
operation_ids:
    - aig-config-get-gateway-log-detail
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Gateway Log Detail

`GET /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}`

Operation ID: `aig-config-get-gateway-log-detail`

Retrieves detailed information for a specific AI Gateway log entry.

## Definition

```yaml
{"operationId": "aig-config-get-gateway-log-detail", "summary": "Get Gateway Log Detail", "description": "Retrieves detailed information for a specific AI Gateway log entry.", "parameters": [{"name": "id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}], "responses": {"200": {"description": "Returns the log details", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"cached": {"type": "boolean"}, "cost": {"type": "number"}, "created_at": {"type": "string", "format": "date-time"}, "custom_cost": {"type": "boolean"}, "duration": {"type": "integer"}, "id": {"type": "string"}, "metadata": {"type": "string"}, "model": {"type": "string"}, "model_type": {"type": "string"}, "path": {"type": "string"}, "provider": {"type": "string"}, "request_content_type": {"type": "string"}, "request_head": {"type": "string"}, "request_head_complete": {"type": "boolean"}, "request_size": {"type": "integer"}, "request_type": {"type": "string"}, "response_content_type": {"type": "string"}, "response_head": {"type": "string"}, "response_head_complete": {"type": "boolean"}, "response_size": {"type": "integer"}, "status_code": {"type": "integer"}, "step": {"type": "integer"}, "success": {"type": "boolean"}, "tokens_in": {"type": "integer", "nullable": true}, "tokens_out": {"type": "integer", "nullable": true}}, "required": ["id", "created_at", "provider", "model", "path", "duration", "success", "cached", "tokens_in", "tokens_out"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Logs"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.logs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
