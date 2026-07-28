---
title: access_pingone-2
page_id: schema-access-pingone-2-bd738331
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_pingone-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config-2"}, {"properties": {"ping_env_id": {"description": "Your PingOne environment identifier", "type": "string", "example": "342b5660-0c32-4936-a5a4-ce21fae57b0a"}}, "type": "object"}]}}, "type": "object"}], "title": "PingOne"}
```
