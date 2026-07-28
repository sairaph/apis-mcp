---
title: tls-certificates-and-hostnames_sslsettings
page_id: schema-tls-certificates-and-hostnames-sslsettings-abfa0604
path: schemas
description: SSL specific settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_sslsettings

SSL specific settings.

```yaml
{"description": "SSL specific settings.", "type": "object", "properties": {"ciphers": {"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["ECDHE-RSA-AES128-GCM-SHA256", "AES128-SHA"], "uniqueItems": true}, "early_hints": {"description": "Whether or not Early Hints is enabled.", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}, "http2": {"description": "Whether or not HTTP2 is enabled.", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}, "min_tls_version": {"description": "The minimum TLS version supported.", "type": "string", "example": "1.2", "enum": ["1.0", "1.1", "1.2", "1.3"], "x-auditable": true}, "tls_1_3": {"description": "Whether or not TLS 1.3 is enabled.", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}}
```
