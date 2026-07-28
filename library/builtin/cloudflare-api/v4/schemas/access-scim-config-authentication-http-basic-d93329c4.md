---
title: access_scim_config_authentication_http_basic
page_id: schema-access-scim-config-authentication-http-basic-d93329c4
path: schemas
description: Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_scim_config_authentication_http_basic

Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.

```yaml
{"description": "Attributes for configuring HTTP Basic authentication scheme for SCIM provisioning to an application.", "type": "object", "properties": {"password": {"description": "Password used to authenticate with the remote SCIM service.", "type": "string", "x-sensitive": true}, "scheme": {"description": "The authentication scheme to use when making SCIM requests to this application.", "type": "string", "enum": ["httpbasic"]}, "user": {"description": "User name used to authenticate with the remote SCIM service.", "type": "string"}}, "required": ["scheme", "user", "password"], "title": "HTTP Basic"}
```
