---
title: r2-slurper_R2TargetSchema
page_id: schema-r2-slurper-r2targetschema-cc72cb29
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_R2TargetSchema

```yaml
{"type": "object", "properties": {"bucket": {"type": "string"}, "jurisdiction": {"$ref": "#/components/schemas/r2-slurper_Jurisdiction"}, "secret": {"$ref": "#/components/schemas/r2-slurper_S3LikeCredsSchema"}, "vendor": {"type": "string", "enum": ["r2"]}}, "required": ["vendor", "bucket", "secret"]}
```
