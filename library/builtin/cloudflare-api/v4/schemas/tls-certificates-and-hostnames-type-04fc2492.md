---
title: tls-certificates-and-hostnames_type
page_id: schema-tls-certificates-and-hostnames-type-04fc2492
path: schemas
description: The type 'legacy_custom' enables support for legacy clients which do not include SNI in the TLS handshake.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_type

The type 'legacy_custom' enables support for legacy clients which do not include SNI in the TLS handshake.

```yaml
{"description": "The type 'legacy_custom' enables support for legacy clients which do not include SNI in the TLS handshake.", "type": "string", "example": "sni_custom", "default": "legacy_custom", "enum": ["legacy_custom", "sni_custom"], "x-auditable": true}
```
