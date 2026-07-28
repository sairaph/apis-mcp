---
title: List Provider Configs
page_id: operation-get-accounts-account-id-ai-gateway-gateways-gateway-id-provider-configs-2a118d8f
path: operations/ai-gateway-provider-configs
description: Lists all AI Gateway evaluator types configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/provider_configs
operation_ids:
    - aig-config-list-providers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Provider Configs

`GET /accounts/{account_id}/ai-gateway/gateways/{gateway_id}/provider_configs`

Operation ID: `aig-config-list-providers`

Lists all AI Gateway evaluator types configured for the account.

## Definition

```yaml
{"operationId": "aig-config-list-providers", "summary": "List Provider Configs", "description": "Lists all AI Gateway evaluator types configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "gateway_id", "in": "path", "required": true, "schema": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List objects", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"alias": {"type": "string"}, "default_config": {"type": "boolean"}, "gateway_id": {"description": "gateway id", "type": "string", "example": "my-gateway", "maxLength": 64, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$", "x-auditable": true}, "id": {"type": "string"}, "modified_at": {"type": "string", "format": "date-time"}, "provider_slug": {"type": "string"}, "rate_limit": {"type": "number"}, "rate_limit_period": {"type": "number", "default": 60}, "secret_id": {"type": "string"}, "secret_preview": {"type": "string"}}, "required": ["id", "provider_slug", "secret_preview", "default_config", "gateway_id", "modified_at", "alias", "secret_id"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Provider Configs"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
