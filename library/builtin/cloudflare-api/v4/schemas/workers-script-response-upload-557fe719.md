---
title: workers_script-response-upload
page_id: schema-workers-script-response-upload-557fe719
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_script-response-upload

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_script-response"}, {"properties": {"entry_point": {"description": "The entry point for the script.", "type": "string", "example": "index.js", "readOnly": true}, "startup_time_ms": {"type": "integer", "example": 10, "readOnly": true, "x-auditable": true}}, "required": ["startup_time_ms"], "type": "object"}]}
```
