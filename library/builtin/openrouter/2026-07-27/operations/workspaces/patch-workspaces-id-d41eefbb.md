---
title: Update a workspace
page_id: operation-patch-workspaces-id-799c6941
path: operations/workspaces
description: Update an existing workspace by ID or slug. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /workspaces/{id}
operation_ids:
    - updateWorkspace
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Update a workspace

`PATCH /workspaces/{id}`

Operation ID: `updateWorkspace`

Update an existing workspace by ID or slug. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Update an existing workspace by ID or slug. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "updateWorkspace", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"name": "Updated Workspace", "slug": "updated-workspace"}, "schema": {"$ref": "#/components/schemas/UpdateWorkspaceRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"created_at": "2025-08-24T10:30:00Z", "created_by": "user_abc123", "default_guardrail_id": "595d5849-7e86-51fd-a7c0-705c34e4afff", "default_image_model": "openai/dall-e-3", "default_provider_sort": "price", "default_text_model": "openai/gpt-4o", "description": "Production environment workspace", "id": "550e8400-e29b-41d4-a716-446655440000", "io_logging_api_key_ids": null, "io_logging_sampling_rate": 1, "is_data_discount_logging_enabled": true, "is_observability_broadcast_enabled": false, "is_observability_io_logging_enabled": false, "name": "Updated Workspace", "slug": "updated-workspace", "updated_at": "2025-08-25T10:00:00Z"}}, "schema": {"$ref": "#/components/schemas/UpdateWorkspaceResponse"}}}, "description": "Workspace updated successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Update a workspace", "tags": ["Workspaces"], "x-speakeasy-name-override": "update"}
```
