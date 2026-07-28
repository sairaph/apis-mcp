---
title: tls-certificates-and-hostnames_hostname_association
page_id: schema-tls-certificates-and-hostnames-hostname-association-b1e315e8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_hostname_association

```yaml
{"type": "object", "properties": {"hostnames": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_hostnames-2"}, "mtls_certificate_id": {"description": "The UUID for a certificate that was uploaded to the mTLS Certificate Management endpoint. If no mtls_certificate_id is given, the hostnames will be associated to your active Cloudflare Managed CA.", "type": "string", "maxLength": 36, "minLength": 36, "x-auditable": true}}}
```
