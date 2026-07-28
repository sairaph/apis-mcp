---
title: resource-tagging_tag_key_summary_response_collection
page_id: schema-resource-tagging-tag-key-summary-response-collection-d9a618b2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tag_key_summary_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/resource-tagging_api-response-common"}, {"properties": {"result": {"description": "Contains an array of tag keys with their distinct values.", "type": "array", "items": {"$ref": "#/components/schemas/resource-tagging_tag_key_summary"}, "example": [{"key": "environment", "values": ["production", "staging"]}, {"key": "team", "values": ["engineering", "platform"]}]}, "result_info": {"$ref": "#/components/schemas/resource-tagging_cursor_result_info"}}}]}
```
