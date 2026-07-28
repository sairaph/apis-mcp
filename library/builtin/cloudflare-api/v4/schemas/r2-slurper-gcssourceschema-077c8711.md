---
title: r2-slurper_GCSSourceSchema
page_id: schema-r2-slurper-gcssourceschema-077c8711
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_GCSSourceSchema

```yaml
{"type": "object", "properties": {"bucket": {"type": "string"}, "keys": {"type": "array", "items": {"type": "string"}, "maxItems": 10000, "nullable": true}, "pathPrefix": {"type": "string", "nullable": true}, "secret": {"$ref": "#/components/schemas/r2-slurper_GCSLikeCredsSchema"}, "vendor": {"type": "string", "enum": ["gcs"]}}, "required": ["vendor", "bucket", "secret"]}
```
