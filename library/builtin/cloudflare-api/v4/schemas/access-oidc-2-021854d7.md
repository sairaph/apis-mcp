---
title: access_oidc-2
page_id: schema-access-oidc-2-021854d7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_oidc-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"allOf": [{"$ref": "#/components/schemas/access_generic-oauth-config-2"}, {"properties": {"auth_url": {"description": "The authorization_endpoint URL of your IdP", "type": "string", "example": "https://accounts.google.com/o/oauth2/auth"}, "certs_url": {"description": "The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens", "type": "string", "example": "https://www.googleapis.com/oauth2/v3/certs"}, "claims": {"description": "List of custom claims that will be pulled from your id_token and added to your signed Access JWT token.", "type": "array", "items": {"type": "string"}, "example": ["given_name", "locale"]}, "scopes": {"description": "OAuth scopes", "type": "array", "items": {"type": "string"}, "example": ["openid", "email", "profile"]}, "token_url": {"description": "The token_endpoint URL of your IdP", "type": "string", "example": "https://accounts.google.com/o/oauth2/token"}}, "type": "object"}]}}, "type": "object"}], "title": "Generic OAuth"}
```
