---
title: organizations-api_CreateSingleMember
page_id: schema-organizations-api-createsinglemember-cc39980e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_CreateSingleMember

```yaml
{"type": "object", "properties": {"status": {"type": "string", "enum": ["active", "canceled"]}, "user": {"type": "object", "properties": {"email": {"type": "string"}}, "required": ["email"]}}, "required": ["user"]}
```
