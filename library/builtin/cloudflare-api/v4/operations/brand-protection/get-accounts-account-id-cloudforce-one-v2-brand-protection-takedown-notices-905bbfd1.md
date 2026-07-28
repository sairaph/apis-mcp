---
title: List takedown notices
page_id: operation-get-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-noti-e9e5d341
path: operations/brand-protection
description: List all takedown notices for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices
operation_ids:
    - get_TakedownNoticeList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List takedown notices

`GET /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices`

Operation ID: `get_TakedownNoticeList`

List all takedown notices for the account.

## Definition

```yaml
{"operationId": "get_TakedownNoticeList", "summary": "List takedown notices", "description": "List all takedown notices for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Takedown notices listed successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"notices": {"type": "array", "items": {"properties": {"createdAt": {"type": "string", "nullable": true}, "domain": {"type": "string"}, "id": {"type": "number"}, "matchId": {"type": "number", "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "number", "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "domain", "queryId", "status", "matchId", "matchType", "createdAt", "updatedAt"], "type": "object"}}}, "required": ["notices"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
