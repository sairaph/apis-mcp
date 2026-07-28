---
title: zero-trust-gateway_items-input
page_id: schema-zero-trust-gateway-items-input-c3cce64b
path: schemas
description: Add items to the list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_items-input

Add items to the list.

```yaml
{"description": "Add items to the list.", "type": "array", "items": {"properties": {"description": {"$ref": "#/components/schemas/zero-trust-gateway_description_item"}, "value": {"$ref": "#/components/schemas/zero-trust-gateway_value"}}, "type": "object"}, "x-stainless-collection-type": "set"}
```
