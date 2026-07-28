---
title: Fetch an email message
page_id: operation-get-accounts-account-id-email-sending-messages-message-id-724f928d
path: operations/email-sending
description: Returns the raw RFC 5322 MIME message for the given account and message id.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/sending/messages/{message_id}
operation_ids:
    - email-sending-get-email-message
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch an email message

`GET /accounts/{account_id}/email/sending/messages/{message_id}`

Operation ID: `email-sending-get-email-message`

Returns the raw RFC 5322 MIME message for the given account and message id.

## Definition

```yaml
{"operationId": "email-sending-get-email-message", "summary": "Fetch an email message", "description": "Returns the raw RFC 5322 MIME message for the given account and message id.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345", "pattern": "^\\d+$"}}, {"name": "message_id", "in": "path", "required": true, "schema": {"type": "string", "example": "<aB3xK9mP@example.com>"}}], "responses": {"200": {"description": "The raw RFC 5322 MIME message.", "content": {"message/rfc822": {"schema": {"type": "string", "format": "binary"}}}}, "404": {"description": "Account or message not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-sending_EmailApiError"}}}}, "500": {"description": "An unexpected error occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/email-sending_EmailApiError"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Email Sending"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.email.sending.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
