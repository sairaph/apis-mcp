---
title: Get feedback emails statistics
page_id: operation-get-accounts-account-id-email-sending-feedback-5aff7630
path: operations/email-sending-feedback
description: Returns the count of feedback emails received for the specified account and date range.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/email/sending/feedback
operation_ids:
    - get_publicFeedbackStatus
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get feedback emails statistics

`GET /accounts/{account_id}/email/sending/feedback`

Operation ID: `get_publicFeedbackStatus`

Returns the count of feedback emails received for the specified account and date range.

## Definition

```yaml
{"operationId": "get_publicFeedbackStatus", "summary": "Get feedback emails statistics", "description": "Returns the count of feedback emails received for the specified account and date range.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "54442216"}}, {"name": "start_at", "in": "query", "schema": {"type": "string", "format": "date-time", "example": "2026-04-08T00:00:00Z"}}, {"name": "end_at", "in": "query", "schema": {"type": "string", "format": "date-time", "example": "2026-04-15T00:00:00Z"}}], "responses": {"200": {"description": "Returns feedback statistics.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"count": {"type": "number"}, "end_at": {"type": "string", "format": "date-time"}, "start_at": {"type": "string", "format": "date-time"}}, "required": ["count", "start_at", "end_at"]}}, "required": ["result"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Email Sending feedback"]}
```
