---
title: access_target_attributes
page_id: schema-access-target-attributes-010969e5
path: schemas
description: Contains a map of target attribute keys to target attribute values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_target_attributes

Contains a map of target attribute keys to target attribute values.

```yaml
{"description": "Contains a map of target attribute keys to target attribute values.", "type": "object", "example": {"hostname": ["test-server", "production-server"]}, "additionalProperties": {"example": ["test-server", "production-server"], "items": {"type": "string"}, "type": "array"}}
```
