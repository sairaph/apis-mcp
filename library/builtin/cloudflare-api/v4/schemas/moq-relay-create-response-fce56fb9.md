---
title: moq_relay_create_response
page_id: schema-moq-relay-create-response-fce56fb9
path: schemas
description: |-
    Relay with its auto-created default token pair (one full-access
    [publish, subscribe] and one [subscribe]-only), each with its one-time
    secret, wrapped in the issuers envelope.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_relay_create_response

Relay with its auto-created default token pair (one full-access
[publish, subscribe] and one [subscribe]-only), each with its one-time
secret, wrapped in the issuers envelope.

```yaml
{"description": "Relay with its auto-created default token pair (one full-access\n[publish, subscribe] and one [subscribe]-only), each with its one-time\nsecret, wrapped in the issuers envelope.\n", "type": "object", "properties": {"config": {"$ref": "#/components/schemas/moq_relay_config"}, "created": {"type": "string", "format": "date-time"}, "issuers": {"description": "Token collection (discriminated union on `type`). On create this\nholds the auto-created default pair, each including its one-time\nsecret.\n", "type": "array", "items": {"$ref": "#/components/schemas/moq_issuer"}}, "modified": {"type": "string", "format": "date-time"}, "name": {"type": "string", "example": "Production Live Stream"}, "uid": {"description": "Server-generated unique identifier (32 hex chars).", "type": "string", "example": "a1b2c3d4e5f67890a1b2c3d4e5f67890"}}, "required": ["uid", "created", "modified", "name", "issuers", "config"]}
```
