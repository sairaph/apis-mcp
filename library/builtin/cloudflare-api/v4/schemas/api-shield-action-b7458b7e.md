---
title: api-shield_action
page_id: schema-api-shield-action-b7458b7e
path: schemas
description: Action to take on requests that match operations included in `selector` and fail `expression`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_action

Action to take on requests that match operations included in `selector` and fail `expression`.

```yaml
{"description": "Action to take on requests that match operations included in `selector` and fail `expression`.", "type": "string", "example": "log", "enum": ["log", "block"], "x-auditable": true}
```
