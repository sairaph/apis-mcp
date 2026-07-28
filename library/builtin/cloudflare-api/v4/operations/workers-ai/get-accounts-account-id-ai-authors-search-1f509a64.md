---
title: Author Search
page_id: operation-get-accounts-account-id-ai-authors-search-f7bbf87a
path: operations/workers-ai
description: Searches Workers AI models by author or organization name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai/authors/search
operation_ids:
    - workers-ai-search-author
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Author Search

`GET /accounts/{account_id}/ai/authors/search`

Operation ID: `workers-ai-search-author`

Searches Workers AI models by author or organization name.

## Definition

```yaml
{"operationId": "workers-ai-search-author", "summary": "Author Search", "description": "Searches Workers AI models by author or organization name.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353"}}], "responses": {"200": {"description": "Returns a list of authors", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}, "400": {"description": "Bad Request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers AI"], "x-api-token-group": ["Workers AI Write", "Workers AI Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.ai"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
