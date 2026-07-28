---
title: access_bookmark_props
page_id: schema-access-bookmark-props-faa4ddec
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_bookmark_props

```yaml
{"type": "object", "properties": {"app_launcher_visible": {"$ref": "#/components/schemas/access_app_launcher_visible"}, "domain": {"description": "The URL or domain of the bookmark.", "type": "string", "example": "https://mybookmark.com"}, "logo_url": {"$ref": "#/components/schemas/access_logo_url"}, "name": {"$ref": "#/components/schemas/access_name-8"}, "tags": {"$ref": "#/components/schemas/access_tags"}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "bookmark"}]}}, "title": "Bookmark application"}
```
