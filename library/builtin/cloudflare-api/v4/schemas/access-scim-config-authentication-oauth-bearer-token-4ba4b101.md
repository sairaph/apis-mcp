---
title: access_scim_config_authentication_oauth_bearer_token
page_id: schema-access-scim-config-authentication-oauth-bearer-token-4ba4b101
path: schemas
description: Attributes for configuring OAuth Bearer Token authentication scheme for SCIM provisioning to an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_scim_config_authentication_oauth_bearer_token

Attributes for configuring OAuth Bearer Token authentication scheme for SCIM provisioning to an application.

```yaml
{"description": "Attributes for configuring OAuth Bearer Token authentication scheme for SCIM provisioning to an application.", "type": "object", "properties": {"scheme": {"description": "The authentication scheme to use when making SCIM requests to this application.", "type": "string", "enum": ["oauthbearertoken"]}, "token": {"description": "Token used to authenticate with the remote SCIM service.", "type": "string", "x-sensitive": true}}, "required": ["scheme", "token"], "title": "OAuth Bearer Token"}
```
