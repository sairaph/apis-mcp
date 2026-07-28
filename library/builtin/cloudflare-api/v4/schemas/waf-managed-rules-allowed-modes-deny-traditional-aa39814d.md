---
title: waf-managed-rules_allowed_modes_deny_traditional
page_id: schema-waf-managed-rules-allowed-modes-deny-traditional-aa39814d
path: schemas
description: Defines the list of possible actions of the WAF rule when it is triggered.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_allowed_modes_deny_traditional

Defines the list of possible actions of the WAF rule when it is triggered.

```yaml
{"description": "Defines the list of possible actions of the WAF rule when it is triggered.", "type": "array", "items": {"$ref": "#/components/schemas/waf-managed-rules_mode_deny_traditional"}, "example": ["default", "disable", "simulate", "block", "challenge"], "readOnly": true}
```
