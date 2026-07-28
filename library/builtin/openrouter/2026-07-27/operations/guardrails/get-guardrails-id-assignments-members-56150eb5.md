---
title: List member assignments for a guardrail
page_id: operation-get-guardrails-id-assignments-members-37940c31
path: operations/guardrails
description: List all organization member assignments for a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /guardrails/{id}/assignments/members
operation_ids:
    - listGuardrailMemberAssignments
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List member assignments for a guardrail

`GET /guardrails/{id}/assignments/members`

Operation ID: `listGuardrailMemberAssignments`

List all organization member assignments for a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all organization member assignments for a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listGuardrailMemberAssignments", "parameters": [{"description": "The unique identifier of the guardrail", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, {"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "organization_id": "org_xyz789", "user_id": "user_abc123"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListMemberAssignmentsResponse"}}}, "description": "List of member assignments"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List member assignments for a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "listGuardrailMemberAssignments", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
