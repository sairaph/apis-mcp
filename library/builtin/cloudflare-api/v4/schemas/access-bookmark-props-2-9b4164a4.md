---
title: access_bookmark_props-2
page_id: schema-access-bookmark-props-2-9b4164a4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_bookmark_props-2

```yaml
{"type": "object", "properties": {"app_launcher_visible": {"default": true}, "domain": {"description": "The URL or domain of the bookmark.", "example": "https://mybookmark.com"}, "logo_url": {"$ref": "#/components/schemas/access_logo_url"}, "name": {"$ref": "#/components/schemas/access_name-8"}, "type": {"description": "The application type.", "type": "string", "example": "bookmark"}}, "required": ["type", "domain"], "title": "Bookmark Application"}
```
