---
title: OutputSubagentServerToolItem
page_id: schema-outputsubagentservertoolitem-a5b701f6
path: schemas
description: An openrouter:subagent server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputSubagentServerToolItem

An openrouter:subagent server tool output item

```yaml
{"description": "An openrouter:subagent server tool output item", "example": {"id": "st_tmp_abc123", "status": "completed", "type": "openrouter:subagent"}, "properties": {"error": {"description": "Error message when the subagent task did not produce an outcome.", "type": "string"}, "id": {"type": "string"}, "instance_name": {"description": "Provider-safe function name of the specific subagent instance that produced this item (e.g. `openrouter_subagent__1`). Present only on items from non-default instances — the second and later subagent entries in the request `tools` array. The first (default) instance omits it, even when multiple subagents are configured. When a replayed item echoes this field back, the transcript rehydrates the call under that instance's tool. This identity is positional: it is derived from the index of the subagent entry in the request `tools` array, so keep the order of subagent entries stable across requests in a conversation.", "example": "openrouter_subagent__1", "type": "string"}, "model": {"description": "Slug of the worker model that executed the task.", "type": "string"}, "name": {"description": "Configured name of the subagent that executed the task (the `name` on its tool entry). Present only for named subagents; omitted for an unnamed (default) subagent.", "example": "summarizer", "type": "string"}, "outcome": {"description": "The worker model's result (the outcome text returned to the delegating model).", "type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "task_description": {"description": "The task description the delegating model sent to the worker.", "type": "string"}, "task_name": {"description": "The short task identifier the delegating model supplied.", "type": "string"}, "type": {"enum": ["openrouter:subagent"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
