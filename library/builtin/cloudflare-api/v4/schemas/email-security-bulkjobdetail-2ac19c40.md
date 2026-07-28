---
title: email-security_BulkJobDetail
page_id: schema-email-security-bulkjobdetail-2ac19c40
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BulkJobDetail

```yaml
{"type": "object", "properties": {"action_params": {"$ref": "#/components/schemas/email-security_BulkJobActionParams"}, "action_type": {"type": "string", "enum": ["MOVE", "RELEASE"]}, "comment": {"type": "string", "nullable": true}, "completed_at": {"type": "string", "format": "date-time", "nullable": true}, "created_at": {"type": "string", "format": "date-time"}, "job_id": {"type": "string", "format": "uuid"}, "messages_failed": {"type": "integer"}, "messages_pending": {"type": "integer"}, "messages_successful": {"type": "integer"}, "search_params": {"$ref": "#/components/schemas/email-security_BulkSearchParams"}, "started_at": {"type": "string", "format": "date-time", "nullable": true}, "status": {"type": "string", "enum": ["PENDING", "DISCOVERING", "PROCESSING", "COMPLETED", "FAILED", "CANCELLED", "SKIPPED"]}, "status_message": {"type": "string", "nullable": true}, "total_messages_discovered": {"type": "integer"}}, "required": ["job_id", "action_type", "action_params", "search_params", "status", "total_messages_discovered", "messages_pending", "messages_successful", "messages_failed", "created_at"]}
```
