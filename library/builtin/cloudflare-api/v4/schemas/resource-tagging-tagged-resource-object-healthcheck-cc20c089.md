---
title: resource-tagging_tagged_resource_object_healthcheck
page_id: schema-resource-tagging-tagged-resource-object-healthcheck-cc20c089
path: schemas
description: Response for healthcheck resources
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tagged_resource_object_healthcheck

Response for healthcheck resources

```yaml
{"description": "Response for healthcheck resources", "allOf": [{"properties": {"type": {"type": "string", "enum": ["healthcheck"]}}, "required": ["type"], "type": "object"}, {"$ref": "#/components/schemas/resource-tagging_tagged_resource_object_zone_level_base"}]}
```
