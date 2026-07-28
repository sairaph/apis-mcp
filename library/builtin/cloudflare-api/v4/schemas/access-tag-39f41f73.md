---
title: access_tag
page_id: schema-access-tag-39f41f73
path: schemas
description: A tag
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_tag

A tag

```yaml
{"description": "A tag", "type": "object", "properties": {"app_count": {"description": "The number of applications that have this tag", "type": "integer", "example": 1, "x-auditable": true, "x-stainless-skip": true}, "created_at": {"$ref": "#/components/schemas/access_created_at"}, "name": {"$ref": "#/components/schemas/access_name-13"}, "updated_at": {"$ref": "#/components/schemas/access_updated_at"}}, "required": ["name"]}
```
