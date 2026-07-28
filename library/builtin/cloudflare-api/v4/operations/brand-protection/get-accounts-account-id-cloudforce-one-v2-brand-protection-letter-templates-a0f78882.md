---
title: List letter templates
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-letter-templa-6d2a93f6
path: operations/brand-protection
description: List user-defined takedown letter templates. System example templates are available via the /templates/examples endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates
operation_ids:
    - get_LetterTemplateList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List letter templates

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates`

Operation ID: `get_LetterTemplateList`

List user-defined takedown letter templates. System example templates are available via the /templates/examples endpoint.

## Definition

```yaml
{"operationId": "get_LetterTemplateList", "summary": "List letter templates", "description": "List user-defined takedown letter templates. System example templates are available via the /templates/examples endpoint.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Templates listed successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"templates": {"type": "array", "items": {"properties": {"category": {"type": "string"}, "createdAt": {"type": "string", "nullable": true}, "description": {"type": "string", "nullable": true}, "id": {"type": "string"}, "name": {"type": "string"}, "source": {"type": "string", "enum": ["system", "user"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "source", "name", "category", "description", "createdAt", "updatedAt"], "type": "object"}}}, "required": ["templates"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
