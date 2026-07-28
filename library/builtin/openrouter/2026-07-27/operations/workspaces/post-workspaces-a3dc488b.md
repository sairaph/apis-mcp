---
title: Create a workspace
page_id: operation-post-workspaces-d780ceb0
path: operations/workspaces
description: Create a new workspace for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /workspaces
operation_ids:
    - createWorkspace
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create a workspace

`POST /workspaces`

Operation ID: `createWorkspace`

Create a new workspace for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Create a new workspace for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "createWorkspace", "requestBody": {"content": {"application/json": {"example": {"default_image_model": "openai/dall-e-3", "default_provider_sort": "price", "default_text_model": "openai/gpt-4o", "description": "Production environment workspace", "name": "Production", "slug": "production"}, "schema": {"$ref": "#/components/schemas/CreateWorkspaceRequest"}}}, "required": true}, "responses": {"201": {"content": {"application/json": {"example": {"data": {"created_at": "2025-08-24T10:30:00Z", "created_by": "user_abc123", "default_guardrail_id": "595d5849-7e86-51fd-a7c0-705c34e4afff", "default_image_model": "openai/dall-e-3", "default_provider_sort": "price", "default_text_model": "openai/gpt-4o", "description": "Production environment workspace", "id": "550e8400-e29b-41d4-a716-446655440000", "io_logging_api_key_ids": null, "io_logging_sampling_rate": 1, "is_data_discount_logging_enabled": true, "is_observability_broadcast_enabled": false, "is_observability_io_logging_enabled": false, "name": "Production", "slug": "production", "updated_at": null}}, "schema": {"$ref": "#/components/schemas/CreateWorkspaceResponse"}}}, "description": "Workspace created successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Create a workspace", "tags": ["Workspaces"], "x-speakeasy-name-override": "create"}
```
