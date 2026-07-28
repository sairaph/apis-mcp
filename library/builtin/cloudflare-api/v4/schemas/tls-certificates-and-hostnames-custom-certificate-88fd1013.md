---
title: tls-certificates-and-hostnames_custom-certificate
page_id: schema-tls-certificates-and-hostnames-custom-certificate-88fd1013
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom-certificate

```yaml
{"type": "object", "properties": {"bundle_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_bundle_method"}, "custom_csr_id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_id"}, "expires_on": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_expires_on"}, "geo_restrictions": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_geo_restrictions"}, "hosts": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hosts"}, "id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, "issuer": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_issuer"}, "keyless_server": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless-certificate"}, "modified_on": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_modified_on"}, "policy_restrictions": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_policy_restrictions"}, "priority": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_priority"}, "signature": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_signature"}, "status": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_status"}, "uploaded_on": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_uploaded_on"}, "zone_id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, "required": ["id", "zone_id"]}
```
