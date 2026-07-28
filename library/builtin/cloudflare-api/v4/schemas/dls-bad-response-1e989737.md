---
title: dls_bad_response
page_id: schema-dls-bad-response-1e989737
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_bad_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/dls_response"}, {"properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/dls_coded_message"}, "minLength": 1}, "result": {"$ref": "#/components/schemas/dls_none"}}, "required": ["errors"]}]}
```
