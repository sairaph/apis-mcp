---
title: access_custom_page
page_id: schema-access-custom-page-daee3d27
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_custom_page

```yaml
{"type": "object", "properties": {"app_count": {"$ref": "#/components/schemas/access_app_count-2"}, "created_at": {"$ref": "#/components/schemas/access_created_at"}, "custom_html": {"description": "Custom page HTML.", "type": "string", "example": "<html><body><h1>Access Denied</h1></body></html>", "x-auditable": true}, "name": {"$ref": "#/components/schemas/access_name-11"}, "type": {"$ref": "#/components/schemas/access_type-2"}, "uid": {"$ref": "#/components/schemas/access_uuid"}, "updated_at": {"$ref": "#/components/schemas/access_updated_at"}}, "required": ["name", "custom_html", "type"]}
```
