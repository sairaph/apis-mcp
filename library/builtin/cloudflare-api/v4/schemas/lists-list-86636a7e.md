---
title: lists_list
page_id: schema-lists-list-86636a7e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_list

```yaml
{"type": "object", "properties": {"created_on": {"$ref": "#/components/schemas/lists_created_on"}, "description": {"$ref": "#/components/schemas/lists_description"}, "id": {"$ref": "#/components/schemas/lists_list_id"}, "kind": {"$ref": "#/components/schemas/lists_kind"}, "modified_on": {"$ref": "#/components/schemas/lists_modified_on"}, "name": {"$ref": "#/components/schemas/lists_name"}, "num_items": {"$ref": "#/components/schemas/lists_num_items"}, "num_referencing_filters": {"$ref": "#/components/schemas/lists_num_referencing_filters"}}, "required": ["id", "name", "kind", "num_items", "num_referencing_filters", "created_on", "modified_on"]}
```
