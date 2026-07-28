---
title: api-shield_expression
page_id: schema-api-shield-expression-a9417c5e
path: schemas
description: |-
    Rule expression. Requests that fail to match this expression will be subject to `action`.

    For details on expressions, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_expression

Rule expression. Requests that fail to match this expression will be subject to `action`.

For details on expressions, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).

```yaml
{"description": "Rule expression. Requests that fail to match this expression will be subject to `action`.\n\nFor details on expressions, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).\n", "type": "string", "example": "is_jwt_valid(\"52973293-cb04-4a97-8f55-e7d2ad1107dd\") or is_jwt_valid(\"46eab8d1-6376-45e3-968f-2c649d77d423\")", "x-auditable": true}
```
