---
title: vectorize_create-index-request
page_id: schema-vectorize-create-index-request-ee71044f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_create-index-request

```yaml
{"type": "object", "properties": {"config": {"allOf": [{"$ref": "#/components/schemas/vectorize_index-configuration"}]}, "description": {"$ref": "#/components/schemas/vectorize_index-description"}, "name": {"$ref": "#/components/schemas/vectorize_index-name"}}, "required": ["name", "config"]}
```
