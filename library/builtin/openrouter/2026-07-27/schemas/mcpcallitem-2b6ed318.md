---
title: McpCallItem
page_id: schema-mcpcallitem-2b6ed318
path: schemas
description: An MCP tool call with its output or error
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# McpCallItem

An MCP tool call with its output or error

```yaml
{"description": "An MCP tool call with its output or error", "example": {"arguments": "{\"query\":\"SELECT * FROM users\"}", "id": "mcp-call-abc123", "name": "query_database", "output": "[{\"id\":1,\"name\":\"Alice\"}]", "server_label": "database-server", "type": "mcp_call"}, "properties": {"arguments": {"type": "string"}, "error": {"type": ["string", "null"]}, "id": {"type": "string"}, "name": {"type": "string"}, "output": {"type": ["string", "null"]}, "server_label": {"type": "string"}, "type": {"enum": ["mcp_call"], "type": "string"}}, "required": ["type", "id", "name", "arguments", "server_label"], "type": "object"}
```
