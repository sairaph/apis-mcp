---
title: workers_completed-upload-assets-response
page_id: schema-workers-completed-upload-assets-response-d1521269
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_completed-upload-assets-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"jwt": {"description": "A \"completion\" JWT which can be redeemed when creating a Worker version.", "type": "string", "x-sensitive": true}}}}, "type": "object"}]}
```
