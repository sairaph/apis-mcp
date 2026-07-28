---
title: OutputBrowserUseServerToolItem
page_id: schema-outputbrowseruseservertoolitem-651dee43
path: schemas
description: An openrouter:browser_use server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputBrowserUseServerToolItem

An openrouter:browser_use server tool output item

```yaml
{"description": "An openrouter:browser_use server tool output item", "example": {"action": "screenshot", "id": "bu_tmp_abc123", "status": "completed", "type": "openrouter:browser_use"}, "properties": {"action": {"type": "string"}, "id": {"type": "string"}, "screenshotB64": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:browser_use"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
