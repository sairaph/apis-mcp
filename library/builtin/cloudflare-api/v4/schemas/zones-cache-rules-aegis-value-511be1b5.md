---
title: zones_cache-rules_aegis_value
page_id: schema-zones-cache-rules-aegis-value-511be1b5
path: schemas
description: Value of the zone setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache-rules_aegis_value

Value of the zone setting.

```yaml
{"description": "Value of the zone setting.", "type": "object", "properties": {"enabled": {"description": "Whether the feature is enabled or not.", "type": "boolean", "x-auditable": true}, "pool_id": {"description": "Egress pool id which refers to a grouping of dedicated egress IPs through which Cloudflare will connect to origin.", "type": "string", "example": "pool-id", "x-auditable": true}}}
```
