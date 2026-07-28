---
title: cache-rules_supported_cloud_region
page_id: schema-cache-rules-supported-cloud-region-d684e7bd
path: schemas
description: A single supported cloud region with associated Tiered Cache upper-tier colocations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_supported_cloud_region

A single supported cloud region with associated Tiered Cache upper-tier colocations.

```yaml
{"description": "A single supported cloud region with associated Tiered Cache upper-tier colocations.", "type": "object", "properties": {"name": {"description": "Cloud vendor region identifier.", "type": "string", "example": "us-east-1"}, "upper_tier_colos": {"description": "Cloudflare Tiered Cache upper-tier colocation codes co-located with this cloud region. Requests from zones with a matching origin mapping will be routed through these colos.", "type": "array", "items": {"type": "string"}, "example": ["IAD", "EWR"]}}, "required": ["name", "upper_tier_colos"]}
```
