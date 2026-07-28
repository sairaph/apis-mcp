---
title: Get Model Schema
page_id: operation-get-accounts-account-id-ai-models-schema-3f8a5249
path: operations/workers-ai
description: Retrieves the input and output JSON schema definition for a Workers AI model.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/models/schema
operation_ids:
    - workers-ai-get-model-schema
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Model Schema

`GET /accounts/{account_id}/ai/models/schema`

Operation ID: `workers-ai-get-model-schema`

Retrieves the input and output JSON schema definition for a Workers AI model.

## Definition

```yaml
{"operationId": "workers-ai-get-model-schema", "summary": "Get Model Schema", "description": "Retrieves the input and output JSON schema definition for a Workers AI model.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}, {"name": "model", "in": "query", "description": "Model Name", "required": true, "schema": {"description": "Model Name", "type": "string"}}], "responses": {"200": {"description": "Model Schema", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"input": {"type": "object", "properties": {"additionalProperties": {"type": "boolean", "example": true}, "description": {"type": "string", "example": "JSON Schema definition for the model's input parameters"}, "type": {"type": "string", "example": "object"}}, "required": ["type", "description", "additionalProperties"]}, "output": {"type": "object", "properties": {"additionalProperties": {"type": "boolean", "example": true}, "description": {"type": "string", "example": "JSON Schema definition for the model's output format"}, "type": {"type": "string", "example": "object"}}, "required": ["type", "description", "additionalProperties"]}}, "required": ["input", "output"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
