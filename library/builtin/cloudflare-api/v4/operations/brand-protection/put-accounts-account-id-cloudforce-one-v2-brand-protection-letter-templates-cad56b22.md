---
title: Update letter template
page_id: operation-put-accounts-account-id-cloudforce-one-v2-brand-protection-letter-templa-c8bc1408
path: operations/brand-protection
description: Update a user-defined takedown letter template. System templates cannot be modified.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/{template_id}
operation_ids:
    - put_LetterTemplateUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update letter template

`PUT /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/{template_id}`

Operation ID: `put_LetterTemplateUpdate`

Update a user-defined takedown letter template. System templates cannot be modified.

## Definition

```yaml
{"operationId": "put_LetterTemplateUpdate", "summary": "Update letter template", "description": "Update a user-defined takedown letter template. System templates cannot be modified.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "template_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"body": {"type": "string", "minLength": 1}, "category": {"type": "string", "maxLength": 100, "minLength": 1}, "description": {"type": "string", "maxLength": 1000, "nullable": true}, "name": {"type": "string", "maxLength": 255, "minLength": 1}}}}}}, "responses": {"200": {"description": "Template updated successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"body": {"type": "string"}, "category": {"type": "string"}, "createdAt": {"type": "string", "nullable": true}, "description": {"type": "string", "nullable": true}, "id": {"type": "string"}, "name": {"type": "string"}, "source": {"type": "string", "enum": ["system", "user"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "source", "name", "category", "description", "body", "createdAt", "updatedAt"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
