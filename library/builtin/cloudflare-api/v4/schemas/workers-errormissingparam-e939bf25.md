---
title: workers_ErrorMissingParam
page_id: schema-workers-errormissingparam-e939bf25
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorMissingParam

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that a required URL parameter is missing.", "type": "integer", "enum": [10003]}, "message": {"description": "Message explaining which required parameter is missing and suggesting to check the URL.", "type": "string"}}, "required": ["code", "message"]}
```
