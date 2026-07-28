---
title: access_saml_certificate
page_id: schema-access-saml-certificate-b81491f9
path: schemas
description: A single SAML encryption certificate with metadata
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saml_certificate

A single SAML encryption certificate with metadata

```yaml
{"description": "A single SAML encryption certificate with metadata", "type": "object", "properties": {"is_current": {"description": "Indicates whether this is the currently active certificate", "type": "boolean", "example": true}, "not_after": {"description": "Certificate expiration date. Certificates are automatically rotated 30 days before expiration.", "type": "string", "format": "date-time", "example": "2027-05-07T19:11:00Z"}, "public_certificate": {"description": "PEM-encoded X.509 certificate containing the public key.\nConfigure this certificate in your external SAML Identity Provider to enable encryption.\n", "type": "string", "example": "-----BEGIN CERTIFICATE-----\nMIIEpzCCA4+gAwIBAgIUTh2VSDDJ0oB/gabio6j1L9QwWoUwDQYJKoZIhvcNAQEL\n...\n-----END CERTIFICATE-----\n"}, "uid": {"description": "Unique identifier for the certificate", "type": "string", "format": "uuid", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}}, "required": ["uid", "public_certificate", "not_after", "is_current"]}
```
