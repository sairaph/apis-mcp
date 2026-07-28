---
title: tls-certificates-and-hostnames_custom_csr
page_id: schema-tls-certificates-and-hostnames-custom-csr-73826532
path: schemas
description: A custom Certificate Signing Request (CSR).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_csr

A custom Certificate Signing Request (CSR).

```yaml
{"description": "A custom Certificate Signing Request (CSR).", "type": "object", "properties": {"common_name": {"description": "The common name (domain) for the CSR.", "type": "string", "example": "example.com", "maxLength": 64, "x-auditable": true}, "country": {"description": "Two-letter ISO 3166-1 alpha-2 country code.", "type": "string", "example": "US", "x-auditable": true}, "created_at": {"description": "When the CSR was created.", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00Z", "readOnly": true, "x-auditable": true}, "csr": {"description": "The PEM-encoded Certificate Signing Request.", "type": "string", "example": "-----BEGIN CERTIFICATE REQUEST-----\nMIICYzCCAUsCAQAwHj...", "readOnly": true, "x-auditable": true}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_identifier"}, "key_type": {"description": "The key algorithm used to generate the CSR.", "type": "string", "example": "rsa2048", "enum": ["rsa2048", "p256v1"], "x-auditable": true}, "locality": {"description": "City or locality name.", "type": "string", "example": "San Francisco", "x-auditable": true}, "organization": {"description": "Organization name.", "type": "string", "example": "Cloudflare, Inc.", "x-auditable": true}, "organizational_unit": {"description": "Organizational unit for the CSR subject.", "type": "string", "example": "Engineering", "x-auditable": true}, "sans": {"description": "Subject Alternative Names included in the CSR.", "type": "array", "items": {"type": "string"}, "example": ["example.com", "www.example.com"], "x-auditable": true}, "state": {"description": "State or province name.", "type": "string", "example": "California", "x-auditable": true}}, "required": ["id", "key_type", "created_at"]}
```
