---
title: dlp_AddinAccountMapping
page_id: schema-dlp-addinaccountmapping-2885ce1e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_AddinAccountMapping

```yaml
{"type": "object", "properties": {"addin_identifier_token": {"type": "string", "format": "uuid"}, "auth_requirements": {"$ref": "#/components/schemas/dlp_AddinAuth"}}, "required": ["auth_requirements", "addin_identifier_token"]}
```
