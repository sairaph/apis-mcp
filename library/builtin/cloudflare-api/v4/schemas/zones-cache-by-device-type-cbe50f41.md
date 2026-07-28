---
title: zones_cache_by_device_type
page_id: schema-zones-cache-by-device-type-cbe50f41
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache_by_device_type

```yaml
{"type": "object", "properties": {"id": {"description": "Separate cached content based on the visitor's device type.\n", "type": "string", "enum": ["cache_by_device_type"], "x-auditable": true}, "value": {"description": "The status of Cache By Device Type.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Cache By Device Type", "x-stainless-skip": ["terraform"]}
```
