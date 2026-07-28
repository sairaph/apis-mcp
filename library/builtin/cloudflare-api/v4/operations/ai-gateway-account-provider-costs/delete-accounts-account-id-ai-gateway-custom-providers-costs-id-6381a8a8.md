---
title: Delete a Account Provider Cost
page_id: operation-delete-accounts-account-id-ai-gateway-custom-providers-costs-id-397bde90
path: operations/ai-gateway-account-provider-costs
description: Deletes an AI Gateway dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-gateway/custom-providers/costs/{id}
operation_ids:
    - aig-config-delete-account-provider-cost
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Account Provider Cost

`DELETE /accounts/{account_id}/ai-gateway/custom-providers/costs/{id}`

Operation ID: `aig-config-delete-account-provider-cost`

Deletes an AI Gateway dataset.

## Definition

```yaml
{"operationId": "aig-config-delete-account-provider-cost", "summary": "Delete a Account Provider Cost", "description": "Deletes an AI Gateway dataset.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "3ebbcb006d4d46d7bb6a8c7f14676cb0"}}, {"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns the Object if it was successfully deleted", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"account_provider_id": {"type": "string", "format": "uuid"}, "changed_by": {"type": "string", "default": "manual"}, "cost_in": {"type": "number"}, "cost_out": {"type": "number"}, "cost_type": {"type": "string", "default": "tokens"}, "created_at": {"type": "string", "format": "date-time"}, "enable": {"type": "boolean"}, "id": {"type": "string", "format": "uuid"}, "model": {"type": "string"}, "model_rule": {"type": "string", "default": "equals", "enum": ["equals", "starts-with", "contains"]}, "modified_at": {"type": "string", "format": "date-time"}, "token_pricing": {"type": "string"}, "weight": {"type": "integer"}}, "required": ["account_provider_id", "model", "id", "created_at", "modified_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway Account Provider Costs"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
