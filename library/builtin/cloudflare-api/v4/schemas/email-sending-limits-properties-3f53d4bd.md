---
title: email_sending_limits_properties
page_id: schema-email-sending-limits-properties-3f53d4bd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_sending_limits_properties

```yaml
{"type": "object", "properties": {"quota": {"description": "The resolved daily sending quota for the account. Null when the quota is not yet available.", "type": "object", "nullable": true, "properties": {"unit": {"description": "The time period for the quota.", "type": "string", "example": "day", "enum": ["day", "hour"]}, "value": {"description": "The quota limit.", "type": "integer", "example": 1000}}, "readOnly": true}, "usage": {"description": "The account's current daily sending usage. Null when there is no resolved quota or usage is temporarily unavailable.", "type": "object", "nullable": true, "properties": {"over_quota": {"description": "Whether the account has exceeded its daily sending quota.", "type": "boolean", "example": false}, "resets_at": {"description": "When the current daily quota window resets. Null when there is no active window.", "type": "string", "format": "date-time", "example": "2026-07-08T00:00:00Z", "nullable": true, "readOnly": true}, "sent": {"description": "Emails sent against the daily quota in the current window.", "type": "integer", "example": 42}}, "readOnly": true}}}
```
