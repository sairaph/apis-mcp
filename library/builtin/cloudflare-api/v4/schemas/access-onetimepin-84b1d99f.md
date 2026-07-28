---
title: access_onetimepin
page_id: schema-access-onetimepin-84b1d99f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_onetimepin

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"type": "object", "properties": {"redirect_url": {"type": "string", "readOnly": true}}}, "type": {"type": "string", "enum": ["onetimepin"]}}}], "title": "One Time Pin"}
```
