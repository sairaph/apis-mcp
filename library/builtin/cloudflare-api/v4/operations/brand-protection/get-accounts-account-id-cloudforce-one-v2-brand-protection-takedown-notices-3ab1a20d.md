---
title: Get takedown notice
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-noti-8c72690a
path: operations/brand-protection
description: Get a specific takedown notice by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}
operation_ids:
    - get_TakedownNoticeGet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get takedown notice

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}`

Operation ID: `get_TakedownNoticeGet`

Get a specific takedown notice by ID.

## Definition

```yaml
{"operationId": "get_TakedownNoticeGet", "summary": "Get takedown notice", "description": "Get a specific takedown notice by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "responses": {"200": {"description": "Takedown notice retrieved successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "nullable": true}, "domain": {"type": "string"}, "id": {"type": "number"}, "matchId": {"type": "number", "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "number", "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "domain", "queryId", "status", "matchId", "matchType", "createdAt", "updatedAt"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
