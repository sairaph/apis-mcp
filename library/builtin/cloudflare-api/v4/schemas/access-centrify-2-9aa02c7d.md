---
title: access_centrify-2
page_id: schema-access-centrify-2-9aa02c7d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_centrify-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config-2"}, {"properties": {"centrify_account": {"description": "Your centrify account url", "type": "string", "example": "https://abc123.my.centrify.com/"}, "centrify_app_id": {"description": "Your centrify app id", "type": "string", "example": "exampleapp"}}, "type": "object"}]}}, "type": "object"}], "title": "Centrify"}
```
