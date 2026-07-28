---
title: Bulk assign keys to a guardrail
page_id: operation-post-guardrails-id-assignments-keys-f6a03aa3
path: operations/guardrails
description: Assign multiple API keys to a specific guardrail. A key may hold at most one guardrail; assigning replaces any existing assignment. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /guardrails/{id}/assignments/keys
operation_ids:
    - bulkAssignKeysToGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Bulk assign keys to a guardrail

`POST /guardrails/{id}/assignments/keys`

Operation ID: `bulkAssignKeysToGuardrail`

Assign multiple API keys to a specific guardrail. A key may hold at most one guardrail; assigning replaces any existing assignment. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Assign multiple API keys to a specific guardrail. A key may hold at most one guardrail; assigning replaces any existing assignment. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "bulkAssignKeysToGuardrail", "parameters": [{"description": "The unique identifier of the guardrail", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"key_hashes": ["c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93"]}, "schema": {"$ref": "#/components/schemas/BulkAssignKeysRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"assigned_count": 3}, "schema": {"$ref": "#/components/schemas/BulkAssignKeysResponse"}}}, "description": "Assignment result"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Bulk assign keys to a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "bulkAssignKeys"}
```
