---
title: mq_api-v4-success
page_id: schema-mq-api-v4-success-66c0d5ee
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_api-v4-success

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/mq_api-v4-error"}, "messages": {"$ref": "#/components/schemas/mq_api-v4-message"}, "success": {"description": "Indicates if the API call was successful or not.", "type": "boolean", "enum": [true], "x-auditable": true}}}
```
