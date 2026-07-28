---
title: access_onetimepin-2
page_id: schema-access-onetimepin-2-aa191c39
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_onetimepin-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"type": "object", "properties": {"redirect_url": {"type": "string", "readOnly": true}}}, "type": {"enum": ["onetimepin"]}}}], "title": "One Time Pin"}
```
