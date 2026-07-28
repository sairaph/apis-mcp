---
title: vectorize_create-metadata-index-request
page_id: schema-vectorize-create-metadata-index-request-8787ea61
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_create-metadata-index-request

```yaml
{"type": "object", "properties": {"indexType": {"description": "Specifies the type of metadata property to index.", "type": "string", "enum": ["string", "number", "boolean"], "x-auditable": true}, "propertyName": {"description": "Specifies the metadata property to index.", "type": "string", "example": "random_metadata_property", "x-auditable": true}}, "required": ["propertyName", "indexType"]}
```
