---
title: zones_cache-rules_base
page_id: schema-zones-cache-rules-base-e61f8686
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache-rules_base

```yaml
{"type": "object", "properties": {"id": {"description": "Identifier of the zone setting.", "type": "string", "x-auditable": true}, "modified_on": {"description": "Last time this setting was modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "nullable": true, "readOnly": true, "x-auditable": true}}, "required": ["id"]}
```
