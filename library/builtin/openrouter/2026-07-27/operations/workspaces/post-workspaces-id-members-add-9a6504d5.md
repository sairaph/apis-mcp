---
title: Bulk add members to a workspace
page_id: operation-post-workspaces-id-members-add-87b24a02
path: operations/workspaces
description: Add multiple organization members to a workspace. Members are assigned the same role they hold in the organization. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /workspaces/{id}/members/add
operation_ids:
    - bulkAddWorkspaceMembers
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Bulk add members to a workspace

`POST /workspaces/{id}/members/add`

Operation ID: `bulkAddWorkspaceMembers`

Add multiple organization members to a workspace. Members are assigned the same role they hold in the organization. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Add multiple organization members to a workspace. Members are assigned the same role they hold in the organization. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "bulkAddWorkspaceMembers", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"user_ids": ["user_abc123", "user_def456"]}, "schema": {"$ref": "#/components/schemas/BulkAddWorkspaceMembersRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"added_count": 1, "data": [{"created_at": "2025-08-24T10:30:00Z", "id": "660e8400-e29b-41d4-a716-446655440000", "role": "member", "user_id": "user_abc123", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}]}, "schema": {"$ref": "#/components/schemas/BulkAddWorkspaceMembersResponse"}}}, "description": "Members added successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Bulk add members to a workspace", "tags": ["Workspaces"], "x-speakeasy-name-override": "bulkAddMembers"}
```
