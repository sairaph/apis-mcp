---
title: Create a new Provider Configs
page_id: operation-post-accounts-account-id-ai-gateway-gateways-gateway-id-provider-configs-02744740
path: operations/ai-gateway-provider-configs
description: Creates a new AI Gateway.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/provider_configs
operation_ids:
    - aig-config-create-providers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Provider Configs

`POST /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/provider_configs`

Operation ID: `aig-config-create-providers`

Creates a new AI Gateway.

## Definition

```yaml
{"operationId": "aig-config-create-providers", "summary": "Create a new Provider Configs", "description": "Creates a new AI Gateway.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"alias": {"type": "string"}, "default_config": {"type": "boolean"}, "provider_slug": {"type": "string"}, "rate_limit": {"type": "number"}, "rate_limit_period": {"type": "number", "default": 60}, "secret": {"type": "string"}, "secret_id": {"type": "string"}}, "required": ["provider_slug", "default_config", "alias"]}}}}, "responses": {"200": {"description": "Returns the created Object", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"alias": {"type": "string"}, "default_config": {"type": "boolean"}, "gateway_id": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}, "id": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "provider_slug": {"type": "string"}, "rate_limit": {"type": "number"}, "rate_limit_period": {"type": "number", "default": 60}, "secret_id": {"type": "string"}, "secret_preview": {"type": "string"}}, "required": ["id", "provider_slug", "secret_preview", "default_config", "gateway_id", "modified_at", "alias", "secret_id"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7001}, "message": {"type": "string", "example": "Input Validation Error"}, "path": {"type": "array", "items": {"example": "body", "type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Provider Configs"], "x-api-token-group": ["Secrets Store Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
