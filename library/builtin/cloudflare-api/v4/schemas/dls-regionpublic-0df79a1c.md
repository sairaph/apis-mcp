---
title: dls_RegionPublic
page_id: schema-dls-regionpublic-0df79a1c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_RegionPublic

```yaml
{"type": "object", "properties": {"created_on": {"type": "string", "format": "date-time"}, "id": {"type": "string"}, "modified_on": {"type": "string", "format": "date-time"}, "name": {"type": "string"}, "region_key": {"type": "string", "maxLength": 128, "minLength": 1, "pattern": "^[a-z0-9_-]+$"}, "version": {"type": "integer"}, "version_created_on": {"type": "string", "format": "date-time"}}, "required": ["id", "name", "region_key", "created_on", "modified_on", "version", "version_created_on"]}
```
