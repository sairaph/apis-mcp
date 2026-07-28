---
title: zero-trust-gateway_class
page_id: schema-zero-trust-gateway-class-ebb79381
path: schemas
description: Specify which account types can create policies for this category. `blocked` Blocks unconditionally for all accounts. `removalPending` Allows removal from policies but disables addition. `noBlock` Prevents blocking.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_class

Specify which account types can create policies for this category. `blocked` Blocks unconditionally for all accounts. `removalPending` Allows removal from policies but disables addition. `noBlock` Prevents blocking.

```yaml
{"description": "Specify which account types can create policies for this category. `blocked` Blocks unconditionally for all accounts. `removalPending` Allows removal from policies but disables addition. `noBlock` Prevents blocking.", "type": "string", "example": "premium", "enum": ["free", "premium", "blocked", "removalPending", "noBlock"], "x-auditable": true}
```
