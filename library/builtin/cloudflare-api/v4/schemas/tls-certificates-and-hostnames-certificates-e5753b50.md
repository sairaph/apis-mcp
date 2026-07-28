---
title: tls-certificates-and-hostnames_certificates
page_id: schema-tls-certificates-and-hostnames-certificates-e5753b50
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_certificates

```yaml
{"type": "object", "properties": {"certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate-4"}, "csr": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_csr"}, "expires_on": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_expires_on-3"}, "hostnames": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostnames"}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, "request_type": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_request_type"}, "requested_validity": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_requested_validity"}}, "required": ["hostnames", "csr", "requested_validity", "request_type"]}
```
