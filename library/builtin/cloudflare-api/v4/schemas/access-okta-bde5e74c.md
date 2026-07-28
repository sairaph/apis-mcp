---
title: access_okta
page_id: schema-access-okta-bde5e74c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_okta

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config"}, {"$ref": "#/components/schemas/access_custom-claims-support"}, {"properties": {"authorization_server_id": {"description": "Your okta authorization server id", "type": "string", "example": "aus9o8wzkhckw9TLa0h7z"}, "okta_account": {"description": "Your okta account url", "type": "string", "example": "https://dev-abc123.oktapreview.com", "x-auditable": true}}, "type": "object"}]}}, "type": "object"}], "title": "Okta"}
```
