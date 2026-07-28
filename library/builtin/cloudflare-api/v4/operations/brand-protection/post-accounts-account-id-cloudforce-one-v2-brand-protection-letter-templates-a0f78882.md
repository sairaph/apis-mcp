---
title: Create letter template
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-letter-templ-8782a3f0
path: operations/brand-protection
description: Create a new user-defined takedown letter template
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates
operation_ids:
    - post_LetterTemplateCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create letter template

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates`

Operation ID: `post_LetterTemplateCreate`

Create a new user-defined takedown letter template

## Definition

```yaml
{"operationId": "post_LetterTemplateCreate", "summary": "Create letter template", "description": "Create a new user-defined takedown letter template", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"body": {"type": "string", "minLength": 1}, "category": {"type": "string", "maxLength": 100, "minLength": 1}, "description": {"type": "string", "maxLength": 1000}, "name": {"type": "string", "maxLength": 255, "minLength": 1}}, "required": ["name", "category", "body"]}}}}, "responses": {"200": {"description": "Template created successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"body": {"type": "string"}, "category": {"type": "string"}, "createdAt": {"type": "string", "nullable": true}, "description": {"type": "string", "nullable": true}, "id": {"type": "string"}, "name": {"type": "string"}, "source": {"type": "string", "enum": ["system", "user"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "source", "name", "category", "description", "body", "createdAt", "updatedAt"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
