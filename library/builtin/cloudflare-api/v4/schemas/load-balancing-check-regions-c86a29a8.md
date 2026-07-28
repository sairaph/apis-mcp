---
title: load-balancing_check_regions
page_id: schema-load-balancing-check-regions-c86a29a8
path: schemas
description: A list of regions from which to run health checks. Null means every Cloudflare data center.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_check_regions

A list of regions from which to run health checks. Null means every Cloudflare data center.

```yaml
{"description": "A list of regions from which to run health checks. Null means every Cloudflare data center.", "type": "array", "items": {"description": "WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (ENTERPRISE customers only).", "enum": ["WNAM", "ENAM", "WEU", "EEU", "NSAM", "SSAM", "OC", "ME", "NAF", "SAF", "SAS", "SEAS", "NEAS", "ALL_REGIONS"], "type": "string", "x-auditable": true}, "example": ["WEU", "ENAM"], "nullable": true}
```
