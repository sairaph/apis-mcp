---
title: digital-experience-monitoring_schemas-ip_info
page_id: schema-digital-experience-monitoring-schemas-ip-info-69fac050
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_schemas-ip_info

```yaml
{"type": "object", "properties": {"address": {"type": "string", "nullable": true}, "asn": {"type": "integer", "nullable": true}, "aso": {"type": "string", "nullable": true}, "location": {"type": "object", "properties": {"city": {"type": "string", "nullable": true}, "country_iso": {"type": "string", "nullable": true}, "state_iso": {"type": "string", "nullable": true}, "zip": {"type": "string", "nullable": true}}}, "name": {"type": "string", "nullable": true}, "netmask": {"type": "string", "nullable": true}, "version": {"description": "IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).", "type": "integer", "example": 1}}}
```
