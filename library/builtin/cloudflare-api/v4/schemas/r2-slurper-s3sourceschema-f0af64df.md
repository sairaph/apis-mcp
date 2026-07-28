---
title: r2-slurper_S3SourceSchema
page_id: schema-r2-slurper-s3sourceschema-f0af64df
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_S3SourceSchema

```yaml
{"type": "object", "properties": {"bucket": {"type": "string"}, "endpoint": {"description": "Custom S3-compatible endpoint that must use https://.", "type": "string", "format": "uri", "nullable": true}, "keys": {"type": "array", "items": {"type": "string"}, "maxItems": 10000, "nullable": true}, "pathPrefix": {"type": "string", "nullable": true}, "region": {"type": "string", "nullable": true}, "secret": {"$ref": "#/components/schemas/r2-slurper_S3LikeCredsSchema"}, "vendor": {"type": "string", "enum": ["s3"]}}, "required": ["vendor", "bucket", "secret"]}
```
