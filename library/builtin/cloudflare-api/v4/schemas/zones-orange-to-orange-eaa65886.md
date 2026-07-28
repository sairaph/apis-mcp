---
title: zones_orange_to_orange
page_id: schema-zones-orange-to-orange-eaa65886
path: schemas
description: Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also on Cloudflare.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_orange_to_orange

Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also on Cloudflare.

```yaml
{"description": "Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also on Cloudflare.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "orange_to_orange", "enum": ["orange_to_orange"]}, "value": {"$ref": "#/components/schemas/zones_orange_to_orange_value"}}}], "title": "Orange to Orange"}
```
