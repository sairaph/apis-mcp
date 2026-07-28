---
title: cache-rules_origin_cloud_region_v2_batch_item_result
page_id: schema-cache-rules-origin-cloud-region-v2-batch-item-result-c80e341b
path: schemas
description: Result for a single item in a batch operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_v2_batch_item_result

Result for a single item in a batch operation.

```yaml
{"description": "Result for a single item in a batch operation.", "type": "object", "properties": {"error": {"description": "Error message explaining why the item failed. Present only on failed items.", "type": "string"}, "origin_ip": {"description": "The origin IP address for this item.", "type": "string", "example": "192.0.2.1"}, "region": {"description": "Cloud vendor region identifier. Present on succeeded items (the new value for upsert, the deleted value for delete).", "type": "string", "example": "us-east-1"}, "vendor": {"description": "Cloud vendor identifier. Present on succeeded items (the new value for upsert, the deleted value for delete).", "type": "string", "example": "aws"}}, "required": ["origin_ip"]}
```
