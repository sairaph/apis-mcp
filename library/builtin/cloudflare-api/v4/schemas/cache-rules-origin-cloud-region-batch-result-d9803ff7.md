---
title: cache-rules_origin_cloud_region_batch_result
page_id: schema-cache-rules-origin-cloud-region-batch-result-d9803ff7
path: schemas
description: Response result for a batch origin cloud region operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_batch_result

Response result for a batch origin cloud region operation.

```yaml
{"description": "Response result for a batch origin cloud region operation.", "type": "object", "properties": {"editable": {"description": "Whether the setting can be modified by the current user.", "type": "boolean"}, "id": {"type": "string", "example": "origin_public_cloud_region", "enum": ["origin_public_cloud_region"], "x-auditable": true}, "modified_on": {"description": "Time the mapping set was last modified. Null when no items were successfully applied.", "type": "string", "format": "date-time", "nullable": true, "x-auditable": true}, "value": {"type": "object", "properties": {"failed": {"description": "Items that could not be applied, with error details.", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_batch_item_result"}, "type": "array"}, "succeeded": {"description": "Items that were successfully applied.", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_batch_item_result"}, "type": "array"}}, "required": ["succeeded", "failed"]}}, "required": ["id", "editable", "value"]}
```
