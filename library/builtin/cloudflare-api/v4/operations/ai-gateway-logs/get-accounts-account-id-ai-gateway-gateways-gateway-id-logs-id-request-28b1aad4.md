---
title: Get Gateway Log Request
page_id: operation-get-accounts-account-id-ai-gateway-gateways-gateway-id-logs-id-request-f8fadc57
path: operations/ai-gateway-logs
description: Retrieves the original request payload for an AI Gateway log entry.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}/request
operation_ids:
    - aig-config-get-gateway-log-request
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Gateway Log Request

`GET /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs/{id}/request`

Operation ID: `aig-config-get-gateway-log-request`

Retrieves the original request payload for an AI Gateway log entry.

## Definition

```yaml
{"operationId": "aig-config-get-gateway-log-request", "summary": "Get Gateway Log Request", "description": "Retrieves the original request payload for an AI Gateway log entry.", "parameters": [{"name": "id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}], "responses": {"200": {"description": "Returns the request body from a specific log", "content": {"application/json": {"schema": {"type": "object"}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Logs"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.logs", "x-fern-sdk-method-name": "request", "x-forge-hidden": true}
```
