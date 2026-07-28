---
title: r2-slurper_api-v4-failure
page_id: schema-r2-slurper-api-v4-failure-3d14b7b0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_api-v4-failure

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/r2-slurper_api-v4-error"}, "messages": {"$ref": "#/components/schemas/r2-slurper_api-v4-message"}, "success": {"description": "Indicates if the API call was successful or not.", "type": "boolean", "example": false, "enum": [false], "x-auditable": true}}}
```
