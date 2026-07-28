---
title: Delete a Dataset
page_id: operation-delete-accounts-account-id-ai-gateway-gateways-gateway-id-datasets-id-c1a34190
path: operations/ai-gateway-datasets
description: Deletes an AI Gateway dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets/{id}
operation_ids:
    - aig-config-delete-dataset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Dataset

`DELETE /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/datasets/{id}`

Operation ID: `aig-config-delete-dataset`

Deletes an AI Gateway dataset.

## Definition

```yaml
{"operationId": "aig-config-delete-dataset", "summary": "Delete a Dataset", "description": "Deletes an AI Gateway dataset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "x-auditable": true}}], "responses": {"200": {"description": "Returns the Object if it was successfully deleted", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "enable": {"type": "boolean", "x-auditable": true}, "filters": {"type": "array", "items": {"properties": {"key": {"type": "string", "enum": ["created_at", "request_content_type", "response_content_type", "success", "cached", "provider", "model", "cost", "tokens", "tokens_in", "tokens_out", "duration", "feedback"], "x-auditable": true}, "operator": {"type": "string", "enum": ["eq", "contains", "lt", "gt"], "x-auditable": true}, "value": {"type": "array", "items": {"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}, "x-auditable": true}}, "required": ["key", "operator", "value"], "type": "object"}, "x-auditable": true}, "gateway_id": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string", "x-auditable": true}}, "required": ["gateway_id", "name", "filters", "enable", "id", "created_at", "modified_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Datasets"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
