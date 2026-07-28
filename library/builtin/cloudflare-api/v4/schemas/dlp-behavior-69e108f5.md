---
title: dlp_Behavior
page_id: schema-dlp-behavior-69e108f5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_Behavior

```yaml
{"type": "object", "properties": {"description": {"type": "string"}, "enabled": {"type": "boolean"}, "name": {"type": "string"}, "risk_level": {"$ref": "#/components/schemas/dlp_RiskLevel"}}, "required": ["name", "description", "risk_level", "enabled"]}
```
