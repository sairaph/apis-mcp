---
title: zero-trust-gateway_items
page_id: schema-zero-trust-gateway-items-4e23ae82
path: schemas
description: Provide the list items.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_items

Provide the list items.

```yaml
{"description": "Provide the list items.", "type": "array", "items": {"properties": {"created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description_item"}, "value": {"$ref": "#/components/schemas/zero-trust-gateway_value"}}, "type": "object"}, "x-stainless-collection-type": "set"}
```
