---
title: ObservabilityFilterRulesConfigNullable
page_id: schema-observabilityfilterrulesconfignullable-0eb27a38
path: schemas
description: Optional structured filter rules controlling which events are forwarded.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ObservabilityFilterRulesConfigNullable

Optional structured filter rules controlling which events are forwarded.

```yaml
{"description": "Optional structured filter rules controlling which events are forwarded.", "example": null, "properties": {"enabled": {"default": true, "type": "boolean"}, "groups": {"items": {"$ref": "#/components/schemas/ObservabilityFilterRuleGroup"}, "type": "array"}}, "required": ["groups"], "type": ["object", "null"]}
```
