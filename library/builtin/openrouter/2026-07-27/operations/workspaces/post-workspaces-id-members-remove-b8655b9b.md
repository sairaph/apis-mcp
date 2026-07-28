---
title: Bulk remove members from a workspace
page_id: operation-post-workspaces-id-members-remove-eae81775
path: operations/workspaces
description: Remove multiple members from a workspace. Members with active API keys in the workspace cannot be removed. SCIM-managed members cannot be removed; changes must be made in your identity provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /workspaces/{id}/members/remove
operation_ids:
    - bulkRemoveWorkspaceMembers
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Bulk remove members from a workspace

`POST /workspaces/{id}/members/remove`

Operation ID: `bulkRemoveWorkspaceMembers`

Remove multiple members from a workspace. Members with active API keys in the workspace cannot be removed. SCIM-managed members cannot be removed; changes must be made in your identity provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Remove multiple members from a workspace. Members with active API keys in the workspace cannot be removed. SCIM-managed members cannot be removed; changes must be made in your identity provider. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "bulkRemoveWorkspaceMembers", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"user_ids": ["user_abc123", "user_def456"]}, "schema": {"$ref": "#/components/schemas/BulkRemoveWorkspaceMembersRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"removed_count": 2}, "schema": {"$ref": "#/components/schemas/BulkRemoveWorkspaceMembersResponse"}}}, "description": "Members removed successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Bulk remove members from a workspace", "tags": ["Workspaces"], "x-speakeasy-name-override": "bulkRemoveMembers"}
```
