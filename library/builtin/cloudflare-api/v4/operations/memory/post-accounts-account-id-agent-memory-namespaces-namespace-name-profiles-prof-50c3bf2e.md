---
title: Ingest messages
page_id: operation-post-accounts-account-id-agent-memory-namespaces-namespace-name-profiles-43334b96
path: operations/memory
description: Processes a conversation and extracts structured memories from it. Agent Memory identifies facts, events, instructions, and tasks automatically.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/agent-memory/namespaces/{namespace_name}/profiles/{profile_name}/ingest
operation_ids:
    - agent-memory-ingest
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Ingest messages

`POST /accounts/{account_id}/agent-memory/namespaces/{namespace_name}/profiles/{profile_name}/ingest`

Operation ID: `agent-memory-ingest`

Processes a conversation and extracts structured memories from it. Agent Memory identifies facts, events, instructions, and tasks automatically.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare Account ID.", "required": true, "schema": {"type": "string", "minLength": 1, "pattern": "^\\d+$"}}]
```

## Definition

```yaml
{"operationId": "agent-memory-ingest", "summary": "Ingest messages", "description": "Processes a conversation and extracts structured memories from it. Agent Memory identifies facts, events, instructions, and tasks automatically.", "parameters": [{"name": "namespace_name", "in": "path", "required": true, "schema": {"description": "Namespace name.", "type": "string", "example": "support-agent", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9-]*[a-z0-9])?$"}}, {"name": "profile_name", "in": "path", "required": true, "schema": {"description": "Profile name.", "type": "string", "example": "my-profile", "maxLength": 100, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9-]*[a-z0-9])?$"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"messages": {"description": "Conversation messages to extract memories from.", "type": "array", "items": {"properties": {"content": {"description": "Text content of the message.", "example": "I really prefer dark mode in the dashboard.", "minLength": 1, "type": "string"}, "role": {"description": "Author role of the message.", "type": "string", "example": "user", "enum": ["system", "user", "assistant"]}, "timestamp": {"description": "Message timestamp.", "type": "string", "example": "2025-09-21T14:30:00Z"}}, "required": ["role", "content"], "type": "object"}, "minItems": 1}, "sessionId": {"description": "Session identifier.", "type": "string", "example": "chat-session-42", "nullable": true}}, "required": ["messages"]}}}}, "responses": {"200": {"description": "Conversation ingested.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "Always empty for a successful response.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"type": "object", "nullable": true}, "success": {"description": "Always true for a successful response.", "enum": [true]}}, "required": ["result", "success", "errors", "messages"]}}}}, "4XX": {"description": "Error response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "One or more errors describing what went wrong.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": [{"code": 1003, "message": "namespace not found"}]}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"description": "Always null for an error response.", "type": "object", "nullable": true}, "success": {"description": "Always false for an error response.", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Memory"]}
```
