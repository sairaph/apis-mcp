---
title: zones_sha1_support
page_id: schema-zones-sha1-support-8f4dea65
path: schemas
description: Allow SHA1 support.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_sha1_support

Allow SHA1 support.

```yaml
{"description": "Allow SHA1 support.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "Zone setting identifier.", "example": "sha1_support", "enum": ["sha1_support"]}, "value": {"$ref": "#/components/schemas/zones_sha1_support_value"}}}], "title": "Toggle SHA1 support"}
```
