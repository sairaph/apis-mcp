---
title: iam_oauth_client_common
page_id: schema-iam-oauth-client-common-95f3d2c2
path: schemas
description: Fields shared by OAuth client responses and create/update requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_common

Fields shared by OAuth client responses and create/update requests.

```yaml
{"description": "Fields shared by OAuth client responses and create/update requests.", "type": "object", "properties": {"allowed_cors_origins": {"description": "Array of allowed CORS origins.", "type": "array", "items": {"type": "string"}, "example": ["https://example.com"], "x-auditable": true}, "client_name": {"description": "Human-readable name of the OAuth client.", "type": "string", "example": "My OAuth App", "x-auditable": true}, "client_uri": {"description": "URL of the home page of the client.", "type": "string", "example": "https://example.com", "x-auditable": true}, "grant_types": {"$ref": "#/components/schemas/iam_oauth_client_grant_types"}, "logo_uri": {"description": "URL of the client's logo.", "type": "string", "example": "https://example.com/logo.png", "x-auditable": true}, "policy_uri": {"description": "URL that points to a privacy policy document.", "type": "string", "example": "https://example.com/privacy", "x-auditable": true}, "post_logout_redirect_uris": {"description": "Array of allowed post-logout redirect URIs.", "type": "array", "items": {"type": "string"}, "example": ["https://example.com/logout"], "x-auditable": true}, "redirect_uris": {"description": "Array of allowed redirect URIs for the client.", "type": "array", "items": {"type": "string"}, "example": ["https://example.com/callback"], "x-auditable": true}, "response_types": {"$ref": "#/components/schemas/iam_oauth_client_response_types"}, "scopes": {"description": "Array of OAuth scopes the client is allowed to request. Colon-delimited scopes are not accepted. Dot-delimited scopes are validated against available OAuth API scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and `openid` are added or removed automatically based on `grant_types` and `response_types`.", "type": "array", "items": {"type": "string"}, "example": ["account.read"], "x-auditable": true}, "token_endpoint_auth_method": {"$ref": "#/components/schemas/iam_oauth_client_token_endpoint_auth_method"}, "tos_uri": {"description": "URL that points to a terms of service document.", "type": "string", "example": "https://example.com/tos", "x-auditable": true}}, "title": "OAuth Client Common Fields"}
```
