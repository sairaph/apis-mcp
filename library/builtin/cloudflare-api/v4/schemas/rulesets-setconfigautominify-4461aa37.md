---
title: rulesets_SetConfigAutominify
page_id: schema-rulesets-setconfigautominify-4461aa37
path: schemas
description: Which file extensions to minify automatically.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetConfigAutominify

Which file extensions to minify automatically.

```yaml
{"description": "Which file extensions to minify automatically.", "type": "object", "properties": {"css": {"description": "Whether to minify CSS files.", "type": "boolean", "example": true, "default": false, "title": "CSS"}, "html": {"description": "Whether to minify HTML files.", "type": "boolean", "example": true, "default": false, "title": "HTML"}, "js": {"description": "Whether to minify JavaScript files.", "type": "boolean", "example": true, "default": false, "title": "JavaScript"}}, "title": "Auto-Minify"}
```
