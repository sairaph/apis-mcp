---
title: Create account Email Sending suppression
page_id: operation-post-accounts-account-id-email-sending-suppressions-edaa0e57
path: operations/email-sending-suppressions
description: Create account Email Sending suppression
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email/sending/suppressions
operation_ids:
    - post_publicCreateSendingSuppression
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create account Email Sending suppression

`POST /accounts/{account_id}/email/sending/suppressions`

Operation ID: `post_publicCreateSendingSuppression`

Create account Email Sending suppression

## Definition

```yaml
{"operationId": "post_publicCreateSendingSuppression", "summary": "Create account Email Sending suppression", "description": "Create account Email Sending suppression", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "note": {"type": "string", "example": "Imported from CRM", "maxLength": 1000}}, "required": ["email"]}}}}, "responses": {"200": {"description": "Suppression mutation completed.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"nullable": true}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object", "properties": {"id": {"type": "string", "format": "uuid"}}, "required": ["id"]}, "success": {"type": "boolean", "example": true}}, "required": ["errors", "messages", "success", "result"]}}}}, "404": {"description": "Suppression not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Suppression not found"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["errors", "messages", "result", "success"]}}}}, "409": {"description": "Suppression already exists.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "Suppression not found"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["errors", "messages", "result", "success"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
