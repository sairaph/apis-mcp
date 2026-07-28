---
title: application
page_id: schema-application-1fe28920
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# application

```yaml
{"title": "Application", "required": ["id", "object"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "name": {"maxLength": 5000, "type": "string", "description": "The name of the application.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["application"]}}, "description": "", "x-expandableFields": []}
```
