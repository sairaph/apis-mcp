---
title: StopServerToolsWhenCondition
page_id: schema-stopservertoolswhencondition-e387c4f4
path: schemas
description: A single condition that, when met, halts the server-tool agent loop.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StopServerToolsWhenCondition

A single condition that, when met, halts the server-tool agent loop.

```yaml
{"description": "A single condition that, when met, halts the server-tool agent loop.", "discriminator": {"mapping": {"finish_reason_is": "#/components/schemas/StopServerToolsWhenFinishReasonIs", "has_tool_call": "#/components/schemas/StopServerToolsWhenHasToolCall", "max_cost": "#/components/schemas/StopServerToolsWhenMaxCost", "max_tokens_used": "#/components/schemas/StopServerToolsWhenMaxTokensUsed", "step_count_is": "#/components/schemas/StopServerToolsWhenStepCountIs"}, "propertyName": "type"}, "example": {"step_count": 5, "type": "step_count_is"}, "oneOf": [{"$ref": "#/components/schemas/StopServerToolsWhenStepCountIs"}, {"$ref": "#/components/schemas/StopServerToolsWhenHasToolCall"}, {"$ref": "#/components/schemas/StopServerToolsWhenMaxTokensUsed"}, {"$ref": "#/components/schemas/StopServerToolsWhenMaxCost"}, {"$ref": "#/components/schemas/StopServerToolsWhenFinishReasonIs"}]}
```
