---
title: Delete a guardrail
page_id: operation-delete-guardrails-id-b3b1a1b9
path: operations/guardrails
description: Delete an existing guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /guardrails/{id}
operation_ids:
    - deleteGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete a guardrail

`DELETE /guardrails/{id}`

Operation ID: `deleteGuardrail`

Delete an existing guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Delete an existing guardrail. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "deleteGuardrail", "parameters": [{"description": "The unique identifier of the guardrail to delete", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail to delete", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"deleted": true}, "schema": {"$ref": "#/components/schemas/DeleteGuardrailResponse"}}}, "description": "Guardrail deleted successfully"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "delete"}
```
