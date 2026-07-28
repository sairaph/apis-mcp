---
title: workers_script-response-collection
page_id: schema-workers-script-response-collection-25d8fbba
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_script-response-collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/workers_script-response"}, {"properties": {"routes": {"description": "Routes associated with the Worker.", "type": "array", "items": {"$ref": "#/components/schemas/workers_route"}, "nullable": true}}, "type": "object"}]}}}, "required": ["result"], "type": "object"}]}
```
