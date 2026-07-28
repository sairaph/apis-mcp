---
title: healthchecks_check_regions
page_id: schema-healthchecks-check-regions-edef1ce5
path: schemas
description: A list of regions from which to run health checks. Null means Cloudflare will pick a default region.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# healthchecks_check_regions

A list of regions from which to run health checks. Null means Cloudflare will pick a default region.

```yaml
{"description": "A list of regions from which to run health checks. Null means Cloudflare will pick a default region.", "type": "array", "items": {"description": "WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS and ENTERPRISE customers only).", "enum": ["WNAM", "ENAM", "WEU", "EEU", "NSAM", "SSAM", "OC", "ME", "NAF", "SAF", "IN", "SEAS", "NEAS", "ALL_REGIONS"], "type": "string"}, "example": ["WEU", "ENAM"], "nullable": true, "x-auditable": true}
```
