---
title: Create takedown letter
page_id: operation-post-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-not-686ceb6b
path: operations/brand-protection
description: Generate a takedown letter for a notice and persist the rendered text. The PDF is generated lazily on download.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters
operation_ids:
    - post_TakedownLetterCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create takedown letter

`POST /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters`

Operation ID: `post_TakedownLetterCreate`

Generate a takedown letter for a notice and persist the rendered text. The PDF is generated lazily on download.

## Definition

```yaml
{"operationId": "post_TakedownLetterCreate", "summary": "Create takedown letter", "description": "Generate a takedown letter for a notice and persist the rendered text. The PDF is generated lazily on download.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"fields": {"type": "object", "properties": {"discoveryDate": {"type": "string", "pattern": "^\\d{4}-\\d{2}-\\d{2}$"}, "domain": {"type": "string", "maxLength": 500}, "generationDate": {"type": "string", "pattern": "^\\d{4}-\\d{2}-\\d{2}$"}, "jurisdiction": {"type": "string", "maxLength": 500}, "registrantEmail": {"type": "string", "maxLength": 500}, "registrar": {"type": "string", "maxLength": 500}, "registrarEmail": {"type": "string", "maxLength": 500}, "resolutionByDate": {"type": "string", "pattern": "^\\d{4}-\\d{2}-\\d{2}$"}, "senderCompany": {"type": "string", "maxLength": 500}, "senderEmail": {"type": "string", "maxLength": 500}, "senderName": {"type": "string", "maxLength": 500}, "senderTitle": {"type": "string", "maxLength": 500}, "trademarkName": {"type": "string", "maxLength": 500}, "trademarkNumber": {"type": "string", "maxLength": 500}, "trademarkOwner": {"type": "string", "maxLength": 500}}}, "templateId": {"type": "string", "minLength": 1}}, "required": ["templateId", "fields"]}}}}, "responses": {"200": {"description": "Takedown letter created successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "nullable": true}, "id": {"type": "number"}, "letterText": {"type": "string"}, "letterType": {"type": "string"}, "takedownNoticeId": {"type": "number"}, "templateId": {"type": "string"}}, "required": ["id", "takedownNoticeId", "templateId", "letterType", "letterText", "createdAt"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
