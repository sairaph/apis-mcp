---
title: Delete takedown notice
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-n-0189aa98
path: operations/brand-protection
description: Delete a takedown notice and all associated letters. PDFs are also removed from storage.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}
operation_ids:
    - delete_TakedownNoticeDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete takedown notice

`DELETE /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}`

Operation ID: `delete_TakedownNoticeDelete`

Delete a takedown notice and all associated letters. PDFs are also removed from storage.

## Definition

```yaml
{"operationId": "delete_TakedownNoticeDelete", "summary": "Delete takedown notice", "description": "Delete a takedown notice and all associated letters. PDFs are also removed from storage.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "responses": {"200": {"description": "Takedown notice deleted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
