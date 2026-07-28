---
title: Delete a workspace
page_id: operation-delete-workspaces-id-a877cee8
path: operations/workspaces
description: Delete an existing workspace. The default workspace cannot be deleted. Workspaces with active API keys cannot be deleted; remove the keys first. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /workspaces/{id}
operation_ids:
    - deleteWorkspace
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete a workspace

`DELETE /workspaces/{id}`

Operation ID: `deleteWorkspace`

Delete an existing workspace. The default workspace cannot be deleted. Workspaces with active API keys cannot be deleted; remove the keys first. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Delete an existing workspace. The default workspace cannot be deleted. Workspaces with active API keys cannot be deleted; remove the keys first. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "deleteWorkspace", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"deleted": true}, "schema": {"$ref": "#/components/schemas/DeleteWorkspaceResponse"}}}, "description": "Workspace deleted successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete a workspace", "tags": ["Workspaces"], "x-speakeasy-name-override": "delete"}
```
