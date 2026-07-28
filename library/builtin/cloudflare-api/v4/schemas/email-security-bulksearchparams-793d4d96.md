---
title: email-security_BulkSearchParams
page_id: schema-email-security-bulksearchparams-793d4d96
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_BulkSearchParams

```yaml
{"type": "object", "properties": {"action_log": {"description": "Deprecated, use `GET /investigate/{investigate_id}/action_log` instead. End of life: November 1, 2026.", "type": "boolean", "default": false, "deprecated": true, "x-stainless-deprecation-message": "Use GET /investigate/{investigate_id}/action_log instead."}, "alert_id": {"type": "string", "nullable": true}, "delivery_status": {"$ref": "#/components/schemas/email-security_MessageDeliveryStatus"}, "detections_only": {"type": "boolean", "default": true}, "domain": {"type": "string", "nullable": true}, "end": {"description": "End of search date range.", "type": "string", "format": "date-time", "example": "2022-07-25T14:30:00Z"}, "exact_subject": {"type": "string", "nullable": true}, "final_disposition": {"$ref": "#/components/schemas/email-security_DispositionLabel"}, "message_action": {"type": "string", "enum": ["PREVIEW", "QUARANTINE_RELEASED", "MOVED"], "nullable": true}, "message_id": {"type": "string", "nullable": true}, "metric": {"type": "string", "nullable": true}, "query": {"nullable": true, "type": "string"}, "recipient": {"type": "string", "nullable": true}, "sender": {"type": "string", "nullable": true}, "start": {"description": "Beginning of search date range.", "type": "string", "format": "date-time", "example": "2022-06-25T14:30:00Z"}, "subject": {"type": "string", "nullable": true}, "submissions": {"type": "boolean", "default": false}}}
```
