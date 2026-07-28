---
title: access_saml_certificate_set-2
page_id: schema-access-saml-certificate-set-2-dff29edb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saml_certificate_set-2

```yaml
{"type": "object", "properties": {"created_at": {"description": "When the certificate set was created", "type": "string", "format": "date-time", "example": "2024-03-21T10:30:00Z"}, "current_certificate": {"description": "The current active certificate", "allOf": [{"$ref": "#/components/schemas/access_saml_certificate-2"}]}, "previous_certificate": {"description": "The previous certificate (maintained during rotation period). May be null when no rotation has occurred. Mirrors the structure of `saml_certificate`.", "type": "object", "nullable": true}, "uid": {"description": "Unique identifier for the certificate set", "type": "string", "example": "a5bb4b3f-c2d1-4e6a-8f9b-1d3e4f5a6b7c"}, "updated_at": {"description": "When the certificate set was last updated", "type": "string", "format": "date-time", "example": "2024-03-21T10:30:00Z"}}, "required": ["uid", "created_at", "updated_at"]}
```
