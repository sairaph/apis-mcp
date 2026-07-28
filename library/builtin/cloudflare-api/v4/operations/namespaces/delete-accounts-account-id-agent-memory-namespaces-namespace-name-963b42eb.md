---
title: Delete a namespace
page_id: operation-delete-accounts-account-id-agent-memory-namespaces-namespace-name-9ff29df9
path: operations/namespaces
description: Deletes a namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/agent-memory/namespaces/{namespace_name}
operation_ids:
    - agent-memory-namespace-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a namespace

`DELETE /accounts/{account_id}/agent-memory/namespaces/{namespace_name}`

Operation ID: `agent-memory-namespace-delete`

Deletes a namespace.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Cloudflare Account ID.", "required": true, "schema": {"type": "string", "minLength": 1, "pattern": "^\\d+$"}}]
```

## Definition

```yaml
{"operationId": "agent-memory-namespace-delete", "summary": "Delete a namespace", "description": "Deletes a namespace.", "parameters": [{"name": "namespace_name", "in": "path", "required": true, "schema": {"description": "Namespace name.", "type": "string", "example": "support-agent", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9]([a-z0-9-]*[a-z0-9])?$"}}], "responses": {"200": {"description": "Namespace deleted.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "Always empty for a successful response.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"type": "object", "nullable": true}, "success": {"description": "Always true for a successful response.", "enum": [true]}}, "required": ["result", "success", "errors", "messages"]}}}}, "4XX": {"description": "Error response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"description": "One or more errors describing what went wrong.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": [{"code": 1003, "message": "namespace not found"}]}, "messages": {"description": "Informational, non-error messages, if any.", "type": "array", "items": {"properties": {"code": {"description": "Machine-readable status code.", "type": "number", "example": 1003}, "documentation_url": {"description": "Link to relevant documentation, when available.", "type": "string"}, "message": {"description": "Human-readable message describing the error or status.", "type": "string", "example": "namespace not found"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the request field that caused the error.", "type": "string", "example": "/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": []}, "result": {"description": "Always null for an error response.", "type": "object", "nullable": true}, "success": {"description": "Always false for an error response.", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Namespaces"]}
```
