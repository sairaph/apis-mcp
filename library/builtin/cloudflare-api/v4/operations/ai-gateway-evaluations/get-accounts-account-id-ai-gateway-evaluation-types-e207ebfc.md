---
title: List Evaluators
page_id: operation-get-accounts-account-id-ai-gateway-evaluation-types-5e39360c
path: operations/ai-gateway-evaluations
description: Lists all available evaluator types for scoring AI gateway responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/evaluation-types
operation_ids:
    - aig-config-list-evaluators
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Evaluators

`GET /accounts/{account_id}/ai-gateway/evaluation-types`

Operation ID: `aig-config-list-evaluators`

Lists all available evaluator types for scoring AI gateway responses.

## Definition

```yaml
{"operationId": "aig-config-list-evaluators", "summary": "List Evaluators", "description": "Lists all available evaluator types for scoring AI gateway responses.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "0d37909e38d3e99c29fa2cd343ac421a"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order_by", "in": "query", "schema": {"type": "string", "default": "mandatory"}}, {"name": "order_by_direction", "in": "query", "schema": {"type": "string", "default": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "Returns a list of Evaluators", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "description": {"type": "string"}, "enable": {"type": "boolean"}, "id": {"type": "string"}, "mandatory": {"type": "boolean"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"type": "string"}, "type": {"type": "string"}}, "required": ["name", "type", "mandatory", "description", "enable", "id", "created_at", "modified_at"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"type": "boolean"}}, "required": ["success", "result", "result_info"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Evaluations"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai-gateway.evaluation-types", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
