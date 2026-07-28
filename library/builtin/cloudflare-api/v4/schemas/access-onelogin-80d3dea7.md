---
title: access_onelogin
page_id: schema-access-onelogin-80d3dea7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_onelogin

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config"}, {"$ref": "#/components/schemas/access_custom-claims-support"}, {"properties": {"onelogin_account": {"description": "Your OneLogin account url", "type": "string", "example": "https://mycompany.onelogin.com", "x-auditable": true}}, "type": "object"}]}}, "type": "object"}], "title": "OneLogin"}
```
