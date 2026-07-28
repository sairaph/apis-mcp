---
title: Get account Email Sending suppression
page_id: operation-get-accounts-account-id-email-sending-suppressions-suppression-id-67528cbc
path: operations/email-sending-suppressions
description: Get account Email Sending suppression
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/sending/suppressions/{suppression_id}
operation_ids:
    - get_publicGetSendingSuppression
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account Email Sending suppression

`GET /accounts/{account_id}/email/sending/suppressions/{suppression_id}`

Operation ID: `get_publicGetSendingSuppression`

Get account Email Sending suppression

## Definition

```yaml
{"operationId": "get_publicGetSendingSuppression", "summary": "Get account Email Sending suppression", "description": "Get account Email Sending suppression", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "suppression_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid", "example": "396a5436-d4b0-42a6-b3fc-48e8fa522321"}}], "responses": {"200": {"description": "Returns account-level Email Sending suppression.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"nullable": true}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "reason": {"type": "string", "example": "hard_bounce"}}, "required": ["id", "email", "reason", "created_at", "expires_at"]}, "success": {"type": "boolean", "example": true}}, "required": ["errors", "messages", "success", "result"]}}}}, "404": {"description": "Suppression not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Suppression not found"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["errors", "messages", "result", "success"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
