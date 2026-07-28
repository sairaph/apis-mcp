---
title: cache-rules_supported_cloud_regions_result
page_id: schema-cache-rules-supported-cloud-regions-result-74c7344f
path: schemas
description: Cloud vendors and their supported regions for origin cloud region mappings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_supported_cloud_regions_result

Cloud vendors and their supported regions for origin cloud region mappings.

```yaml
{"description": "Cloud vendors and their supported regions for origin cloud region mappings.", "type": "object", "properties": {"obtained_codes": {"description": "Whether Cloudflare airport codes (IATA colo identifiers) were successfully resolved for the `upper_tier_colos` field on each region. When `false`, the `upper_tier_colos` arrays may be empty or incomplete.", "type": "boolean"}, "vendors": {"description": "Map of vendor name to list of supported regions.", "type": "object", "additionalProperties": {"items": {"$ref": "#/components/schemas/cache-rules_supported_cloud_region"}, "type": "array"}}}, "required": ["vendors", "obtained_codes"]}
```
