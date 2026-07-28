---
title: McpListToolsItem
page_id: schema-mcplisttoolsitem-14879a60
path: schemas
description: List of available MCP tools from a server
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# McpListToolsItem

List of available MCP tools from a server

```yaml
{"description": "List of available MCP tools from a server", "example": {"id": "mcp-list-abc123", "server_label": "database-server", "tools": [{"description": "Execute a database query", "input_schema": {"properties": {"query": {"type": "string"}}, "type": "object"}, "name": "query_database"}], "type": "mcp_list_tools"}, "properties": {"error": {"type": ["string", "null"]}, "id": {"type": "string"}, "server_label": {"type": "string"}, "tools": {"items": {"properties": {"annotations": {}, "description": {"type": ["string", "null"]}, "input_schema": {"additionalProperties": {}, "type": "object"}, "name": {"type": "string"}}, "required": ["name", "input_schema"], "type": "object"}, "type": "array"}, "type": {"enum": ["mcp_list_tools"], "type": "string"}}, "required": ["type", "id", "server_label", "tools"], "type": "object"}
```
