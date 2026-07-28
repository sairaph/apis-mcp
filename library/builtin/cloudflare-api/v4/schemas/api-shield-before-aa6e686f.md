---
title: api-shield_before
page_id: schema-api-shield-before-aa6e686f
path: schemas
description: Move rule to after rule with ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_before

Move rule to after rule with ID.

```yaml
{"description": "Move rule to after rule with ID.", "type": "object", "properties": {"before": {"description": "Move rule to before rule with this ID.", "type": "string", "format": "uuid", "example": "0d9bf70c-92e1-4bb3-9411-34a3bcc59003", "maxLength": 36, "x-auditable": true}}}
```
