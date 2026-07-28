---
title: moq_cloudflare_jwt_issuer
page_id: schema-moq-cloudflare-jwt-issuer-942a85bb
path: schemas
description: Cloudflare-managed, Ed25519-signed JWT tokens (the only V1 type).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_cloudflare_jwt_issuer

Cloudflare-managed, Ed25519-signed JWT tokens (the only V1 type).

```yaml
{"description": "Cloudflare-managed, Ed25519-signed JWT tokens (the only V1 type).", "type": "object", "properties": {"cloudflare_tokens": {"description": "Always present ([] when empty).", "type": "array", "items": {"$ref": "#/components/schemas/moq_cloudflare_jwt_token"}}, "issuer": {"type": "string", "enum": ["cloudflare"]}, "type": {"type": "string", "enum": ["cloudflare_jwt"]}}, "required": ["type", "issuer", "cloudflare_tokens"]}
```
