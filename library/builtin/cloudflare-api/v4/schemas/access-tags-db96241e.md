---
title: access_tags
page_id: schema-access-tags-db96241e
path: schemas
description: The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_tags

The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.

```yaml
{"description": "The tags you want assigned to an application. Tags are used to filter applications in the App Launcher dashboard.", "type": "array", "items": {"description": "The tag associated with an application.", "example": "engineers", "type": "string"}, "default": [], "x-stainless-collection-type": "set"}
```
