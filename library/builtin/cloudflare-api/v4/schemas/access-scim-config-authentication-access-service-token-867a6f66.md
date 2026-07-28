---
title: access_scim_config_authentication_access_service_token
page_id: schema-access-scim-config-authentication-access-service-token-867a6f66
path: schemas
description: Attributes for configuring Access Service Token authentication scheme for SCIM provisioning to an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_scim_config_authentication_access_service_token

Attributes for configuring Access Service Token authentication scheme for SCIM provisioning to an application.

```yaml
{"description": "Attributes for configuring Access Service Token authentication scheme for SCIM provisioning to an application.", "type": "object", "properties": {"client_id": {"description": "Client ID of the Access service token used to authenticate with the remote service.", "type": "string"}, "client_secret": {"description": "Client secret of the Access service token used to authenticate with the remote service.", "type": "string", "x-sensitive": true}, "scheme": {"description": "The authentication scheme to use when making SCIM requests to this application.", "type": "string", "enum": ["access_service_token"]}}, "required": ["scheme", "client_id", "client_secret"], "title": "Access Service Token"}
```
