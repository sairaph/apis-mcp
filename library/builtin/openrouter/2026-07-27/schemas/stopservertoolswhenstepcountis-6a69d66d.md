---
title: StopServerToolsWhenStepCountIs
page_id: schema-stopservertoolswhenstepcountis-6a69d66d
path: schemas
description: Stop after the agent loop has executed this many steps.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhenStepCountIs

Stop after the agent loop has executed this many steps.

```yaml
{"description": "Stop after the agent loop has executed this many steps.", "example": {"step_count": 5, "type": "step_count_is"}, "properties": {"step_count": {"type": "integer"}, "type": {"enum": ["step_count_is"], "type": "string"}}, "required": ["type", "step_count"], "type": "object"}
```
