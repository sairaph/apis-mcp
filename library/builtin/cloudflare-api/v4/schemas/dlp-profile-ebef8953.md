---
title: dlp_Profile
page_id: schema-dlp-profile-ebef8953
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_Profile

```yaml
{"oneOf": [{"allOf": [{"$ref": "#/components/schemas/dlp_CustomProfile"}, {"properties": {"type": {"type": "string", "enum": ["custom"]}}, "required": ["type"], "type": "object"}], "title": "Custom Profile"}, {"allOf": [{"$ref": "#/components/schemas/dlp_PredefinedProfile"}, {"properties": {"type": {"type": "string", "enum": ["predefined"]}}, "required": ["type"], "type": "object"}], "title": "Predefined Profile"}, {"allOf": [{"$ref": "#/components/schemas/dlp_IntegrationProfile"}, {"properties": {"type": {"type": "string", "enum": ["integration"]}}, "required": ["type"], "type": "object"}], "title": "Integration Profile"}]}
```
