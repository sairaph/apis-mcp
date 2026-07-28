---
title: Update a guardrail
page_id: operation-patch-guardrails-id-59dcd55b
path: operations/guardrails
description: 'Update an existing guardrail. Collection fields use replace semantics: send the full desired set on every update. [Management key](/docs/guides/overview/auth/management-api-keys) required.'
source: https://openrouter.ai/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /guardrails/{id}
operation_ids:
    - updateGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Update a guardrail

`PATCH /guardrails/{id}`

Operation ID: `updateGuardrail`

Update an existing guardrail. Collection fields use replace semantics: send the full desired set on every update. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Update an existing guardrail. Collection fields use replace semantics: send the full desired set on every update. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "updateGuardrail", "parameters": [{"description": "The unique identifier of the guardrail to update", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail to update", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"description": "Updated description", "limit_usd": 75, "name": "Updated Guardrail Name", "reset_interval": "weekly"}, "schema": {"$ref": "#/components/schemas/UpdateGuardrailRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"allowed_models": null, "allowed_providers": ["openai"], "created_at": "2025-08-24T10:30:00Z", "description": "Updated description", "enforce_zdr": null, "enforce_zdr_anthropic": true, "enforce_zdr_google": true, "enforce_zdr_openai": true, "enforce_zdr_other": true, "enforce_zdr_xai": true, "id": "550e8400-e29b-41d4-a716-446655440000", "ignored_models": null, "ignored_providers": null, "include_byok_in_budgets": true, "limit_usd": 75, "name": "Updated Guardrail Name", "reset_interval": "weekly", "updated_at": "2025-08-24T16:00:00Z", "workspace_id": "0df9e665-d932-5740-b2c7-b52af166bc11"}}, "schema": {"$ref": "#/components/schemas/UpdateGuardrailResponse"}}}, "description": "Guardrail updated successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Update a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "update"}
```
