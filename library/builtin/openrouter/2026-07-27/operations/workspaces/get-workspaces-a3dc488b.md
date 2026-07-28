---
title: List workspaces
page_id: operation-get-workspaces-835dfccd
path: operations/workspaces
description: List all workspaces for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /workspaces
operation_ids:
    - listWorkspaces
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List workspaces

`GET /workspaces`

Operation ID: `listWorkspaces`

List all workspaces for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all workspaces for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listWorkspaces", "parameters": [{"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"created_at": "2025-08-24T10:30:00Z", "created_by": "user_abc123", "default_guardrail_id": "595d5849-7e86-51fd-a7c0-705c34e4afff", "default_image_model": "openai/dall-e-3", "default_provider_sort": "price", "default_text_model": "openai/gpt-4o", "description": "Production environment workspace", "id": "550e8400-e29b-41d4-a716-446655440000", "io_logging_api_key_ids": null, "io_logging_sampling_rate": 1, "is_data_discount_logging_enabled": true, "is_observability_broadcast_enabled": false, "is_observability_io_logging_enabled": false, "name": "Production", "slug": "production", "updated_at": "2025-08-24T15:45:00Z"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListWorkspacesResponse"}}}, "description": "List of workspaces"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List workspaces", "tags": ["Workspaces"], "x-speakeasy-name-override": "list", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
