---
title: OutputMcpServerToolItem
page_id: schema-outputmcpservertoolitem-d3498511
path: schemas
description: An openrouter:mcp server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputMcpServerToolItem

An openrouter:mcp server tool output item

```yaml
{"description": "An openrouter:mcp server tool output item", "example": {"id": "mcp_tmp_abc123", "serverLabel": "my-server", "status": "completed", "toolName": "get_data", "type": "openrouter:mcp"}, "properties": {"id": {"type": "string"}, "serverLabel": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "toolName": {"type": "string"}, "type": {"enum": ["openrouter:mcp"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
