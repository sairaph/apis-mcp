---
title: waitingroom_rule_position
page_id: schema-waitingroom-rule-position-983a445d
path: schemas
description: Reorder the position of a rule
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_rule_position

Reorder the position of a rule

```yaml
{"description": "Reorder the position of a rule", "type": "object", "oneOf": [{"properties": {"index": {"description": "Places the rule in the exact position specified by the integer number <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset from the specified position number onward are shifted one position (no rule is overwritten).", "type": "integer", "x-auditable": true}}, "type": "object"}, {"properties": {"before": {"description": "Places the rule before rule <RULE_ID>. Use this argument with an empty rule ID value (\"\") to set the rule as the first rule in the ruleset.", "type": "string", "example": "<RULE_ID>", "x-auditable": true}}, "type": "object"}, {"properties": {"after": {"description": "Places the rule after rule <RULE_ID>. Use this argument with an empty rule ID value (\"\") to set the rule as the last rule in the ruleset.", "type": "string", "example": "<RULE_ID>", "x-auditable": true}}, "type": "object"}]}
```
