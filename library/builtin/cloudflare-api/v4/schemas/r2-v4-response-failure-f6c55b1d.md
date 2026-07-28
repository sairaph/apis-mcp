---
title: r2_v4_response_failure
page_id: schema-r2-v4-response-failure-f6c55b1d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_v4_response_failure

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/r2_errors"}, "messages": {"$ref": "#/components/schemas/r2_messages"}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
