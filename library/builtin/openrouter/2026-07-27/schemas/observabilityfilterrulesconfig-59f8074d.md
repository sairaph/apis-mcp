---
title: ObservabilityFilterRulesConfig
page_id: schema-observabilityfilterrulesconfig-59f8074d
path: schemas
description: Optional structured filter rules controlling which events are forwarded.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ObservabilityFilterRulesConfig

Optional structured filter rules controlling which events are forwarded.

```yaml
{"description": "Optional structured filter rules controlling which events are forwarded.", "example": null, "properties": {"enabled": {"default": true, "type": "boolean"}, "groups": {"items": {"$ref": "#/components/schemas/ObservabilityFilterRuleGroup"}, "type": "array"}}, "required": ["groups"], "type": ["object", "null"]}
```
