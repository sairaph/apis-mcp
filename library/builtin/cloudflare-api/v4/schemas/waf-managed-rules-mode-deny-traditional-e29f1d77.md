---
title: waf-managed-rules_mode_deny_traditional
page_id: schema-waf-managed-rules-mode-deny-traditional-e29f1d77
path: schemas
description: Defines the action that the current WAF rule will perform when triggered. Applies to traditional (deny) WAF rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_mode_deny_traditional

Defines the action that the current WAF rule will perform when triggered. Applies to traditional (deny) WAF rules.

```yaml
{"description": "Defines the action that the current WAF rule will perform when triggered. Applies to traditional (deny) WAF rules.", "type": "string", "example": "block", "enum": ["default", "disable", "simulate", "block", "challenge"], "x-auditable": true}
```
