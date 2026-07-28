---
title: access_pingone
page_id: schema-access-pingone-7693d413
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_pingone

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config"}, {"$ref": "#/components/schemas/access_custom-claims-support"}, {"properties": {"ping_env_id": {"description": "Your PingOne environment identifier", "type": "string", "example": "342b5660-0c32-4936-a5a4-ce21fae57b0a", "x-auditable": true}}, "type": "object"}]}}, "type": "object"}], "title": "PingOne"}
```
