---
title: resource-tagging_set_tags_request_zone_level_base
page_id: schema-resource-tagging-set-tags-request-zone-level-base-938f0a5a
path: schemas
description: Request body schema for setting tags on zone-level resources with no extra requirements.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_set_tags_request_zone_level_base

Request body schema for setting tags on zone-level resources with no extra requirements.

```yaml
{"description": "Request body schema for setting tags on zone-level resources with no extra requirements.", "allOf": [{"$ref": "#/components/schemas/resource-tagging_delete_tags_request_zone_level_base"}, {"properties": {"tags": {"$ref": "#/components/schemas/resource-tagging_tags"}}}]}
```
