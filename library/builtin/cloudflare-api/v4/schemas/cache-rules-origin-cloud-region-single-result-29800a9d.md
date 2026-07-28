---
title: cache-rules_origin_cloud_region_single_result
page_id: schema-cache-rules-origin-cloud-region-single-result-29800a9d
path: schemas
description: Response result for a single origin cloud region mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_single_result

Response result for a single origin cloud region mapping.

```yaml
{"description": "Response result for a single origin cloud region mapping.", "type": "object", "properties": {"editable": {"description": "Whether the setting can be modified by the current user.", "type": "boolean"}, "id": {"type": "string", "example": "origin_public_cloud_region", "enum": ["origin_public_cloud_region"], "x-auditable": true}, "modified_on": {"description": "Time the mapping was last modified.", "type": "string", "format": "date-time", "x-auditable": true}, "value": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_entry"}}, "required": ["id", "editable", "value"]}
```
