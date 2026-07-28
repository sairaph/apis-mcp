---
title: api-shield_position
page_id: schema-api-shield-position-a869757d
path: schemas
description: Update rule order among zone rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_position

Update rule order among zone rules.

```yaml
{"description": "Update rule order among zone rules.", "type": "object", "oneOf": [{"$ref": "#/components/schemas/api-shield_index"}, {"$ref": "#/components/schemas/api-shield_before"}, {"$ref": "#/components/schemas/api-shield_after"}], "writeOnly": true}
```
