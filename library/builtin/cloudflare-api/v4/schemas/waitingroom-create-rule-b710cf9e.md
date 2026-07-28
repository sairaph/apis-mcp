---
title: waitingroom_create_rule
page_id: schema-waitingroom-create-rule-b710cf9e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_create_rule

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/waitingroom_rule_action"}, "description": {"$ref": "#/components/schemas/waitingroom_rule_description"}, "enabled": {"$ref": "#/components/schemas/waitingroom_rule_enabled"}, "expression": {"$ref": "#/components/schemas/waitingroom_rule_expression"}}, "required": ["action", "expression"]}
```
