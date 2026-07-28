---
title: zones_advanced_ddos
page_id: schema-zones-advanced-ddos-9f309bb9
path: schemas
description: Advanced protection from Distributed Denial of Service (DDoS) attacks on your website. This is an uneditable value that is 'on' in the case of Business and Enterprise zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_advanced_ddos

Advanced protection from Distributed Denial of Service (DDoS) attacks on your website. This is an uneditable value that is 'on' in the case of Business and Enterprise zones.

```yaml
{"description": "Advanced protection from Distributed Denial of Service (DDoS) attacks on your website. This is an uneditable value that is 'on' in the case of Business and Enterprise zones.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "advanced_ddos", "enum": ["advanced_ddos"]}, "value": {"$ref": "#/components/schemas/zones_advanced_ddos_value"}}}], "title": "Advanced DDoS Protection"}
```
