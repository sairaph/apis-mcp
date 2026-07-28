---
title: Patch Gateway Log
page_id: operation-patch-accounts-account-id-ai-gateway-gateways-gateway-id-logs-id-003aadbc
path: operations/ai-gateway-logs
description: Updates metadata for an AI Gateway log entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}
operation_ids:
    - aig-config-patch-gateway-log
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Gateway Log

`PATCH /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}`

Operation ID: `aig-config-patch-gateway-log`

Updates metadata for an AI Gateway log entry.

## Definition

```yaml
{"operationId": "aig-config-patch-gateway-log", "summary": "Patch Gateway Log", "description": "Updates metadata for an AI Gateway log entry.", "parameters": [{"name": "id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"feedback": {"type": "number", "maximum": 1, "minimum": -1, "nullable": true}, "metadata": {"type": "object", "additionalProperties": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}, "nullable": true}, "score": {"type": "number", "maximum": 100, "minimum": 0, "nullable": true}}}}}}, "responses": {"200": {"description": "Returns the log details", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Logs"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.logs", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
