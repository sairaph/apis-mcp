---
title: one_Permission
page_id: schema-one-permission-d1d59cdf
path: schemas
description: Permission/scope with severity for display.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_Permission

Permission/scope with severity for display.

```yaml
{"description": "Permission/scope with severity for display.", "type": "object", "properties": {"display_name": {"description": "Human-readable permission name.", "type": "string"}, "scope": {"description": "Vendor-native scope identifier.", "type": "string"}, "severity": {"description": "Permission sensitivity level.\n\n* `low` - low\n* `medium` - medium\n* `high` - high\n* `critical` - critical", "type": "string", "enum": ["low", "medium", "high", "critical"]}}, "required": ["display_name", "scope", "severity"]}
```
