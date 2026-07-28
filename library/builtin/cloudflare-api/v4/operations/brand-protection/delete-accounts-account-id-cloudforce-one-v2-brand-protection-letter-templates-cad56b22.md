---
title: Delete letter template
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-brand-protection-letter-tem-b051ce3f
path: operations/brand-protection
description: Delete a user-defined takedown letter template. System templates cannot be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/{template_id}
operation_ids:
    - delete_LetterTemplateDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete letter template

`DELETE /accounts/{account_id}/cloudforce-one/v2/brand-protection/letter/templates/{template_id}`

Operation ID: `delete_LetterTemplateDelete`

Delete a user-defined takedown letter template. System templates cannot be deleted.

## Definition

```yaml
{"operationId": "delete_LetterTemplateDelete", "summary": "Delete letter template", "description": "Delete a user-defined takedown letter template. System templates cannot be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}, {"name": "template_id", "in": "path", "required": true, "schema": {"type": "string", "minLength": 1}}], "responses": {"200": {"description": "Template deleted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "message"]}}}}}, "security": [{"api_token": []}], "tags": ["Brand Protection"], "x-api-token-group": ["Cloudforce One Write"]}
```
