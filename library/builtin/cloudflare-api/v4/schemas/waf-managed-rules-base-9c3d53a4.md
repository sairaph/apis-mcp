---
title: waf-managed-rules_base
page_id: schema-waf-managed-rules-base-9c3d53a4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_base

```yaml
{"properties": {"description": {"$ref": "#/components/schemas/waf-managed-rules_schemas-description"}, "group": {"description": "Defines the rule group to which the current WAF rule belongs.", "type": "object", "properties": {"id": {"$ref": "#/components/schemas/waf-managed-rules_components-schemas-identifier"}, "name": {"$ref": "#/components/schemas/waf-managed-rules_name"}}, "readOnly": true}, "id": {"$ref": "#/components/schemas/waf-managed-rules_rule_components-schemas-identifier"}, "package_id": {"$ref": "#/components/schemas/waf-managed-rules_identifier"}, "priority": {"$ref": "#/components/schemas/waf-managed-rules_priority"}}}
```
