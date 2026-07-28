---
title: secret_service_resource_scope
page_id: schema-secret-service-resource-scope-d2d0c17f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# secret_service_resource_scope

```yaml
{"title": "SecretServiceResourceScope", "required": ["type"], "type": "object", "properties": {"type": {"type": "string", "description": "The secret scope type.", "enum": ["account", "user"]}, "user": {"maxLength": 5000, "type": "string", "description": "The user ID, if type is set to \"user\""}}, "description": "", "x-expandableFields": []}
```
