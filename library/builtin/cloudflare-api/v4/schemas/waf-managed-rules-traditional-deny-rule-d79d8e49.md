---
title: waf-managed-rules_traditional_deny_rule
page_id: schema-waf-managed-rules-traditional-deny-rule-d79d8e49
path: schemas
description: When triggered, traditional WAF rules cause the firewall to immediately act upon the request based on the configuration of the rule. A 'deny' rule will immediately respond to the request based on the configured rule action/mode (for example, 'block') and no other rules will be processed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_traditional_deny_rule

When triggered, traditional WAF rules cause the firewall to immediately act upon the request based on the configuration of the rule. A 'deny' rule will immediately respond to the request based on the configured rule action/mode (for example, 'block') and no other rules will be processed.

```yaml
{"description": "When triggered, traditional WAF rules cause the firewall to immediately act upon the request based on the configuration of the rule. A 'deny' rule will immediately respond to the request based on the configured rule action/mode (for example, 'block') and no other rules will be processed.", "allOf": [{"$ref": "#/components/schemas/waf-managed-rules_base"}, {"properties": {"allowed_modes": {"$ref": "#/components/schemas/waf-managed-rules_allowed_modes_deny_traditional"}, "default_mode": {"$ref": "#/components/schemas/waf-managed-rules_default_mode"}, "mode": {"$ref": "#/components/schemas/waf-managed-rules_mode_deny_traditional"}}, "type": "object"}], "required": ["id", "description", "priority", "allowed_modes", "default_mode", "mode", "group", "package_id"], "title": "Traditional (deny) WAF rule"}
```
