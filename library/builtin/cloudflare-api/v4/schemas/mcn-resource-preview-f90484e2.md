---
title: mcn_resource_preview
page_id: schema-mcn-resource-preview-f90484e2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_resource_preview

```yaml
{"type": "object", "properties": {"cloud_type": {"$ref": "#/components/schemas/mcn_cloud_type"}, "detail": {"type": "string"}, "id": {"$ref": "#/components/schemas/mcn_resource_id"}, "name": {"type": "string"}, "resource_type": {"$ref": "#/components/schemas/mcn_resource_type"}, "title": {"type": "string"}}, "required": ["id", "cloud_type", "resource_type", "name", "title", "detail"]}
```
