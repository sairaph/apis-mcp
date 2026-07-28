---
title: Get a namespace
page_id: operation-get-accounts-account-id-agent-memory-namespaces-namespace-name-056d7062
path: operations/namespaces
description: Gets a namespace by name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/agent-memory/namespaces/{namespace_name}
operation_ids:
    - agent-memory-namespace-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a namespace

`GET /accounts/{account_id}/agent-memory/namespaces/{namespace_name}`

Operation ID: `agent-memory-namespace-get`

Gets a namespace by name.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare Account ID.", "required": true, "schema": {"type": "string", "minLength": 1, "pattern": "^\\d+$"}}]
```

## Definition

```yaml
{"operationId": "agent-memory-namespace-get", "summary": "Get a namespace", "description": "Gets a namespace by name.", "parameters": [{"name": "namespace_name", "in": "path", "required": true, "schema": {"description": "Namespace name.", "type": "string", "example": "support-agent", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9-]*[a-z0-9])?$"}}], "responses": {"200": {"description": "The requested namespace.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "Always empty for a successful response.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"type": "object", "properties": {"created_at": {"description": "Time the namespace was created.", "type": "string", "format": "date-time", "example": "2025-09-21T14:30:00.000Z"}, "id": {"description": "Unique identifier of the namespace.", "type": "string", "example": "01JSGCEXAMPLE000000000000"}, "name": {"description": "Namespace name.", "type": "string", "example": "support-agent"}, "updated_at": {"description": "Time the namespace was last updated.", "type": "string", "format": "date-time", "example": "2025-09-21T14:30:00.000Z"}}, "required": ["id", "name", "created_at", "updated_at"]}, "success": {"description": "Always true for a successful response.", "enum": [true]}}, "required": ["result", "success", "errors", "messages"]}}}}, "4XX": {"description": "Error response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "One or more errors describing what went wrong.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": [{"code": 1003, "message": "namespace not found"}]}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"description": "Always null for an error response.", "type": "object", "nullable": true}, "success": {"description": "Always false for an error response.", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Namespaces"]}
```
