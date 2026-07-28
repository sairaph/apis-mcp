---
title: security-center_auditLog
page_id: schema-security-center-auditlog-3dd15932
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_auditLog

```yaml
{"type": "object", "properties": {"changed_at": {"description": "The timestamp when the change occurred.", "type": "string", "format": "date-time", "x-auditable": true}, "changed_by": {"description": "The actor that made the change. 'system' for automated changes, or a user identifier.", "type": "string", "example": "system", "x-auditable": true}, "current_value": {"description": "The value of the field after the change. Null if the field was cleared.", "type": "string", "nullable": true, "x-auditable": true}, "field_changed": {"description": "The field that was changed.", "type": "string", "enum": ["status", "user_classification"], "x-auditable": true}, "id": {"description": "UUIDv7 identifier for the audit log entry, time-ordered.", "type": "string", "format": "uuid", "example": "019b0000-0000-7000-8000-000000000001", "x-auditable": true}, "issue_id": {"description": "The ID of the insight this audit log entry relates to.", "type": "string", "x-auditable": true}, "previous_value": {"description": "The value of the field before the change. Null if the field was not previously set.", "type": "string", "nullable": true, "x-auditable": true}, "rationale": {"description": "Optional rationale provided for the change.", "type": "string", "nullable": true, "x-auditable": true}, "zone_id": {"description": "The zone ID associated with the insight. Only present for zone-level insights.", "type": "integer", "format": "int64", "x-auditable": true}}}
```
