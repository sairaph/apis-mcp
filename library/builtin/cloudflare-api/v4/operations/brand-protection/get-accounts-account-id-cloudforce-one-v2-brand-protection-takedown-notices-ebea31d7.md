---
title: List takedown letters
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-noti-2ea18cf9
path: operations/brand-protection
description: List all letters for a specific takedown notice.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters
operation_ids:
    - get_TakedownLetterList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List takedown letters

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters`

Operation ID: `get_TakedownLetterList`

List all letters for a specific takedown notice.

## Definition

```yaml
{"operationId": "get_TakedownLetterList", "summary": "List takedown letters", "description": "List all letters for a specific takedown notice.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "responses": {"200": {"description": "Takedown letters listed successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"letters": {"type": "array", "items": {"properties": {"createdAt": {"type": "string", "nullable": true}, "id": {"type": "number"}, "letterType": {"type": "string"}, "takedownNoticeId": {"type": "number"}, "templateId": {"type": "string"}}, "required": ["id", "takedownNoticeId", "templateId", "letterType", "createdAt"], "type": "object"}}}, "required": ["letters"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
