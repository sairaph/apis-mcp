---
title: magic-visibility-pcaps_pcaps_collection_response
page_id: schema-magic-visibility-pcaps-pcaps-collection-response-e1ba27bc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_pcaps_collection_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_response_simple"}, {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_response_full"}]}}}, "type": "object"}]}
```
