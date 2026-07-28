---
title: zero-trust-gateway_inspection-settings
page_id: schema-zero-trust-gateway-inspection-settings-4d6b8243
path: schemas
description: Define the proxy inspection mode.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_inspection-settings

Define the proxy inspection mode.

```yaml
{"description": "Define the proxy inspection mode.", "type": "object", "properties": {"mode": {"description": "Define the proxy inspection mode.   1. static: Gateway applies static inspection to HTTP on TCP(80). With TLS decryption on, Gateway inspects HTTPS traffic on TCP(443) and UDP(443).   2. dynamic: Gateway applies protocol detection to inspect HTTP and HTTPS traffic on any port. TLS decryption must remain on to inspect HTTPS traffic.", "type": "string", "example": "static", "enum": ["static", "dynamic"], "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
