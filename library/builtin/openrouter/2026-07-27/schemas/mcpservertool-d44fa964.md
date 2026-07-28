---
title: McpServerTool
page_id: schema-mcpservertool-d44fa964
path: schemas
description: MCP (Model Context Protocol) tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# McpServerTool

MCP (Model Context Protocol) tool configuration

```yaml
{"description": "MCP (Model Context Protocol) tool configuration", "example": {"server_label": "my-server", "server_url": "https://example.com/mcp", "type": "mcp"}, "properties": {"allowed_tools": {"anyOf": [{"items": {"type": "string"}, "type": "array"}, {"properties": {"read_only": {"type": "boolean"}, "tool_names": {"items": {"type": "string"}, "type": "array"}}, "type": "object"}, {"type": "null"}]}, "authorization": {"type": "string"}, "connector_id": {"enum": ["connector_dropbox", "connector_gmail", "connector_googlecalendar", "connector_googledrive", "connector_microsoftteams", "connector_outlookcalendar", "connector_outlookemail", "connector_sharepoint"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "headers": {"additionalProperties": {"type": "string"}, "type": ["object", "null"]}, "require_approval": {"anyOf": [{"properties": {"always": {"properties": {"tool_names": {"items": {"type": "string"}, "type": "array"}}, "type": "object"}, "never": {"properties": {"tool_names": {"items": {"type": "string"}, "type": "array"}}, "type": "object"}}, "type": "object"}, {"enum": ["always"], "type": "string"}, {"enum": ["never"], "type": "string"}, {"type": "null"}]}, "server_description": {"type": "string"}, "server_label": {"type": "string"}, "server_url": {"type": "string"}, "type": {"enum": ["mcp"], "type": "string"}}, "required": ["type", "server_label"], "type": "object"}
```
