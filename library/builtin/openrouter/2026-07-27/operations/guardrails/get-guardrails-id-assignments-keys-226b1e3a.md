---
title: List key assignments for a guardrail
page_id: operation-get-guardrails-id-assignments-keys-cc7a4a6a
path: operations/guardrails
description: List all API key assignments for a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /guardrails/{id}/assignments/keys
operation_ids:
    - listGuardrailKeyAssignments
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List key assignments for a guardrail

`GET /guardrails/{id}/assignments/keys`

Operation ID: `listGuardrailKeyAssignments`

List all API key assignments for a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all API key assignments for a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listGuardrailKeyAssignments", "parameters": [{"description": "The unique identifier of the guardrail", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, {"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "key_hash": "c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93", "key_label": "prod-key", "key_name": "Production Key"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListKeyAssignmentsResponse"}}}, "description": "List of key assignments"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List key assignments for a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "listGuardrailKeyAssignments", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
