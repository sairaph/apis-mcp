---
title: McpApprovalRequestItem
page_id: schema-mcpapprovalrequestitem-42115be5
path: schemas
description: Request for approval to execute an MCP tool
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# McpApprovalRequestItem

Request for approval to execute an MCP tool

```yaml
{"description": "Request for approval to execute an MCP tool", "example": {"arguments": "{\"id\":\"123\"}", "id": "approval-abc123", "name": "delete_record", "server_label": "database-server", "type": "mcp_approval_request"}, "properties": {"arguments": {"type": "string"}, "id": {"type": "string"}, "name": {"type": "string"}, "server_label": {"type": "string"}, "type": {"enum": ["mcp_approval_request"], "type": "string"}}, "required": ["type", "id", "name", "arguments", "server_label"], "type": "object"}
```
