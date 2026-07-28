---
title: aaa_api-response-common-failure-2
page_id: schema-aaa-api-response-common-failure-2-1f22ae2a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_api-response-common-failure-2

```yaml
{"type": "object", "properties": {"errors": {"description": "A list of error messages.", "type": "array", "items": {"properties": {"message": {"description": "A text description of this message.", "type": "string", "example": "No route for the URI"}}, "required": ["message"], "type": "object", "uniqueItems": true}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "type": "object", "uniqueItems": true}, "example": []}, "success": {"description": "Indicates whether the API call was failed", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors"]}
```
