---
title: Get account email suppression
page_id: operation-get-accounts-account-id-email-sending-suppression-suppression-id-7a3163f8
path: operations/email-sending-suppressions
description: Retrieves a single email suppression for the specified account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/sending/suppression/{suppression_id}
operation_ids:
    - get_publicGetSuppressionSending
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account email suppression

`GET /accounts/{account_id}/email/sending/suppression/{suppression_id}`

Operation ID: `get_publicGetSuppressionSending`

Retrieves a single email suppression for the specified account.

## Definition

```yaml
{"operationId": "get_publicGetSuppressionSending", "summary": "Get account email suppression", "description": "Retrieves a single email suppression for the specified account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "suppression_id", "in": "path", "required": true, "schema": {"type": "string", "example": "396a5436-d4b0-42a6-b3fc-48e8fa522321"}}], "responses": {"200": {"description": "Returns suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "reason": {"type": "string", "example": "hard_bounce"}, "zones": {"type": "array", "items": {"type": "string"}, "default": []}}, "required": ["id", "email", "reason", "created_at", "expires_at"]}}, "required": ["result"]}}}}, "404": {"description": "Suppression not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string"}, "success": {"type": "boolean"}}, "required": ["success", "error"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
