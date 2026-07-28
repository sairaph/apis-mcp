---
title: Bulk import account Email Sending suppressions
page_id: operation-post-accounts-account-id-email-sending-suppressions-bulk-e5e24efa
path: operations/email-sending-suppressions
description: Bulk import account Email Sending suppressions
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email/sending/suppressions/bulk
operation_ids:
    - post_publicBulkCreateSendingSuppressions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk import account Email Sending suppressions

`POST /accounts/{account_id}/email/sending/suppressions/bulk`

Operation ID: `post_publicBulkCreateSendingSuppressions`

Bulk import account Email Sending suppressions

## Definition

```yaml
{"operationId": "post_publicBulkCreateSendingSuppressions", "summary": "Bulk import account Email Sending suppressions", "description": "Bulk import account Email Sending suppressions", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "array", "items": {"properties": {"email": {"type": "string"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "note": {"type": "string", "maxLength": 1000}}, "required": ["email"], "type": "object"}, "maxItems": 1000}}, "required": ["items"]}}}}, "responses": {"200": {"description": "Bulk import result.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"nullable": true}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object", "properties": {"deduplicated": {"type": "integer", "example": 1}, "errors": {"type": "integer"}, "invalid": {"type": "integer", "example": 1}, "items": {"type": "array", "items": {"properties": {"email": {"type": "string", "format": "email", "example": "user@example.com"}, "error": {"type": "string", "example": "Invalid email"}, "id": {"type": "string", "format": "uuid"}, "index": {"type": "integer"}, "status": {"type": "string", "enum": ["processed", "invalid", "error", "skipped"]}}, "required": ["index", "status"], "type": "object"}}, "processed": {"type": "integer", "example": 2}, "skipped": {"type": "integer"}, "total": {"type": "integer", "example": 4}}, "required": ["processed", "deduplicated", "errors", "invalid", "skipped", "total", "items"]}, "success": {"type": "boolean", "example": true}}, "required": ["errors", "messages", "success", "result"]}}}}, "429": {"description": "Rate limited.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Suppression not found"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["errors", "messages", "result", "success"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
