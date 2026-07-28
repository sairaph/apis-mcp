---
title: aaa_schemas-response_collection
page_id: schema-aaa-schemas-response-collection-9d97554b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_schemas-response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-2"}, {"properties": {"result": {"type": "object", "example": {"email": {"eligible": true, "ready": true, "type": "email"}}, "additionalProperties": {"items": {"$ref": "#/components/schemas/aaa_eligibility"}, "type": "array"}}}}]}
```
