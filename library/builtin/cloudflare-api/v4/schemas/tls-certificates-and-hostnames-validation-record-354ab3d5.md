---
title: tls-certificates-and-hostnames_validation_record
page_id: schema-tls-certificates-and-hostnames-validation-record-354ab3d5
path: schemas
description: Certificate's required validation record.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_validation_record

Certificate's required validation record.

```yaml
{"description": "Certificate's required validation record.", "type": "object", "properties": {"cname": {"description": "The CNAME record hostname for DCV delegation.", "type": "string", "example": "_acme-challenge.example.com", "readOnly": true, "x-auditable": true}, "cname_target": {"description": "The CNAME record target value for DCV delegation.", "type": "string", "example": "dcv.cloudflare.com", "readOnly": true, "x-auditable": true}, "emails": {"description": "The set of email addresses that the certificate authority (CA) will use to complete domain validation.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["administrator@example.com", "webmaster@example.com"]}, "http_body": {"description": "The content that the certificate authority (CA) will expect to find at the http_url during the domain validation.", "type": "string", "example": "ca3-574923932a82475cb8592200f1a2a23d"}, "http_url": {"description": "The url that will be checked during domain validation.", "type": "string", "example": "http://app.example.com/.well-known/pki-validation/ca3-da12a1c25e7b48cf80408c6c1763b8a2.txt"}, "status": {"description": "Status of the validation record.", "type": "string", "example": "pending", "readOnly": true}, "txt_name": {"description": "The hostname that the certificate authority (CA) will check for a TXT record during domain validation .", "type": "string", "example": "_acme-challenge.app.example.com", "x-auditable": true}, "txt_value": {"description": "The TXT record that the certificate authority (CA) will check during domain validation.", "type": "string", "example": "810b7d5f01154524b961ba0cd578acc2"}}}
```
