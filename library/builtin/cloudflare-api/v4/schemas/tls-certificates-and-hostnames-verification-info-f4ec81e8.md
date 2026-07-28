---
title: tls-certificates-and-hostnames_verification_info
page_id: schema-tls-certificates-and-hostnames-verification-info-f4ec81e8
path: schemas
description: Certificate's required verification information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_verification_info

Certificate's required verification information.

```yaml
{"description": "Certificate's required verification information.", "type": "object", "properties": {"record_name": {"description": "Name of CNAME record.", "type": "string", "format": "hostname", "example": "b3b90cfedd89a3e487d3e383c56c4267.example.com", "enum": ["record_name", "http_url", "cname", "txt_name"], "x-auditable": true}, "record_target": {"description": "Target of CNAME record.", "type": "string", "format": "hostname", "example": "6979be7e4cfc9e5c603e31df7efac9cc60fee82d.comodoca.com", "enum": ["record_value", "http_body", "cname_target", "txt_value"], "x-auditable": true}}}
```
