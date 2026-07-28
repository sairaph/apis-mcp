---
title: Delete Gateway Logs
page_id: operation-delete-accounts-account-id-ai-gateway-gateways-gateway-id-logs-ef2e6851
path: operations/ai-gateway-logs
description: Deletes gateway log entries matching the specified criteria.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs
operation_ids:
    - aig-config-delete-gateway-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Gateway Logs

`DELETE /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/logs`

Operation ID: `aig-config-delete-gateway-logs`

Deletes gateway log entries matching the specified criteria.

## Definition

```yaml
{"operationId": "aig-config-delete-gateway-logs", "summary": "Delete Gateway Logs", "description": "Deletes gateway log entries matching the specified criteria.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "order_by", "in": "query", "schema": {"type": "string", "default": "created_at", "enum": ["created_at", "provider", "model", "model_type", "success", "cached", "cost", "tokens_in", "tokens_out", "duration", "feedback"]}}, {"name": "order_by_direction", "in": "query", "schema": {"type": "string", "default": "asc", "enum": ["asc", "desc"]}}, {"name": "filters", "in": "query", "schema": {"type": "array", "items": {"properties": {"key": {"type": "string", "enum": ["id", "created_at", "request_content_type", "response_content_type", "request_type", "success", "cached", "provider", "model", "model_type", "cost", "tokens", "tokens_in", "tokens_out", "duration", "feedback", "event_id", "metadata.key", "metadata.value", "authentication", "wholesale", "compatibilityMode", "dlp_action", "user_agent"]}, "operator": {"type": "string", "enum": ["eq", "neq", "contains", "lt", "gt"]}, "value": {"type": "array", "items": {"anyOf": [{"nullable": true, "type": "string"}, {"type": "number"}, {"type": "boolean"}]}}}, "required": ["key", "operator", "value"], "type": "object"}}}, {"name": "limit", "in": "query", "schema": {"type": "integer", "default": 10000, "maximum": 10000, "minimum": 1}}], "responses": {"200": {"description": "Returns if the delete was successful", "content": {"application/json": {"schema": {"type": "object", "properties": {"success": {"type": "boolean"}}, "required": ["success"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Logs"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.logs", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
