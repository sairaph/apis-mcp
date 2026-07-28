---
title: NamespaceTool
page_id: schema-namespacetool-119ae06d
path: schemas
description: Groups function/custom tools under a shared namespace
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# NamespaceTool

Groups function/custom tools under a shared namespace

```yaml
{"description": "Groups function/custom tools under a shared namespace", "example": {"description": "Tools for spawning and managing sub-agents.", "name": "multi_agent_v1", "tools": [{"name": "spawn_agent", "type": "function"}], "type": "namespace"}, "properties": {"description": {"type": "string"}, "name": {"type": "string"}, "tools": {"items": {"anyOf": [{"$ref": "#/components/schemas/NamespaceFunctionTool"}, {"$ref": "#/components/schemas/CustomTool"}]}, "type": "array"}, "type": {"enum": ["namespace"], "type": "string"}}, "required": ["type", "name", "description", "tools"], "type": "object"}
```
