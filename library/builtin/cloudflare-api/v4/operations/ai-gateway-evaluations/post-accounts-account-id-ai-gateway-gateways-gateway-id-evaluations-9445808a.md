---
title: Create a new Evaluation
page_id: operation-post-accounts-account-id-ai-gateway-gateways-gateway-id-evaluations-e2e39288
path: operations/ai-gateway-evaluations
description: Creates a new AI Gateway.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/evaluations
operation_ids:
    - aig-config-create-evaluations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Evaluation

`POST /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/evaluations`

Operation ID: `aig-config-create-evaluations`

Creates a new AI Gateway.

## Definition

```yaml
{"operationId": "aig-config-create-evaluations", "summary": "Create a new Evaluation", "description": "Creates a new AI Gateway.", "parameters": [{"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"dataset_ids": {"type": "array", "items": {"type": "string"}, "maxItems": 5, "minItems": 1}, "evaluation_type_ids": {"type": "array", "items": {"type": "string"}}, "name": {"type": "string", "x-auditable": true}}, "required": ["name", "dataset_ids", "evaluation_type_ids"]}}}}, "responses": {"200": {"description": "Returns the created Object", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "datasets": {"type": "array", "items": {"properties": {"account_id": {"type": "string"}, "account_tag": {"type": "string"}, "created_at": {"type": "string", "format": "date-time"}, "enable": {"type": "boolean", "x-auditable": true}, "filters": {"type": "array", "items": {"properties": {"key": {"type": "string", "enum": ["created_at", "request_content_type", "response_content_type", "success", "cached", "provider", "model", "cost", "tokens", "tokens_in", "tokens_out", "duration", "feedback"], "x-auditable": true}, "operator": {"type": "string", "enum": ["eq", "contains", "lt", "gt"], "x-auditable": true}, "value": {"type": "array", "items": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}, "x-auditable": true}}, "required": ["key", "operator", "value"], "type": "object"}, "x-auditable": true}, "gateway_id": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string", "x-auditable": true}}, "required": ["gateway_id", "name", "filters", "enable", "id", "account_id", "account_tag", "created_at", "modified_at"], "type": "object"}}, "gateway_id": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string", "x-auditable": true}, "processed": {"type": "boolean", "x-auditable": true}, "results": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "evaluation_id": {"type": "string"}, "evaluation_type_id": {"type": "string"}, "id": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "result": {"type": "string"}, "status": {"type": "number"}, "status_description": {"type": "string"}, "total_logs": {"type": "number"}}, "required": ["evaluation_id", "evaluation_type_id", "result", "total_logs", "status", "status_description", "id", "created_at", "modified_at"], "type": "object"}}, "total_logs": {"type": "number", "x-auditable": true}}, "required": ["id", "gateway_id", "name", "created_at", "modified_at", "datasets", "results", "processed", "total_logs"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7001}, "message": {"type": "string", "example": "Input Validation Error"}, "path": {"type": "array", "items": {"example": "body", "type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Evaluations"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
