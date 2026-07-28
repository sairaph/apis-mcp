---
title: StopServerToolsWhen
page_id: schema-stopservertoolswhen-ccfea0d9
path: schemas
description: Stop conditions for the server-tool agent loop. Any condition firing halts the loop (OR logic). When set, this overrides `max_tool_calls`. When a condition fires while the model is still emitting tool calls, the pending tool calls are executed and one final turn is made with tool calls disabled so the response ends with a natural-language answer instead of an unfinished tool call.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhen

Stop conditions for the server-tool agent loop. Any condition firing halts the loop (OR logic). When set, this overrides `max_tool_calls`. When a condition fires while the model is still emitting tool calls, the pending tool calls are executed and one final turn is made with tool calls disabled so the response ends with a natural-language answer instead of an unfinished tool call.

```yaml
{"description": "Stop conditions for the server-tool agent loop. Any condition firing halts the loop (OR logic). When set, this overrides `max_tool_calls`. When a condition fires while the model is still emitting tool calls, the pending tool calls are executed and one final turn is made with tool calls disabled so the response ends with a natural-language answer instead of an unfinished tool call.", "example": [{"step_count": 5, "type": "step_count_is"}, {"max_cost_in_dollars": 0.5, "type": "max_cost"}], "items": {"$ref": "#/components/schemas/StopServerToolsWhenCondition"}, "minItems": 1, "type": "array"}
```
