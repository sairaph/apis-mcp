---
title: resource-tagging_tagged_resource_object_load_balancer
page_id: schema-resource-tagging-tagged-resource-object-load-balancer-aa210b72
path: schemas
description: Response for load_balancer resources
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tagged_resource_object_load_balancer

Response for load_balancer resources

```yaml
{"description": "Response for load_balancer resources", "allOf": [{"properties": {"type": {"type": "string", "enum": ["load_balancer"]}}, "required": ["type"], "type": "object"}, {"$ref": "#/components/schemas/resource-tagging_tagged_resource_object_zone_level_base"}]}
```
