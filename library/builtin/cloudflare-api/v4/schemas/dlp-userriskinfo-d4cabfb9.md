---
title: dlp_UserRiskInfo
page_id: schema-dlp-userriskinfo-d4cabfb9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_UserRiskInfo

```yaml
{"type": "object", "properties": {"email": {"type": "string"}, "event_count": {"type": "integer", "minimum": 0}, "last_event": {"type": "string", "format": "date-time"}, "max_risk_level": {"$ref": "#/components/schemas/dlp_RiskLevel"}, "name": {"type": "string"}, "user_id": {"type": "string", "format": "uuid"}}, "required": ["user_id", "name", "email", "max_risk_level", "event_count", "last_event"]}
```
