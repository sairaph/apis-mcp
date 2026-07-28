---
title: zones_replace_insecure_js
page_id: schema-zones-replace-insecure-js-f5341b5a
path: schemas
description: 'Automatically replace insecure JavaScript libraries with safer and faster alternatives provided under cdnjs and powered by Cloudflare. Currently supports the following libraries: Polyfill under polyfill.io.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_replace_insecure_js

Automatically replace insecure JavaScript libraries with safer and faster alternatives provided under cdnjs and powered by Cloudflare. Currently supports the following libraries: Polyfill under polyfill.io.

```yaml
{"description": "Automatically replace insecure JavaScript libraries with safer and faster alternatives provided under cdnjs and powered by Cloudflare. Currently supports the following libraries: Polyfill under polyfill.io.\n", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "replace_insecure_js", "enum": ["replace_insecure_js"]}, "value": {"$ref": "#/components/schemas/zones_replace_insecure_js_value"}}}], "title": "Replace insecure JavaScript libraries"}
```
