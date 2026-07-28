---
title: List account Email Sending suppressions
page_id: operation-get-accounts-account-id-email-sending-suppressions-f859ef19
path: operations/email-sending-suppressions
description: List account Email Sending suppressions
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/sending/suppressions
operation_ids:
    - get_publicListSendingSuppressions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account Email Sending suppressions

`GET /accounts/{account_id}/email/sending/suppressions`

Operation ID: `get_publicListSendingSuppressions`

List account Email Sending suppressions

## Definition

```yaml
{"operationId": "get_publicListSendingSuppressions", "summary": "List account Email Sending suppressions", "description": "List account Email Sending suppressions", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 100, "maximum": 1000, "minimum": 1}}, {"name": "cursor", "in": "query", "schema": {"description": "Opaque pagination cursor returned as `result_info.next_cursor`. When supplied, `order`, `direction`, `email` and `reason` are taken from the cursor and any conflicting values are rejected.", "type": "string", "example": "eyJ0IjoxLCJvIjoiY3JlYXRlZF9hdCIsImQiOiJkZXNjIn0"}}, {"name": "email", "in": "query", "schema": {"type": "string", "format": "email", "example": "user@example.com"}}, {"name": "reason", "in": "query", "schema": {"type": "string", "enum": ["manual", "complaint", "hard_bounce", "soft_bounce"]}}, {"name": "order", "in": "query", "schema": {"type": "string", "default": "created_at", "enum": ["email", "expires_at", "created_at"]}}, {"name": "direction", "in": "query", "schema": {"type": "string", "default": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "Returns account-level Email Sending suppressions.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"nullable": true}}, "messages": {"type": "array", "items": {"nullable": true}}, "result": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "email": {"type": "string", "format": "email", "example": "user@example.com"}, "expires_at": {"type": "string", "format": "date-time", "nullable": true}, "id": {"type": "string", "format": "uuid"}, "reason": {"type": "string", "example": "hard_bounce"}}, "required": ["id", "email", "reason", "created_at", "expires_at"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"type": "integer", "example": 1}, "next_cursor": {"description": "Opaque cursor for the next page. Pass it back as the `cursor` query parameter. `null` when there are no more results.", "type": "string", "example": "eyJ0IjoxLCJvIjoiY3JlYXRlZF9hdCIsImQiOiJkZXNjIn0", "nullable": true}, "per_page": {"type": "integer", "example": 100}}, "required": ["count", "per_page", "next_cursor"]}, "success": {"type": "boolean", "example": true}}, "required": ["errors", "messages", "success", "result", "result_info"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending suppressions"]}
```
