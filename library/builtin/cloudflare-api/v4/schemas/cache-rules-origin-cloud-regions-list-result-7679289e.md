---
title: cache-rules_origin_cloud_regions_list_result
page_id: schema-cache-rules-origin-cloud-regions-list-result-7679289e
path: schemas
description: Response result for a list of origin cloud region mappings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_regions_list_result

Response result for a list of origin cloud region mappings.

```yaml
{"description": "Response result for a list of origin cloud region mappings.", "type": "object", "properties": {"editable": {"description": "Whether the setting can be modified by the current user.", "type": "boolean"}, "id": {"type": "string", "example": "origin_public_cloud_region", "enum": ["origin_public_cloud_region"], "x-auditable": true}, "modified_on": {"description": "Time the mapping set was last modified. Null when no mappings exist.", "type": "string", "format": "date-time", "nullable": true, "x-auditable": true}, "value": {"type": "array", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_entry"}}}, "required": ["id", "editable", "value"]}
```
