---
title: Create or update a workspace budget
page_id: operation-put-workspaces-id-budgets-interval-6d55e7d9
path: operations/workspaces
description: Create or update the budget for a given interval. Budget limits must strictly decrease as the interval narrows (lifetime > monthly > weekly > daily). [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /workspaces/{id}/budgets/{interval}
operation_ids:
    - upsertWorkspaceBudget
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create or update a workspace budget

`PUT /workspaces/{id}/budgets/{interval}`

Operation ID: `upsertWorkspaceBudget`

Create or update the budget for a given interval. Budget limits must strictly decrease as the interval narrows (lifetime > monthly > weekly > daily). [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Create or update the budget for a given interval. Budget limits must strictly decrease as the interval narrows (lifetime > monthly > weekly > daily). [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "upsertWorkspaceBudget", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}, {"description": "Budget reset interval. Use \"lifetime\" for a one-time budget that never resets.", "example": "monthly", "in": "path", "name": "interval", "required": true, "schema": {"$ref": "#/components/schemas/WorkspaceBudgetInterval"}}], "requestBody": {"content": {"application/json": {"example": {"include_byok_in_budgets": true, "limit_usd": 100}, "schema": {"$ref": "#/components/schemas/UpsertWorkspaceBudgetRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"created_at": "2025-08-24T10:30:00Z", "id": "770e8400-e29b-41d4-a716-446655440000", "limit_usd": 100, "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}, "include_byok_in_budgets": true}, "schema": {"$ref": "#/components/schemas/UpsertWorkspaceBudgetResponse"}}}, "description": "Budget created or updated successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Create or update a workspace budget", "tags": ["Workspaces"], "x-speakeasy-name-override": "setBudget"}
```
