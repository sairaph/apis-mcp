---
title: lists_item_base
page_id: schema-lists-item-base-e3deaa91
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_item_base

```yaml
{"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/lists_item_comment"}, "created_on": {"$ref": "#/components/schemas/lists_created_on"}, "id": {"$ref": "#/components/schemas/lists_item_id"}, "modified_on": {"$ref": "#/components/schemas/lists_modified_on"}}, "example": {"comment": "Private IP address", "created_on": "2020-01-01T08:00:00Z", "id": "2c0fc9fa937b11eaa1b71c4d701ab86e", "modified_on": "2020-01-10T14:00:00Z"}, "required": ["id", "created_on", "modified_on"]}
```
