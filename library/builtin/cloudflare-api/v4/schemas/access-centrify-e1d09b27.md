---
title: access_centrify
page_id: schema-access-centrify-e1d09b27
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_centrify

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config"}, {"$ref": "#/components/schemas/access_custom-claims-support"}, {"properties": {"centrify_account": {"description": "Your centrify account url", "type": "string", "example": "https://abc123.my.centrify.com/", "x-auditable": true}, "centrify_app_id": {"description": "Your centrify app id", "type": "string", "example": "exampleapp", "x-auditable": true}}, "type": "object"}]}}, "type": "object"}], "title": "Centrify"}
```
