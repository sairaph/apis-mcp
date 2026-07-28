---
title: StopServerToolsWhenMaxCost
page_id: schema-stopservertoolswhenmaxcost-6a8dce13
path: schemas
description: Stop once cumulative cost across the loop exceeds this dollar threshold.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhenMaxCost

Stop once cumulative cost across the loop exceeds this dollar threshold.

```yaml
{"description": "Stop once cumulative cost across the loop exceeds this dollar threshold.", "example": {"max_cost_in_dollars": 0.5, "type": "max_cost"}, "properties": {"max_cost_in_dollars": {"format": "double", "type": "number"}, "type": {"enum": ["max_cost"], "type": "string"}}, "required": ["type", "max_cost_in_dollars"], "type": "object"}
```
