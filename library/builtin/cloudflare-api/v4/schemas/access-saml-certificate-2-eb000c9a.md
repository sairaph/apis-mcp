---
title: access_saml_certificate-2
page_id: schema-access-saml-certificate-2-eb000c9a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saml_certificate-2

```yaml
{"type": "object", "properties": {"is_current": {"description": "Indicates whether the certificate can be used for IdP configuration.", "type": "boolean", "example": true}, "not_after": {"description": "Certificate expiration date", "type": "string", "format": "date-time", "example": "2027-03-21T12:00:00Z"}, "public_certificate": {"description": "The public certificate in PEM format", "type": "string", "example": "-----BEGIN CERTIFICATE-----\nMIIGAjCCA+qgAwIBAgIJAI7kymlF7CWT...\n...certificate content...\n-----END CERTIFICATE-----\n"}, "uid": {"description": "Unique identifier for the certificate", "type": "string", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}}, "required": ["uid", "public_certificate", "not_after", "is_current"]}
```
