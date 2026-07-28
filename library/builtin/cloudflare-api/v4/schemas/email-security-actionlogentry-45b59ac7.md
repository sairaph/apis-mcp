---
title: email-security_ActionLogEntry
page_id: schema-email-security-actionlogentry-45b59ac7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_ActionLogEntry

```yaml
{"type": "object", "properties": {"completed_at": {"description": "Timestamp when action completed.", "type": "string", "format": "date-time", "readOnly": true}, "completed_timestamp": {"description": "Deprecated, use `completed_at` instead. End of life: November 1, 2026.", "type": "string", "deprecated": true, "readOnly": true, "x-stainless-deprecation-message": "Use `completed_at` instead."}, "operation": {"description": "Type of action performed.", "type": "string", "enum": ["MOVE", "RELEASE", "RECLASSIFY", "SUBMISSION", "QUARANTINE_RELEASE", "PREVIEW"]}, "properties": {"description": "Additional properties for the action.", "type": "object", "properties": {"folder": {"description": "Target folder for move operations.", "type": "string"}, "requested_by": {"description": "User who requested the action.", "type": "string"}}}, "status": {"description": "Status of the action.", "type": "string", "nullable": true}}, "required": ["completed_at", "operation"]}
```
