---
title: zones_opportunistic_onion
page_id: schema-zones-opportunistic-onion-5f0a8553
path: schemas
description: Add an Alt-Svc header to all legitimate requests from Tor, allowing the connection to use our onion services instead of exit nodes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_opportunistic_onion

Add an Alt-Svc header to all legitimate requests from Tor, allowing the connection to use our onion services instead of exit nodes.

```yaml
{"description": "Add an Alt-Svc header to all legitimate requests from Tor, allowing the connection to use our onion services instead of exit nodes.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "opportunistic_onion", "enum": ["opportunistic_onion"]}, "value": {"$ref": "#/components/schemas/zones_opportunistic_onion_value"}}}], "title": "Zone Enable Onion Routing"}
```
