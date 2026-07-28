---
title: waf-managed-rules_traditional_allow_rule
page_id: schema-waf-managed-rules-traditional-allow-rule-1d574594
path: schemas
description: When triggered, traditional WAF rules cause the firewall to immediately act on the request based on the rule configuration. An 'allow' rule will immediately allow the request and no other rules will be processed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_traditional_allow_rule

When triggered, traditional WAF rules cause the firewall to immediately act on the request based on the rule configuration. An 'allow' rule will immediately allow the request and no other rules will be processed.

```yaml
{"description": "When triggered, traditional WAF rules cause the firewall to immediately act on the request based on the rule configuration. An 'allow' rule will immediately allow the request and no other rules will be processed.", "allOf": [{"$ref": "#/components/schemas/waf-managed-rules_base"}, {"properties": {"allowed_modes": {"$ref": "#/components/schemas/waf-managed-rules_allowed_modes_allow_traditional"}, "mode": {"$ref": "#/components/schemas/waf-managed-rules_mode_allow_traditional"}}, "type": "object"}], "required": ["id", "description", "priority", "allowed_modes", "default_mode", "mode", "group", "package_id"], "title": "Traditional (allow) WAF rule"}
```
