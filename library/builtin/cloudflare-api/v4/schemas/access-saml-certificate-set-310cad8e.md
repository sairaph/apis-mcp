---
title: access_saml_certificate_set
page_id: schema-access-saml-certificate-set-310cad8e
path: schemas
description: A SAML encryption certificate set containing current and optionally previous certificates for encryption key rotation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saml_certificate_set

A SAML encryption certificate set containing current and optionally previous certificates for encryption key rotation.

```yaml
{"description": "A SAML encryption certificate set containing current and optionally previous certificates for encryption key rotation.", "type": "object", "properties": {"created_at": {"description": "Timestamp when the certificate set was created", "type": "string", "format": "date-time", "example": "2026-05-07T19:16:19.821162Z"}, "current_certificate": {"description": "The currently active certificate used for encrypting SAML assertions", "allOf": [{"$ref": "#/components/schemas/access_saml_certificate"}]}, "previous_certificate": {"description": "The previous certificate, maintained during rotation to ensure continuity. Null if no rotation has occurred. Mirrors the structure of `saml_certificate`.", "type": "object", "nullable": true}, "uid": {"description": "Unique identifier for the certificate set", "type": "string", "format": "uuid", "example": "c409ef44-e72c-41c8-8c0b-278c8a6f4fd8"}, "updated_at": {"description": "Timestamp when the certificate set was last updated (e.g., during rotation)", "type": "string", "format": "date-time", "example": "2026-05-07T19:16:19.821162Z"}}, "required": ["uid", "created_at", "updated_at"]}
```
