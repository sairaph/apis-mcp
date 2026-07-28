---
title: List workspace members
page_id: operation-get-workspaces-id-members-8e3e58cf
path: operations/workspaces
description: List all members of a workspace. Returns paginated results. For the default workspace, returns all organization members (implicit membership). [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /workspaces/{id}/members
operation_ids:
    - listWorkspaceMembers
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List workspace members

`GET /workspaces/{id}/members`

Operation ID: `listWorkspaceMembers`

List all members of a workspace. Returns paginated results. For the default workspace, returns all organization members (implicit membership). [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all members of a workspace. Returns paginated results. For the default workspace, returns all organization members (implicit membership). [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listWorkspaceMembers", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}, {"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"created_at": "2025-08-24T10:30:00Z", "id": "660e8400-e29b-41d4-a716-446655440000", "role": "member", "user_id": "user_abc123", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListWorkspaceMembersResponse"}}}, "description": "List of workspace members"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List workspace members", "tags": ["Workspaces"], "x-speakeasy-name-override": "listMembers", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
