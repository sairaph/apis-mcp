---
title: bill-subs-api_component_value
page_id: schema-bill-subs-api-component-value-d00f94b8
path: schemas
description: A component value for a subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bill-subs-api_component_value

A component value for a subscription.

```yaml
{"description": "A component value for a subscription.", "type": "object", "properties": {"default": {"description": "The default amount assigned.", "type": "number", "example": 5, "x-auditable": true}, "display_name": {"description": "A human-readable version of the component name.", "type": "string", "example": "Page Rules"}, "kind": {"description": "The type of component value. \"enum\" for discrete values (including boolean on/off toggles where 0=off and 1=on), \"sum\" for countable quantities, \"usage\" for metered billing components.", "type": "string", "example": "sum", "enum": ["enum", "sum", "usage"]}, "name": {"description": "The name of the component value.", "type": "string", "example": "page_rules", "x-auditable": true}, "price": {"description": "The unit price for the component value.", "type": "number", "example": 5, "x-auditable": true}, "value": {"description": "The amount of the component value assigned.", "type": "number", "example": 20, "x-auditable": true}}}
```
