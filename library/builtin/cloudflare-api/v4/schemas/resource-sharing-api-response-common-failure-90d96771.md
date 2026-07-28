---
title: resource-sharing_api-response-common-failure
page_id: schema-resource-sharing-api-response-common-failure-90d96771
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 7003, "message": "No route for the URI"}], "allOf": [{"$ref": "#/components/schemas/resource-sharing_v4errors"}], "minLength": 1}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false, "x-auditable": true}}, "required": ["success", "errors", "result"]}
```
