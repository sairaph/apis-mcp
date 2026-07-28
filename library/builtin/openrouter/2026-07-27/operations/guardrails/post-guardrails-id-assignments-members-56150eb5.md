---
title: Bulk assign members to a guardrail
page_id: operation-post-guardrails-id-assignments-members-4692c0b6
path: operations/guardrails
description: Assign multiple organization members to a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /guardrails/{id}/assignments/members
operation_ids:
    - bulkAssignMembersToGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Bulk assign members to a guardrail

`POST /guardrails/{id}/assignments/members`

Operation ID: `bulkAssignMembersToGuardrail`

Assign multiple organization members to a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Assign multiple organization members to a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "bulkAssignMembersToGuardrail", "parameters": [{"description": "The unique identifier of the guardrail", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"member_user_ids": ["user_abc123", "user_def456"]}, "schema": {"$ref": "#/components/schemas/BulkAssignMembersRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"assigned_count": 2}, "schema": {"$ref": "#/components/schemas/BulkAssignMembersResponse"}}}, "description": "Assignment result"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Bulk assign members to a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "bulkAssignMembers"}
```
