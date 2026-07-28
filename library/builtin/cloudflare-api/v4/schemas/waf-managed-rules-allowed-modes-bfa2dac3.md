---
title: waf-managed-rules_allowed_modes
page_id: schema-waf-managed-rules-allowed-modes-bfa2dac3
path: schemas
description: Defines the available states for the rule group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_allowed_modes

Defines the available states for the rule group.

```yaml
{"description": "Defines the available states for the rule group.", "type": "array", "items": {"$ref": "#/components/schemas/waf-managed-rules_mode"}, "example": ["on", "off"], "readOnly": true}
```
