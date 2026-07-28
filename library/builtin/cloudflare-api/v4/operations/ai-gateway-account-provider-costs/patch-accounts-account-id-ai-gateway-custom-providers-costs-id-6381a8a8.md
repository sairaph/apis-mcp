---
title: Update a Account Provider Cost
page_id: operation-patch-accounts-account-id-ai-gateway-custom-providers-costs-id-30655d70
path: operations/ai-gateway-account-provider-costs
description: Updates an existing AI Gateway dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/ai-gateway/custom-providers/costs/{id}
operation_ids:
    - aig-config-update-account-provider-cost
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Account Provider Cost

`PATCH /accounts/{account_id}/ai-gateway/custom-providers/costs/{id}`

Operation ID: `aig-config-update-account-provider-cost`

Updates an existing AI Gateway dataset.

## Definition

```yaml
{"operationId": "aig-config-update-account-provider-cost", "summary": "Update a Account Provider Cost", "description": "Updates an existing AI Gateway dataset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"cost_in": {"type": "number"}, "cost_out": {"type": "number"}, "cost_type": {"type": "string", "default": "tokens"}, "enable": {"type": "boolean"}, "model": {"type": "string"}, "model_rule": {"type": "string", "default": "equals", "enum": ["equals", "starts-with", "contains"]}, "token_pricing": {"type": "object", "properties": {"input_audio_tokens": {"type": "number"}, "input_cache_creation_tokens": {"type": "number"}, "input_cached_tokens": {"type": "number"}, "input_image_count": {"type": "number"}, "input_image_tokens": {"type": "number"}, "input_text_tokens": {"type": "number"}, "input_tokens": {"type": "number"}, "input_video_tokens": {"type": "number"}, "output_image_count": {"type": "number"}, "output_reasoning_tokens": {"type": "number"}, "output_tokens": {"type": "number"}, "total_tokens": {"type": "number"}}}}}}}}, "responses": {"200": {"description": "Returns the updated Object", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"account_provider_id": {"type": "string", "format": "uuid"}, "changed_by": {"type": "string", "default": "manual"}, "cost_in": {"type": "number"}, "cost_out": {"type": "number"}, "cost_type": {"type": "string", "default": "tokens"}, "created_at": {"type": "string", "format": "date-time"}, "enable": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "model": {"type": "string"}, "model_rule": {"type": "string", "default": "equals", "enum": ["equals", "starts-with", "contains"]}, "modified_at": {"type": "string", "format": "date-time"}, "token_pricing": {"type": "string"}, "weight": {"type": "integer"}}, "required": ["account_provider_id", "model", "id", "created_at", "modified_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Input Validation Error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7001}, "message": {"type": "string", "example": "Input Validation Error"}, "path": {"type": "array", "items": {"example": "body", "type": "string"}}}, "required": ["code", "message", "path"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Account Provider Costs"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
