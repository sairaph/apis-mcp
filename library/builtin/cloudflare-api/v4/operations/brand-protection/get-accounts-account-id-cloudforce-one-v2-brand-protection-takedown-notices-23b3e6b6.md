---
title: Get takedown letter
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-noti-1bc4092c
path: operations/brand-protection
description: Get a specific takedown letter by ID, including the rendered text.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters/{letter_id}
operation_ids:
    - get_TakedownLetterGet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get takedown letter

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters/{letter_id}`

Operation ID: `get_TakedownLetterGet`

Get a specific takedown letter by ID, including the rendered text.

## Definition

```yaml
{"operationId": "get_TakedownLetterGet", "summary": "Get takedown letter", "description": "Get a specific takedown letter by ID, including the rendered text.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}, {"name": "letter_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "responses": {"200": {"description": "Takedown letter retrieved successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "nullable": true}, "id": {"type": "number"}, "letterText": {"type": "string"}, "letterType": {"type": "string"}, "takedownNoticeId": {"type": "number"}, "templateId": {"type": "string"}}, "required": ["id", "takedownNoticeId", "templateId", "letterType", "letterText", "createdAt"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
