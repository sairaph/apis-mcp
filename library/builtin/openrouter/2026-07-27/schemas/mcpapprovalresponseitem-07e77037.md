---
title: McpApprovalResponseItem
page_id: schema-mcpapprovalresponseitem-07e77037
path: schemas
description: User response to an MCP tool approval request
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# McpApprovalResponseItem

User response to an MCP tool approval request

```yaml
{"description": "User response to an MCP tool approval request", "example": {"approval_request_id": "approval-abc123", "approve": true, "reason": "Approved for execution", "type": "mcp_approval_response"}, "properties": {"approval_request_id": {"type": "string"}, "approve": {"type": "boolean"}, "id": {"type": ["string", "null"]}, "reason": {"type": ["string", "null"]}, "type": {"enum": ["mcp_approval_response"], "type": "string"}}, "required": ["type", "approval_request_id", "approve"], "type": "object"}
```
