---
title: cache-rules_origin_cloud_region_v2_delete_result
page_id: schema-cache-rules-origin-cloud-region-v2-delete-result-929dd18e
path: schemas
description: Response result for a delete operation. Identifies the deleted mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_v2_delete_result

Response result for a delete operation. Identifies the deleted mapping.

```yaml
{"description": "Response result for a delete operation. Identifies the deleted mapping.", "type": "object", "properties": {"origin_ip": {"description": "The origin IP address whose mapping was deleted.", "type": "string", "example": "192.0.2.1", "x-auditable": true}}, "required": ["origin_ip"]}
```
