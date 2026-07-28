---
title: dlp_RiskEvent
page_id: schema-dlp-riskevent-6684d27b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_RiskEvent

```yaml
{"type": "object", "properties": {"event_details": {}, "id": {"type": "string"}, "name": {"type": "string"}, "risk_level": {"$ref": "#/components/schemas/dlp_RiskLevel"}, "timestamp": {"type": "string", "format": "date-time"}}, "required": ["id", "risk_level", "name", "timestamp"]}
```
