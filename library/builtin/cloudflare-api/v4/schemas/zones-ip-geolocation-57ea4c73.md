---
title: zones_ip_geolocation
page_id: schema-zones-ip-geolocation-57ea4c73
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_ip_geolocation

```yaml
{"type": "object", "properties": {"id": {"description": "Cloudflare adds a CF-IPCountry HTTP header containing the country code that corresponds to the visitor.\n", "type": "string", "example": "ip_geolocation", "enum": ["ip_geolocation"], "x-auditable": true}, "value": {"description": "The status of adding the IP Geolocation Header.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "IP Geolocation Header", "x-stainless-skip": ["terraform"]}
```
