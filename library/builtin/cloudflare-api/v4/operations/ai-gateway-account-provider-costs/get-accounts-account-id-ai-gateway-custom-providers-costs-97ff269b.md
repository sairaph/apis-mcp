---
title: List Account Provider Costs
page_id: operation-get-accounts-account-id-ai-gateway-custom-providers-costs-248025a5
path: operations/ai-gateway-account-provider-costs
description: Lists all AI Gateway evaluator types configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/custom-providers/costs
operation_ids:
    - aig-config-list-account-provider-cost
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Account Provider Costs

`GET /accounts/{account_id}/ai-gateway/custom-providers/costs`

Operation ID: `aig-config-list-account-provider-cost`

Lists all AI Gateway evaluator types configured for the account.

## Definition

```yaml
{"operationId": "aig-config-list-account-provider-cost", "summary": "List Account Provider Costs", "description": "Lists all AI Gateway evaluator types configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}, {"name": "enable", "in": "query", "schema": {"type": "boolean"}}, {"name": "account_provider_id", "in": "query", "schema": {"type": "string", "format": "uuid"}}, {"name": "model_rule", "in": "query", "schema": {"type": "string", "default": "equals", "enum": ["equals", "starts-with", "contains"]}}, {"name": "cost_type", "in": "query", "schema": {"type": "string", "default": "tokens"}}, {"name": "search", "in": "query", "schema": {"description": "Search by model, changed_by", "type": "string"}}], "responses": {"200": {"description": "List objects", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"account_provider_id": {"type": "string", "format": "uuid"}, "changed_by": {"type": "string", "default": "manual"}, "cost_in": {"type": "number"}, "cost_out": {"type": "number"}, "cost_type": {"type": "string", "default": "tokens"}, "created_at": {"type": "string", "format": "date-time"}, "enable": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "model": {"type": "string"}, "model_rule": {"type": "string", "default": "equals", "enum": ["equals", "starts-with", "contains"]}, "modified_at": {"type": "string", "format": "date-time"}, "token_pricing": {"type": "string"}, "weight": {"type": "integer"}}, "required": ["account_provider_id", "model", "id", "created_at", "modified_at"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Account Provider Costs"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
