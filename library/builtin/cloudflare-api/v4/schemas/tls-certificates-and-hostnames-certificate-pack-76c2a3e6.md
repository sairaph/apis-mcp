---
title: tls-certificates-and-hostnames_certificate_pack
page_id: schema-tls-certificates-and-hostnames-certificate-pack-76c2a3e6
path: schemas
description: A certificate pack with all its properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_certificate_pack

A certificate pack with all its properties.

```yaml
{"description": "A certificate pack with all its properties.", "type": "object", "properties": {"certificate_authority": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_authority-2"}, "certificates": {"description": "Array of certificates in this pack.", "type": "array", "items": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_pack_certificate"}}, "cloudflare_branding": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_cloudflare_branding"}, "dcv_delegation_records": {"description": "DCV Delegation records for domain validation.", "type": "array", "items": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validation_record"}, "readOnly": true}, "hosts": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hosts-2"}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, "primary_certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_primary"}, "status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_status-5"}, "type": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_type-2"}, "validation_errors": {"description": "Domain validation errors that have been received by the certificate authority (CA).", "type": "array", "items": {"properties": {"message": {"description": "A domain validation error.", "type": "string", "example": "SERVFAIL looking up CAA for app.example.com", "x-auditable": true}}, "type": "object"}}, "validation_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validation_method"}, "validation_records": {"description": "Certificates' validation records.", "type": "array", "items": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validation_record"}}, "validity_days": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validity_days"}}, "required": ["id", "type", "hosts", "certificates", "status"]}
```
