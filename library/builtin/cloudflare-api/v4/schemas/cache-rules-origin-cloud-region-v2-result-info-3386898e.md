---
title: cache-rules_origin_cloud_region_v2_result_info
page_id: schema-cache-rules-origin-cloud-region-v2-result-info-3386898e
path: schemas
description: Pagination metadata for list responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_v2_result_info

Pagination metadata for list responses.

```yaml
{"description": "Pagination metadata for list responses.", "type": "object", "properties": {"count": {"description": "Number of items returned in this response.", "type": "integer", "example": 2}, "page": {"description": "Current page number.", "type": "integer", "example": 1}, "per_page": {"description": "Number of items per page.", "type": "integer", "example": 20}, "total_count": {"description": "Total number of mappings configured for the zone.", "type": "integer", "example": 2}, "total_pages": {"description": "Total number of pages.", "type": "integer", "example": 1}}, "required": ["page", "per_page", "count", "total_count", "total_pages"]}
```
