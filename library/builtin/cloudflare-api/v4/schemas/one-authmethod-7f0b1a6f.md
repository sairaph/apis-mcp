---
title: one_AuthMethod
page_id: schema-one-authmethod-7f0b1a6f
path: schemas
description: Authentication method available for a vendor.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_AuthMethod

Authentication method available for a vendor.

```yaml
{"description": "Authentication method available for a vendor.", "type": "object", "properties": {"display_name": {"description": "Human-readable auth method name.", "type": "string"}, "id": {"description": "Auth method identifier.", "type": "string"}, "is_default": {"description": "Whether this is the default auth method.", "type": "boolean"}, "supported_environments": {"description": "Environments this auth method supports.", "type": "array", "items": {"type": "string"}}}, "required": ["display_name", "id", "is_default", "supported_environments"]}
```
