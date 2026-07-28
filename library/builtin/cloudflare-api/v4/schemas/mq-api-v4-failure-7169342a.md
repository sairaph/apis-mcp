---
title: mq_api-v4-failure
page_id: schema-mq-api-v4-failure-7169342a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_api-v4-failure

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/mq_api-v4-error"}, "messages": {"$ref": "#/components/schemas/mq_api-v4-message"}, "success": {"description": "Indicates if the API call was successful or not.", "type": "boolean", "example": false, "enum": [false], "x-auditable": true}}}
```
