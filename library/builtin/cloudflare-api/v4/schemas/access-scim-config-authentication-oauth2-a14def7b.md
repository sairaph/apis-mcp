---
title: access_scim_config_authentication_oauth2
page_id: schema-access-scim-config-authentication-oauth2-a14def7b
path: schemas
description: Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning to an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_scim_config_authentication_oauth2

Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning to an application.

```yaml
{"description": "Attributes for configuring OAuth 2 authentication scheme for SCIM provisioning to an application.", "type": "object", "properties": {"authorization_url": {"description": "URL used to generate the auth code used during token generation.", "type": "string"}, "client_id": {"description": "Client ID used to authenticate when generating a token for authenticating with the remote SCIM service.", "type": "string"}, "client_secret": {"description": "Secret used to authenticate when generating a token for authenticating with the remove SCIM service.", "type": "string", "x-sensitive": true}, "scheme": {"description": "The authentication scheme to use when making SCIM requests to this application.", "type": "string", "enum": ["oauth2"]}, "scopes": {"description": "The authorization scopes to request when generating the token used to authenticate with the remove SCIM service.", "type": "array", "items": {"type": "string"}}, "token_url": {"description": "URL used to generate the token used to authenticate with the remote SCIM service.", "type": "string"}}, "required": ["scheme", "client_id", "client_secret", "authorization_url", "token_url"], "title": "OAuth 2"}
```
