---
title: Get all converted formats supported
page_id: operation-get-accounts-account-id-ai-tomarkdown-supported-873b241b
path: operations/workers-ai
description: Lists all file formats supported for conversion to Markdown.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/tomarkdown/supported
operation_ids:
    - workers-ai-get-to-markdown-supported
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get all converted formats supported

`GET /accounts/{account_id}/ai/tomarkdown/supported`

Operation ID: `workers-ai-get-to-markdown-supported`

Lists all file formats supported for conversion to Markdown.

## Definition

```yaml
{"operationId": "workers-ai-get-to-markdown-supported", "summary": "Get all converted formats supported", "description": "Lists all file formats supported for conversion to Markdown.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"extension": {"type": "string", "x-auditable": true}, "mimeType": {"type": "string", "x-auditable": true}}, "required": ["extension", "mimeType"], "type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
