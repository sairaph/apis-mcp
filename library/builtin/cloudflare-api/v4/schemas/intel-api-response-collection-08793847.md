---
title: intel_api-response-collection
page_id: schema-intel-api-response-collection-08793847
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_api-response-collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/intel_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"oneOf": [{"type": "string"}, {"type": "object"}]}, "nullable": true}, "result_info": {"$ref": "#/components/schemas/intel_result_info"}}, "type": "object"}]}
```
