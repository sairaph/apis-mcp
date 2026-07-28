---
title: zones_max_upload
page_id: schema-zones-max-upload-5371f54d
path: schemas
description: Maximum size of an allowable upload.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_max_upload

Maximum size of an allowable upload.

```yaml
{"description": "Maximum size of an allowable upload.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "identifier of the zone setting.", "example": "max_upload", "enum": ["max_upload"]}, "value": {"$ref": "#/components/schemas/zones_max_upload_value"}}}], "title": "Max Upload"}
```
