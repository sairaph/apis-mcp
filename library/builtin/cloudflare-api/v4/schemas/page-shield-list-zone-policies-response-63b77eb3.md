---
title: page-shield_list-zone-policies-response
page_id: schema-page-shield-list-zone-policies-response-63b77eb3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_list-zone-policies-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/page-shield_api-list-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/page-shield_policy-with-id"}}}, "required": ["result"], "type": "object"}]}
```
