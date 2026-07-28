---
title: Create account email suppression
page_id: operation-post-accounts-account-id-email-sending-suppression-4aab8ad2
path: operations/email-sending-suppressions
description: Creates a new email suppression for the specified account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email/sending/suppression
operation_ids:
    - post_publicNewSuppressionSending
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create account email suppression

`POST /accounts/{account_id}/email/sending/suppression`

Operation ID: `post_publicNewSuppressionSending`

Creates a new email suppression for the specified account.

## Definition

```yaml
{"operationId": "post_publicNewSuppressionSending", "summary": "Create account email suppression", "description": "Creates a new email suppression for the specified account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}}, "required": ["email"]}}}}, "responses": {"200": {"description": "Returns suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"type": "string", "format": "uuid"}}, "required": ["id"]}}, "required": ["result"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "string"}}, "success": {"type": "boolean"}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
