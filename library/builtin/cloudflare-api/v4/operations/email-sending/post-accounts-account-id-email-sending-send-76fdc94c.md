---
title: Send an email
page_id: operation-post-accounts-account-id-email-sending-send-d85a4dfe
path: operations/email-sending
description: Send an email for the specified account using the structured builder. Provide the sender, recipients, subject, and at least one of text or html; attachments are optional.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/email/sending/send
operation_ids:
    - email-sending-account-send-builder
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Send an email

`POST /accounts/{account_id}/email/sending/send`

Operation ID: `email-sending-account-send-builder`

Send an email for the specified account using the structured builder. Provide the sender, recipients, subject, and at least one of text or html; attachments are optional.

## Definition

```yaml
{"operationId": "email-sending-account-send-builder", "summary": "Send an email", "description": "Send an email for the specified account using the structured builder. Provide the sender, recipients, subject, and at least one of text or html; attachments are optional.", "parameters": [{"name": "account_id", "in": "path", "description": "Identifier of the account.", "required": true, "schema": {"description": "Identifier of the account.", "type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-sending_EmailBuilder"}}}}, "responses": {"200": {"description": "Email sending results.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/email-sending_EmailSendResponse"}, "result_info": {"type": "object", "properties": {"count": {"type": "number"}, "cursor": {"type": "string"}, "page": {"type": "number"}, "per_page": {"type": "number"}, "total_count": {"type": "number"}}, "required": ["per_page", "count", "total_count"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "result", "messages"]}}}}, "400": {"description": "Invalid request, generally because of the format/content of the email send request. No email will be sent when this happens.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "string", "enum": ["10001", "10200", "10201", "10202"]}, "message": {"type": "string", "enum": ["email.sending.error.invalid_request_schema", "email.sending.error.email.invalid", "email.sending.error.email.no_content_length", "email.sending.error.email.too_big"]}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "403": {"description": "Email sending is disabled, or the sending domain is not configured for email sending.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "string", "enum": ["10203", "10204"]}, "message": {"type": "string", "enum": ["email.sending.error.email.sending_disabled", "email.sending.error.email.sender_not_configured"]}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "429": {"description": "Account/zone has surpassed the rate at which it can send email, please try again later. No email will be sent when this happens.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "enum": [10004]}, "message": {"type": "string", "enum": ["email.sending.error.throttled"]}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}, "500": {"description": "An unexpected error while processing the email send request. No email will be sent when this happens.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "enum": [10002]}, "message": {"type": "string", "enum": ["email.sending.error.internal_server"]}}, "required": ["code", "message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "messages", "errors", "result"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Email Sending"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
