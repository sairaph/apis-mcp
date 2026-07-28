---
title: zones_ciphers
page_id: schema-zones-ciphers-e52ead36
path: schemas
description: An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_ciphers

An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.

```yaml
{"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "default": [], "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "ciphers", "enum": ["ciphers"]}, "value": {"$ref": "#/components/schemas/zones_ciphers_value"}}}], "title": "Zone ciphers allowed for TLS termination"}
```
