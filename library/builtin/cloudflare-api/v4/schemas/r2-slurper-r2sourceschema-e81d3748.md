---
title: r2-slurper_R2SourceSchema
page_id: schema-r2-slurper-r2sourceschema-e81d3748
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_R2SourceSchema

```yaml
{"type": "object", "properties": {"bucket": {"type": "string"}, "jurisdiction": {"$ref": "#/components/schemas/r2-slurper_Jurisdiction"}, "keys": {"type": "array", "items": {"type": "string"}, "maxItems": 10000, "nullable": true}, "pathPrefix": {"type": "string", "nullable": true}, "secret": {"$ref": "#/components/schemas/r2-slurper_S3LikeCredsSchema"}, "vendor": {"type": "string", "enum": ["r2"]}}, "required": ["vendor", "bucket", "secret"]}
```
