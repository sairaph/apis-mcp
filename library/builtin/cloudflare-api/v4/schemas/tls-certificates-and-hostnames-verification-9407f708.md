---
title: tls-certificates-and-hostnames_verification
page_id: schema-tls-certificates-and-hostnames-verification-9407f708
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_verification

```yaml
{"type": "object", "properties": {"brand_check": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_brand_check"}, "cert_pack_uuid": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_cert_pack_uuid"}, "certificate_status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_status"}, "signature": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_signature-2"}, "validation_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validation_method-2"}, "verification_info": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_verification_info"}, "verification_status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_verification_status"}, "verification_type": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_verification_type"}}, "required": ["certificate_status"]}
```
