---
title: waf-managed-rules_allowed_modes_allow_traditional
page_id: schema-waf-managed-rules-allowed-modes-allow-traditional-a8bbeb7d
path: schemas
description: Defines the available modes for the current WAF rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_allowed_modes_allow_traditional

Defines the available modes for the current WAF rule.

```yaml
{"description": "Defines the available modes for the current WAF rule.", "type": "array", "items": {"$ref": "#/components/schemas/waf-managed-rules_mode_allow_traditional"}, "example": ["on", "off"], "readOnly": true}
```
