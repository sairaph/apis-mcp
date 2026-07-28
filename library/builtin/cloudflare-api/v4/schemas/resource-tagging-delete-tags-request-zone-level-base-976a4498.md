---
title: resource-tagging_delete_tags_request_zone_level_base
page_id: schema-resource-tagging-delete-tags-request-zone-level-base-976a4498
path: schemas
description: Request body schema for deleting tags from zone-level resources. Zone ID comes from URL path.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_delete_tags_request_zone_level_base

Request body schema for deleting tags from zone-level resources. Zone ID comes from URL path.

```yaml
{"description": "Request body schema for deleting tags from zone-level resources. Zone ID comes from URL path.", "type": "object", "properties": {"resource_id": {"$ref": "#/components/schemas/resource-tagging_resource_id"}, "resource_type": {"$ref": "#/components/schemas/resource-tagging_zone_resource_type_base_enum"}}, "required": ["resource_type", "resource_id"]}
```
