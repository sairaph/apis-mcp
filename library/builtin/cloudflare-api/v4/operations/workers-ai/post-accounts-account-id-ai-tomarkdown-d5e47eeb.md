---
title: Convert Files into Markdown
page_id: operation-post-accounts-account-id-ai-tomarkdown-ce98c604
path: operations/workers-ai
description: Converts uploaded files into Markdown format using Workers AI.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai/tomarkdown
operation_ids:
    - workers-ai-post-to-markdown
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Convert Files into Markdown

`POST /accounts/{account_id}/ai/tomarkdown`

Operation ID: `workers-ai-post-to-markdown`

Converts uploaded files into Markdown format using Workers AI.

## Definition

```yaml
{"operationId": "workers-ai-post-to-markdown", "summary": "Convert Files into Markdown", "description": "Converts uploaded files into Markdown format using Workers AI.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "requestBody": {"content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"files": {"type": "array", "items": {"format": "binary", "type": "string"}}}, "required": ["files"]}}}}, "responses": {"200": {"description": "Model Schema", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"data": {"type": "string", "x-auditable": true}, "format": {"type": "string", "x-auditable": true}, "mimeType": {"type": "string", "x-auditable": true}, "name": {"type": "string", "x-auditable": true}, "tokens": {"type": "string", "x-auditable": true}}, "required": ["name", "mimeType", "format", "tokens", "data"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
