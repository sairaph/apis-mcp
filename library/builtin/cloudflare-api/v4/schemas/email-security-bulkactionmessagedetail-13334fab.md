---
title: email-security_BulkActionMessageDetail
page_id: schema-email-security-bulkactionmessagedetail-13334fab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BulkActionMessageDetail

```yaml
{"type": "object", "properties": {"action_params": {"$ref": "#/components/schemas/email-security_BulkMessageActionParams"}, "action_type": {"type": "string", "enum": ["MOVE", "RELEASE"]}, "alert_id": {"type": "string", "nullable": true}, "created_at": {"type": "string", "format": "date-time"}, "email_message_id": {"type": "string", "nullable": true}, "message_id": {"type": "string", "format": "uuid"}, "postfix_id": {"type": "string"}, "processed_at": {"type": "string", "format": "date-time", "nullable": true}, "retry_after": {"description": "When to retry the action if it failed.", "type": "string", "format": "date-time", "nullable": true}, "retry_count": {"type": "integer"}, "status": {"type": "string", "enum": ["PENDING", "DISCOVERING", "PROCESSING", "COMPLETED", "FAILED", "CANCELLED", "SKIPPED"]}, "status_message": {"type": "string", "nullable": true}}, "required": ["message_id", "postfix_id", "action_type", "action_params", "status", "retry_count", "created_at"]}
```
