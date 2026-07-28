---
title: tls-certificates-and-hostnames_custom-trust-store
page_id: schema-tls-certificates-and-hostnames-custom-trust-store-9ad88169
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom-trust-store

```yaml
{"type": "object", "properties": {"certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate-3"}, "expires_on": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_expires_on-2"}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, "issuer": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_issuer"}, "signature": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_signature"}, "status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_status-7"}, "updated_at": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_updated_at-2"}, "uploaded_on": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_uploaded_on"}}, "required": ["id", "status", "issuer", "signature", "certificate", "expires_on", "uploaded_on", "updated_at"]}
```
