---
title: zones_vanity_name_servers
page_id: schema-zones-vanity-name-servers-b0ac2d4e
path: schemas
description: |-
    An array of domains used for custom name servers. This is only
    available for Business and Enterprise plans.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_vanity_name_servers

An array of domains used for custom name servers. This is only
available for Business and Enterprise plans.

```yaml
{"description": "An array of domains used for custom name servers. This is only\navailable for Business and Enterprise plans.", "type": "array", "items": {"format": "hostname", "maxLength": 253, "type": "string"}, "example": ["ns1.example.com", "ns2.example.com"], "default": []}
```
