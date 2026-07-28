---
title: tls-certificates-and-hostnames_custom_csr_create_request-2
page_id: schema-tls-certificates-and-hostnames-custom-csr-create-request-2-25ac4d03
path: schemas
description: Request body for creating a custom CSR.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_csr_create_request-2

Request body for creating a custom CSR.

```yaml
{"description": "Request body for creating a custom CSR.", "type": "object", "properties": {"common_name": {"description": "The common name (domain) for the CSR. Must be at most 64 characters.", "type": "string", "example": "example.com", "maxLength": 64}, "country": {"description": "Two-letter ISO 3166-1 alpha-2 country code.", "type": "string", "example": "US"}, "description": {"description": "Optional description for the CSR.", "type": "string", "example": "CSR for example.com wildcard"}, "key_type": {"description": "Key algorithm to use for the CSR. Defaults to rsa2048 if not specified.", "type": "string", "example": "rsa2048", "default": "rsa2048", "enum": ["rsa2048", "p256v1"]}, "locality": {"description": "City or locality name.", "type": "string", "example": "San Francisco"}, "name": {"description": "Human-readable name for the CSR.", "type": "string", "example": "My Custom CSR"}, "organization": {"description": "Organization name.", "type": "string", "example": "Cloudflare, Inc."}, "organizational_unit": {"description": "Organizational unit name.", "type": "string", "example": "Engineering"}, "sans": {"description": "Subject Alternative Names for the CSR. At least one SAN is required.", "type": "array", "items": {"maxLength": 256, "type": "string"}, "example": ["example.com", "www.example.com"], "maxItems": 150, "minItems": 1}, "state": {"description": "State or province name.", "type": "string", "example": "California"}}, "required": ["country", "state", "locality", "organization", "common_name", "sans"]}
```
