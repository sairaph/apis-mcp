---
title: tls-certificates-and-hostnames_certificate_pack_certificate
page_id: schema-tls-certificates-and-hostnames-certificate-pack-certificate-ae8a6dc6
path: schemas
description: An individual certificate within a certificate pack.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_certificate_pack_certificate

An individual certificate within a certificate pack.

```yaml
{"description": "An individual certificate within a certificate pack.", "type": "object", "properties": {"bundle_method": {"description": "Certificate bundle method.", "type": "string", "example": "ubiquitous", "x-auditable": true}, "expires_on": {"description": "When the certificate from the authority expires.", "type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z", "x-auditable": true}, "geo_restrictions": {"description": "Specify the region where your private key can be held locally.", "type": "object", "properties": {"label": {"type": "string", "example": "us", "enum": ["us", "eu", "highest_security"], "x-auditable": true}}}, "hosts": {"description": "Hostnames covered by this certificate.", "type": "array", "items": {"type": "string"}, "example": ["example.com", "*.example.com"], "x-auditable": true}, "id": {"description": "Certificate identifier.", "type": "string", "example": "7e7b8deba8538af625850b7b2530034c", "x-auditable": true}, "issuer": {"description": "The certificate authority that issued the certificate.", "type": "string", "example": "Let's Encrypt", "x-auditable": true}, "modified_on": {"description": "When the certificate was last modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00Z", "x-auditable": true}, "priority": {"description": "The order/priority in which the certificate will be used.", "type": "number", "x-auditable": true}, "signature": {"description": "The type of hash used for the certificate.", "type": "string", "example": "ECDSAWithSHA256", "x-auditable": true}, "status": {"description": "Certificate status.", "type": "string", "example": "active", "x-auditable": true}, "uploaded_on": {"description": "When the certificate was uploaded to Cloudflare.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00Z", "x-auditable": true}, "zone_id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, "required": ["id", "hosts", "status"]}
```
