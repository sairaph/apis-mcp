---
title: access_onelogin-2
page_id: schema-access-onelogin-2-a7faffd7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_onelogin-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config-2"}, {"properties": {"onelogin_account": {"description": "Your OneLogin account url", "type": "string", "example": "https://mycompany.onelogin.com"}}, "type": "object"}]}}, "type": "object"}], "title": "OneLogin"}
```
