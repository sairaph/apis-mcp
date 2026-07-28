---
title: cloudflare-pipelines_Format
page_id: schema-cloudflare-pipelines-format-668e38b2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_Format

```yaml
{"discriminator": {"propertyName": "type"}, "oneOf": [{"allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_JsonFormat"}, {"properties": {"type": {"type": "string", "enum": ["json"]}}, "required": ["type"], "type": "object"}], "title": "Json"}, {"allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_ParquetFormat"}, {"properties": {"type": {"type": "string", "enum": ["parquet"]}}, "required": ["type"], "type": "object"}], "title": "Parquet"}]}
```
