---
title: resource-tagging_tagged_resource_object_worker
page_id: schema-resource-tagging-tagged-resource-object-worker-8e7e71ca
path: schemas
description: Response for worker resources
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tagged_resource_object_worker

Response for worker resources

```yaml
{"description": "Response for worker resources", "allOf": [{"properties": {"type": {"type": "string", "enum": ["worker"]}}, "required": ["type"], "type": "object"}, {"$ref": "#/components/schemas/resource-tagging_tagged_resource_object_account_level_base"}]}
```
