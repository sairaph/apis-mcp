---
title: OutputMemoryServerToolItem
page_id: schema-outputmemoryservertoolitem-401ca078
path: schemas
description: An openrouter:memory server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputMemoryServerToolItem

An openrouter:memory server tool output item

```yaml
{"description": "An openrouter:memory server tool output item", "example": {"action": "read", "id": "mem_tmp_abc123", "key": "user_preference", "status": "completed", "type": "openrouter:memory"}, "properties": {"action": {"enum": ["read", "write", "delete"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "id": {"type": "string"}, "key": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:memory"], "type": "string"}, "value": {}}, "required": ["status", "type"], "type": "object"}
```
