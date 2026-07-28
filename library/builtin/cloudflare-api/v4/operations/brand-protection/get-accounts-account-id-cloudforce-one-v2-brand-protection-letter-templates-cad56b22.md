---
title: Get letter template
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-letter-templa-8501a74d
path: operations/brand-protection
description: Get a specific user-defined takedown letter template by ID. System example templates are available via the /templates/examples endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/{template_id}
operation_ids:
    - get_LetterTemplateGet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get letter template

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/{template_id}`

Operation ID: `get_LetterTemplateGet`

Get a specific user-defined takedown letter template by ID. System example templates are available via the /templates/examples endpoint.

## Definition

```yaml
{"operationId": "get_LetterTemplateGet", "summary": "Get letter template", "description": "Get a specific user-defined takedown letter template by ID. System example templates are available via the /templates/examples endpoint.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "template_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Template retrieved successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"body": {"type": "string"}, "category": {"type": "string"}, "createdAt": {"type": "string", "nullable": true}, "description": {"type": "string", "nullable": true}, "id": {"type": "string"}, "name": {"type": "string"}, "source": {"type": "string", "enum": ["system", "user"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "source", "name", "category", "description", "body", "createdAt", "updatedAt"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
