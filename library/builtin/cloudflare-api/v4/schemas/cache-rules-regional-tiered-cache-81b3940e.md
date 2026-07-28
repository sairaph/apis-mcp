---
title: cache-rules_regional_tiered_cache
page_id: schema-cache-rules-regional-tiered-cache-81b3940e
path: schemas
description: Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_regional_tiered_cache

Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.

```yaml
{"description": "Instructs Cloudflare to check a regional hub data center on the way to your upper tier. This can help improve performance for smart and custom tiered cache topologies.", "type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "type": "string", "example": "tc_regional", "enum": ["tc_regional"], "x-auditable": true}}, "type": "object"}], "title": "Regional Tiered Cache"}
```
