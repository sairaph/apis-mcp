---
title: api-shield_selector-operation-state
page_id: schema-api-shield-selector-operation-state-839e508b
path: schemas
description: |-
    Details how `selector` interacted with an operation:
      - `included` operations are included by `selector` and will be covered by the Token Validation Rule
      - `excluded` operations are excluded by `selector` and will not be covered by the Token Validation Rule
      - `ignored` operations are not included by `selector` and will not be covered by the Token Validation Rule
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_selector-operation-state

Details how `selector` interacted with an operation:
  - `included` operations are included by `selector` and will be covered by the Token Validation Rule
  - `excluded` operations are excluded by `selector` and will not be covered by the Token Validation Rule
  - `ignored` operations are not included by `selector` and will not be covered by the Token Validation Rule

```yaml
{"description": "Details how `selector` interacted with an operation:\n  - `included` operations are included by `selector` and will be covered by the Token Validation Rule\n  - `excluded` operations are excluded by `selector` and will not be covered by the Token Validation Rule\n  - `ignored` operations are not included by `selector` and will not be covered by the Token Validation Rule\n", "type": "string", "example": "included", "enum": ["included", "excluded", "ignored"], "x-auditable": true}
```
