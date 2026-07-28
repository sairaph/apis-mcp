---
title: List example letter templates
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-letter-templa-406f59ea
path: operations/brand-protection
description: List system-provided example templates that can be used as starting points when creating custom templates. These templates cannot be used directly for letter generation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/examples
operation_ids:
    - get_LetterTemplateExamples
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List example letter templates

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/examples`

Operation ID: `get_LetterTemplateExamples`

List system-provided example templates that can be used as starting points when creating custom templates. These templates cannot be used directly for letter generation.

## Definition

```yaml
{"operationId": "get_LetterTemplateExamples", "summary": "List example letter templates", "description": "List system-provided example templates that can be used as starting points when creating custom templates. These templates cannot be used directly for letter generation.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Example templates listed successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"templates": {"type": "array", "items": {"properties": {"body": {"type": "string"}, "category": {"type": "string"}, "createdAt": {"type": "string", "nullable": true}, "description": {"type": "string", "nullable": true}, "id": {"type": "string"}, "name": {"type": "string"}, "source": {"type": "string", "enum": ["system", "user"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "source", "name", "category", "description", "body", "createdAt", "updatedAt"], "type": "object"}}}, "required": ["templates"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
