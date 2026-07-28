---
title: List workspace budgets
page_id: operation-get-workspaces-id-budgets-7bc42c94
path: operations/workspaces
description: List all budgets configured for a workspace. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /workspaces/{id}/budgets
operation_ids:
    - listWorkspaceBudgets
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List workspace budgets

`GET /workspaces/{id}/budgets`

Operation ID: `listWorkspaceBudgets`

List all budgets configured for a workspace. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all budgets configured for a workspace. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listWorkspaceBudgets", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"created_at": "2025-08-24T10:30:00Z", "id": "770e8400-e29b-41d4-a716-446655440000", "limit_usd": 100, "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "include_byok_in_budgets": false}, "schema": {"$ref": "#/components/schemas/ListWorkspaceBudgetsResponse"}}}, "description": "Budgets retrieved successfully"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List workspace budgets", "tags": ["Workspaces"], "x-speakeasy-name-override": "listBudgets"}
```
