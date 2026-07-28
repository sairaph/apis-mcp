---
title: kamino_environments_response
page_id: schema-kamino-environments-response-4d3ddda5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# kamino_environments_response

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/kamino_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/kamino_message"}}, "result": {"$ref": "#/components/schemas/kamino_environments_result"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}
```
