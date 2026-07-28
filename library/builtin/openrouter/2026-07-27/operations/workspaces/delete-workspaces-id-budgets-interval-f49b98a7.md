---
title: Delete a workspace budget
page_id: operation-delete-workspaces-id-budgets-interval-3d3f6875
path: operations/workspaces
description: Remove the budget for a given interval. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /workspaces/{id}/budgets/{interval}
operation_ids:
    - deleteWorkspaceBudget
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete a workspace budget

`DELETE /workspaces/{id}/budgets/{interval}`

Operation ID: `deleteWorkspaceBudget`

Remove the budget for a given interval. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Remove the budget for a given interval. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "deleteWorkspaceBudget", "parameters": [{"description": "The workspace ID (UUID) or slug", "in": "path", "name": "id", "required": true, "schema": {"description": "The workspace ID (UUID) or slug", "example": "production", "minLength": 1, "type": "string"}}, {"description": "Budget reset interval. Use \"lifetime\" for a one-time budget that never resets.", "example": "monthly", "in": "path", "name": "interval", "required": true, "schema": {"$ref": "#/components/schemas/WorkspaceBudgetInterval"}}], "responses": {"200": {"content": {"application/json": {"example": {"deleted": true}, "schema": {"$ref": "#/components/schemas/DeleteWorkspaceBudgetResponse"}}}, "description": "Budget deleted successfully"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete a workspace budget", "tags": ["Workspaces"], "x-speakeasy-name-override": "deleteBudget"}
```
