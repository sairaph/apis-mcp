---
title: firewall_rewrite_action
page_id: schema-firewall-rewrite-action-49fef4f1
path: schemas
description: Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_rewrite_action

Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object.

```yaml
{"description": "Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object.", "type": "object", "properties": {"block": {"$ref": "#/components/schemas/firewall_waf_rewrite_action"}, "challenge": {"$ref": "#/components/schemas/firewall_waf_rewrite_action"}, "default": {"$ref": "#/components/schemas/firewall_waf_rewrite_action"}, "disable": {"$ref": "#/components/schemas/firewall_waf_rewrite_action"}, "simulate": {"$ref": "#/components/schemas/firewall_waf_rewrite_action"}}}
```
