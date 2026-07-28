---
title: moq_cloudflare_jwt_token
page_id: schema-moq-cloudflare-jwt-token-e2f62104
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_cloudflare_jwt_token

```yaml
{"type": "object", "properties": {"created": {"type": "string", "format": "date-time"}, "expires": {"description": "Mandatory; no more than 1 year after `created`.", "type": "string", "format": "date-time"}, "jti": {"description": "Token identity and registry key (32 hex chars).", "type": "string", "example": "f3a1b2c3d4e5f67890a1b2c3d4e5f678"}, "label": {"description": "Optional, customer-set.", "type": "string", "example": "primary-encoder"}, "operations": {"description": "Signed allowlist of what the token may do. V1 coarse roles; the array\nform extends to fine-grained MoQT message names later without a\nbreaking change.\n", "type": "array", "items": {"enum": ["publish", "subscribe"], "type": "string"}, "example": ["publish", "subscribe"], "minItems": 1}, "secret": {"description": "The signed JWT. Present ONLY in create / auto-create responses (shown\nonce); never returned by list, never stored.\n", "type": "string", "example": "eyJhbGciOiJFZDI1NTE5...", "x-sensitive": true}}, "required": ["jti", "operations", "created", "expires"]}
```
