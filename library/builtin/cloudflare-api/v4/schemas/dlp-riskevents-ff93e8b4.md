---
title: dlp_RiskEvents
page_id: schema-dlp-riskevents-ff93e8b4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_RiskEvents

```yaml
{"type": "object", "properties": {"email": {"type": "string"}, "events": {"type": "array", "items": {"$ref": "#/components/schemas/dlp_RiskEvent"}}, "last_reset_time": {"type": "string", "format": "date-time", "nullable": true}, "name": {"type": "string"}, "risk_level": {"allOf": [{"$ref": "#/components/schemas/dlp_RiskLevel"}]}}, "required": ["name", "email", "events"]}
```
