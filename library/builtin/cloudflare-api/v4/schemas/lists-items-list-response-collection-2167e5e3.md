---
title: lists_items-list-response-collection
page_id: schema-lists-items-list-response-collection-2167e5e3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_items-list-response-collection

```yaml
{"allOf": [{"properties": {"result": {"$ref": "#/components/schemas/lists_items"}, "result_info": {"type": "object", "properties": {"cursors": {"type": "object", "properties": {"after": {"type": "string", "example": "yyy", "x-auditable": true}, "before": {"type": "string", "example": "xxx", "x-auditable": true}}}}}}, "type": "object"}, {"$ref": "#/components/schemas/lists_api-response-collection"}]}
```
