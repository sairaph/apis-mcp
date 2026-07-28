---
title: iam_scim_authentication_scheme
page_id: schema-iam-scim-authentication-scheme-b3665543
path: schemas
description: An authentication method supported by the SCIM service.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_authentication_scheme

An authentication method supported by the SCIM service.

```yaml
{"description": "An authentication method supported by the SCIM service.", "type": "object", "properties": {"description": {"description": "A description of the authentication scheme.", "type": "string", "example": "Authentication via OAuth 2.0 Bearer Token"}, "documentationUri": {"description": "An HTTP-addressable URL pointing to the authentication scheme documentation.", "type": "string"}, "name": {"description": "The common authentication scheme name.", "type": "string", "example": "OAuth Bearer Token"}, "primary": {"description": "Indicates if this is the primary authentication scheme.", "type": "boolean", "example": true}, "specUri": {"description": "An HTTP-addressable URL pointing to the authentication scheme specification.", "type": "string", "example": "http://www.rfc-editor.org/info/rfc6750"}, "type": {"description": "The authentication scheme type.", "type": "string", "example": "oauthbearertoken"}}, "required": ["type", "name", "description"], "title": "SCIM Authentication Scheme"}
```
