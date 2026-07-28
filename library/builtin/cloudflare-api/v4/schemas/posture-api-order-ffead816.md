---
title: posture-api_Order
page_id: schema-posture-api-order-ffead816
path: schemas
description: Generic ordering specification.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_Order

Generic ordering specification.

```yaml
{"description": "Generic ordering specification.", "type": "object", "properties": {"direction": {"$ref": "#/components/schemas/posture-api_DirectionEnum"}, "name": {"description": "Field name to order by.", "type": "string", "example": "instance_count"}}, "required": ["direction", "name"]}
```
