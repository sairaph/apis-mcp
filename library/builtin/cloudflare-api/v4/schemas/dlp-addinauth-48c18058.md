---
title: dlp_AddinAuth
page_id: schema-dlp-addinauth-48c18058
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_AddinAuth

```yaml
{"oneOf": [{"properties": {"allowed_microsoft_organizations": {"type": "array", "items": {"type": "string"}}, "type": {"type": "string", "enum": ["Org"]}}, "required": ["allowed_microsoft_organizations", "type"], "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["NoAuth"]}}, "required": ["type"], "type": "object"}]}
```
