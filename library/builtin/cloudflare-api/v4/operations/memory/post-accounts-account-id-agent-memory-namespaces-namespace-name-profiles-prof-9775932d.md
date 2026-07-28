---
title: Remember a memory
page_id: operation-post-accounts-account-id-agent-memory-namespaces-namespace-name-profiles-2e00e697
path: operations/memory
description: Stores a single memory explicitly.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/agent-memory/namespaces/{namespace_name}/profiles/{profile_name}/remember
operation_ids:
    - agent-memory-remember
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remember a memory

`POST /accounts/{account_id}/agent-memory/namespaces/{namespace_name}/profiles/{profile_name}/remember`

Operation ID: `agent-memory-remember`

Stores a single memory explicitly.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare Account ID.", "required": true, "schema": {"type": "string", "minLength": 1, "pattern": "^\\d+$"}}]
```

## Definition

```yaml
{"operationId": "agent-memory-remember", "summary": "Remember a memory", "description": "Stores a single memory explicitly.", "parameters": [{"name": "namespace_name", "in": "path", "required": true, "schema": {"description": "Namespace name.", "type": "string", "example": "support-agent", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9-]*[a-z0-9])?$"}}, {"name": "profile_name", "in": "path", "required": true, "schema": {"description": "Profile name.", "type": "string", "example": "my-profile", "maxLength": 100, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9-]*[a-z0-9])?$"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"content": {"description": "Raw memory content to store.", "example": "I prefer dark mode in the dashboard.", "minLength": 1, "type": "string"}, "sessionId": {"description": "Session identifier.", "type": "string", "example": "chat-session-42", "nullable": true}}, "required": ["content"]}}}}, "responses": {"200": {"description": "Memory stored.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "Always empty for a successful response.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"type": "object", "properties": {"content": {"description": "Full memory content extracted from the conversation.", "example": "The user said they prefer dark mode and find the light theme too bright.", "type": "string"}, "createdAt": {"description": "Time the memory was created.", "type": "string", "format": "date-time", "example": "2025-09-21T14:30:00Z"}, "id": {"description": "Unique identifier of the memory.", "type": "string", "example": "01JSGCEXAMPLE000000000000"}, "sessionId": {"description": "Identifier of the session this memory is associated with.", "type": "string", "example": "chat-session-42", "nullable": true}, "summary": {"description": "Short, human-readable summary of the memory.", "type": "string", "example": "User prefers dark mode in the dashboard."}, "type": {"description": "Classification of a memory: 'fact', 'event', 'instruction', or 'task'.", "type": "string", "example": "fact", "enum": ["fact", "event", "instruction", "task"]}, "updatedAt": {"description": "Time the memory was last updated.", "type": "string", "format": "date-time", "example": "2025-09-21T14:30:00Z"}}, "required": ["id", "type", "summary", "content", "sessionId", "createdAt", "updatedAt"]}, "success": {"description": "Always true for a successful response.", "enum": [true]}}, "required": ["result", "success", "errors", "messages"]}}}}, "4XX": {"description": "Error response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "One or more errors describing what went wrong.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": [{"code": 1003, "message": "namespace not found"}]}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"description": "Always null for an error response.", "type": "object", "nullable": true}, "success": {"description": "Always false for an error response.", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Memory"]}
```
