---
title: email-security_ActionEntry
page_id: schema-email-security-actionentry-56d63e51
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_ActionEntry

```yaml
{"type": "object", "properties": {"completed_at": {"description": "Timestamp when the action completed.", "type": "string", "format": "date-time", "nullable": true}, "operation": {"type": "string", "enum": ["PREVIEW", "QUARANTINE_RELEASE", "SUBMISSION", "MOVE"]}, "properties": {"type": "object", "allOf": [{"$ref": "#/components/schemas/email-security_ActionEntryProperties"}], "nullable": true}, "started_at": {"description": "Timestamp when the action was initiated.", "type": "string", "format": "date-time", "nullable": true}, "status": {"type": "string", "nullable": true}, "success": {"type": "boolean", "nullable": true}}, "required": ["operation"]}
```
