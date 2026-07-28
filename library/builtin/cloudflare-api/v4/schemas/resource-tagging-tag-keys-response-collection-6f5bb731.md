---
title: resource-tagging_tag_keys_response_collection
page_id: schema-resource-tagging-tag-keys-response-collection-6f5bb731
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tag_keys_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/resource-tagging_api-response-common"}, {"properties": {"result": {"description": "Contains an array of distinct tag keys.", "type": "array", "items": {"type": "string"}, "example": ["environment", "team", "region"]}, "result_info": {"$ref": "#/components/schemas/resource-tagging_cursor_result_info"}}}]}
```
