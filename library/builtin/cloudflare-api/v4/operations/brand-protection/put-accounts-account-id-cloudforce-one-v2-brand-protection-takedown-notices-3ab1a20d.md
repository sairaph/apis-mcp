---
title: Update takedown notice
page_id: operation-put-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-noti-16dfab72
path: operations/brand-protection
description: Update a takedown notice (e.g. change status to sent, resolved, etc.).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}
operation_ids:
    - put_TakedownNoticeUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update takedown notice

`PUT /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}`

Operation ID: `put_TakedownNoticeUpdate`

Update a takedown notice (e.g. change status to sent, resolved, etc.).

## Definition

```yaml
{"operationId": "put_TakedownNoticeUpdate", "summary": "Update takedown notice", "description": "Update a takedown notice (e.g. change status to sent, resolved, etc.).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"domain": {"type": "string", "maxLength": 253, "minLength": 1}, "matchId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0, "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "integer", "exclusiveMinimum": true, "minimum": 0, "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}}}}}}, "responses": {"200": {"description": "Takedown notice updated successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "nullable": true}, "domain": {"type": "string"}, "id": {"type": "number"}, "matchId": {"type": "number", "nullable": true}, "matchType": {"type": "string", "enum": ["logo", "domain"], "nullable": true}, "queryId": {"type": "number", "nullable": true}, "status": {"type": "string", "enum": ["draft", "sent", "resolved", "expired"]}, "updatedAt": {"type": "string", "nullable": true}}, "required": ["id", "domain", "queryId", "status", "matchId", "matchType", "createdAt", "updatedAt"]}}}}, "409": {"description": "A takedown notice already exists for the target domain"}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
