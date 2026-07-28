---
title: access_oidc
page_id: schema-access-oidc-05e09a07
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_oidc

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config"}, {"$ref": "#/components/schemas/access_custom-claims-support"}, {"properties": {"auth_url": {"description": "The authorization_endpoint URL of your IdP", "type": "string", "example": "https://accounts.google.com/o/oauth2/auth"}, "certs_url": {"description": "The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens", "type": "string", "example": "https://www.googleapis.com/oauth2/v3/certs"}, "pkce_enabled": {"description": "Enable Proof Key for Code Exchange (PKCE)", "type": "boolean"}, "scopes": {"description": "OAuth scopes", "type": "array", "items": {"type": "string"}, "example": ["openid", "email", "profile"]}, "token_url": {"description": "The token_endpoint URL of your IdP", "type": "string", "example": "https://accounts.google.com/o/oauth2/token"}}, "type": "object"}]}}, "type": "object"}], "title": "Generic OAuth"}
```
