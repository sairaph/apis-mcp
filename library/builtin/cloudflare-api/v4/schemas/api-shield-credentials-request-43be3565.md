---
title: api-shield_credentials-request
page_id: schema-api-shield-credentials-request-43be3565
path: schemas
description: Request payload for create and PUT credentials operations. Provided keys define the complete stored key set. Key identities (`{alg,kid}`) must be unique.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-request

Request payload for create and PUT credentials operations. Provided keys define the complete stored key set. Key identities (`{alg,kid}`) must be unique.

```yaml
{"description": "Request payload for create and PUT credentials operations. Provided keys define the complete stored key set. Key identities (`{alg,kid}`) must be unique.", "type": "object", "properties": {"keys": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-request"}, "maxItems": 4, "minItems": 1}}, "required": ["keys"]}
```
