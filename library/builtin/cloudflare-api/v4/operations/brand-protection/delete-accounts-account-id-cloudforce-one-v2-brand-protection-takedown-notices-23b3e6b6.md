---
title: Delete takedown letter
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-brand-protection-takedown-n-66428fe3
path: operations/brand-protection
description: Delete a specific takedown letter and its associated PDF from storage.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters/{letter_id}
operation_ids:
    - delete_TakedownLetterDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete takedown letter

`DELETE /accounts/{account_id}/cloudforce-one/v2/brand-protection/takedown-notices/{notice_id}/letters/{letter_id}`

Operation ID: `delete_TakedownLetterDelete`

Delete a specific takedown letter and its associated PDF from storage.

## Definition

```yaml
{"operationId": "delete_TakedownLetterDelete", "summary": "Delete takedown letter", "description": "Delete a specific takedown letter and its associated PDF from storage.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "notice_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}, {"name": "letter_id", "in": "path", "required": true, "schema": {"type": "integer", "exclusiveMinimum": true, "minimum": 0}}], "responses": {"200": {"description": "Takedown letter deleted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
