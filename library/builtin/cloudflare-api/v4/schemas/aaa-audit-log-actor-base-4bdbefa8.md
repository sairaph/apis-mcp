---
title: aaa_audit-log-actor-base
page_id: schema-aaa-audit-log-actor-base-4bdbefa8
path: schemas
description: Provides details about the actor who performed the action.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-log-actor-base

Provides details about the actor who performed the action.

```yaml
{"description": "Provides details about the actor who performed the action.", "type": "object", "properties": {"context": {"type": "string", "example": "dash", "enum": ["api_key", "api_token", "dash", "oauth", "origin_ca_key"]}, "email": {"description": "The email of the actor who performed the action.", "type": "string", "format": "email", "example": "alice@example.com"}, "id": {"description": "The ID of the actor who performed the action. If a user performed the action, this will be their User ID.", "type": "string", "example": "f6b5de0326bb5182b8a4840ee01ec774"}, "ip_address": {"description": "The IP address of the request that performed the action.", "type": "string", "example": "198.41.129.166"}, "token_id": {"description": "The API token ID when the actor context is an api_token or oauth.", "type": "string"}, "token_name": {"description": "The API token name when the actor context is an api_token or oauth.", "type": "string"}}}
```
