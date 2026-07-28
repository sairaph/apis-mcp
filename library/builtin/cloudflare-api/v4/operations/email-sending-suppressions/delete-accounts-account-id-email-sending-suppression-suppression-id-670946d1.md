---
title: Delete account email suppression
page_id: operation-delete-accounts-account-id-email-sending-suppression-suppression-id-9cf9e876
path: operations/email-sending-suppressions
description: Deletes an email suppression for the specified account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/email/sending/suppression/{suppression_id}
operation_ids:
    - delete_publicDeleteSuppressionSending
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete account email suppression

`DELETE /accounts/{account_id}/email/sending/suppression/{suppression_id}`

Operation ID: `delete_publicDeleteSuppressionSending`

Deletes an email suppression for the specified account.

## Definition

```yaml
{"operationId": "delete_publicDeleteSuppressionSending", "summary": "Delete account email suppression", "description": "Deletes an email suppression for the specified account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "suppression_id", "in": "path", "required": true, "schema": {"type": "string", "example": "396a5436-d4b0-42a6-b3fc-48e8fa522321"}}], "responses": {"200": {"description": "Deletes suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"success": {"type": "boolean"}}, "required": ["success"]}}}}, "404": {"description": "Suppression not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "error"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
