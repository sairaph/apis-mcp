---
title: Bulk unassign members from a guardrail
page_id: operation-post-guardrails-id-assignments-members-remove-f00af361
path: operations/guardrails
description: Unassign multiple organization members from a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /guardrails/{id}/assignments/members/remove
operation_ids:
    - bulkUnassignMembersFromGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Bulk unassign members from a guardrail

`POST /guardrails/{id}/assignments/members/remove`

Operation ID: `bulkUnassignMembersFromGuardrail`

Unassign multiple organization members from a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Unassign multiple organization members from a specific guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "bulkUnassignMembersFromGuardrail", "parameters": [{"description": "The unique identifier of the guardrail", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"member_user_ids": ["user_abc123", "user_def456"]}, "schema": {"$ref": "#/components/schemas/BulkUnassignMembersRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"unassigned_count": 2}, "schema": {"$ref": "#/components/schemas/BulkUnassignMembersResponse"}}}, "description": "Unassignment result"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Bulk unassign members from a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "bulkUnassignMembers"}
```
