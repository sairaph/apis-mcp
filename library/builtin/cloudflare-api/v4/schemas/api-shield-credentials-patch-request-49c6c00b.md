---
title: api-shield_credentials-patch-request
page_id: schema-api-shield-credentials-patch-request-49c6c00b
path: schemas
description: Request payload for PATCH credentials. Provided keys define the complete stored key set, and stored keys omitted from payload are removed. For each provided key identity (`{alg,kid}`), payload fields overwrite the stored key before validation and omitted fields inherit from the stored key. Key identities must be unique.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-patch-request

Request payload for PATCH credentials. Provided keys define the complete stored key set, and stored keys omitted from payload are removed. For each provided key identity (`{alg,kid}`), payload fields overwrite the stored key before validation and omitted fields inherit from the stored key. Key identities must be unique.

```yaml
{"description": "Request payload for PATCH credentials. Provided keys define the complete stored key set, and stored keys omitted from payload are removed. For each provided key identity (`{alg,kid}`), payload fields overwrite the stored key before validation and omitted fields inherit from the stored key. Key identities must be unique.", "type": "object", "properties": {"keys": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-patch-request"}, "maxItems": 4, "minItems": 1}}, "required": ["keys"]}
```
