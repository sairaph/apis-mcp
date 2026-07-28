---
title: cloudflare-pipelines_FieldType
page_id: schema-cloudflare-pipelines-fieldtype-cc1726af
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_FieldType

```yaml
{"discriminator": {"propertyName": "type"}, "oneOf": [{"properties": {"type": {"type": "string", "enum": ["int32"]}}, "required": ["type"], "title": "Int32", "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["int64"]}}, "required": ["type"], "title": "Int64", "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["float32"]}}, "required": ["type"], "title": "Float32", "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["float64"]}}, "required": ["type"], "title": "Float64", "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["bool"]}}, "required": ["type"], "title": "Bool", "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["string"]}}, "required": ["type"], "title": "String", "type": "object"}, {"properties": {"type": {"type": "string", "enum": ["binary"]}}, "required": ["type"], "title": "Binary", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_TimestampField"}, {"properties": {"type": {"type": "string", "enum": ["timestamp"]}}, "required": ["type"], "type": "object"}], "title": "Timestamp"}, {"properties": {"type": {"type": "string", "enum": ["json"]}}, "required": ["type"], "title": "Json", "type": "object"}, {"allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_StructField"}, {"properties": {"type": {"type": "string", "enum": ["struct"]}}, "required": ["type"], "type": "object"}], "title": "Struct"}, {"allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_ListField"}, {"properties": {"type": {"type": "string", "enum": ["list"]}}, "required": ["type"], "type": "object"}], "title": "List"}]}
```
