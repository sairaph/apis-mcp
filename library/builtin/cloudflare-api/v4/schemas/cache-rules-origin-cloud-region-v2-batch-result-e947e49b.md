---
title: cache-rules_origin_cloud_region_v2_batch_result
page_id: schema-cache-rules-origin-cloud-region-v2-batch-result-e947e49b
path: schemas
description: Response result for a batch origin cloud region operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_v2_batch_result

Response result for a batch origin cloud region operation.

```yaml
{"description": "Response result for a batch origin cloud region operation.", "type": "object", "properties": {"failed": {"description": "Items that could not be applied, with error details.", "type": "array", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_batch_item_result"}}, "succeeded": {"description": "Items that were successfully applied.", "type": "array", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_batch_item_result"}}}, "required": ["succeeded", "failed"]}
```
