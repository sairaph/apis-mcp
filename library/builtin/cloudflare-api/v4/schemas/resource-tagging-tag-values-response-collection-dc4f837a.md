---
title: resource-tagging_tag_values_response_collection
page_id: schema-resource-tagging-tag-values-response-collection-dc4f837a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tag_values_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/resource-tagging_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"type": "string"}, "example": ["production", "staging"]}, "result_info": {"$ref": "#/components/schemas/resource-tagging_cursor_result_info"}}}]}
```
