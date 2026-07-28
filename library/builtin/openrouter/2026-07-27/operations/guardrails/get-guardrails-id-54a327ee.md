---
title: Get a guardrail
page_id: operation-get-guardrails-id-4d581bbd
path: operations/guardrails
description: Get a single guardrail by ID. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /guardrails/{id}
operation_ids:
    - getGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get a guardrail

`GET /guardrails/{id}`

Operation ID: `getGuardrail`

Get a single guardrail by ID. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Get a single guardrail by ID. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "getGuardrail", "parameters": [{"description": "The unique identifier of the guardrail to retrieve", "in": "path", "name": "id", "required": true, "schema": {"description": "The unique identifier of the guardrail to retrieve", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"allowed_models": null, "allowed_providers": ["openai", "anthropic", "google"], "created_at": "2025-08-24T10:30:00Z", "description": "Guardrail for production environment", "enforce_zdr": false, "id": "550e8400-e29b-41d4-a716-446655440000", "ignored_models": null, "ignored_providers": null, "include_byok_in_budgets": false, "limit_usd": 100, "name": "Production Guardrail", "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "0df9e665-d932-5740-b2c7-b52af166bc11"}}, "schema": {"$ref": "#/components/schemas/GetGuardrailResponse"}}}, "description": "Guardrail details"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "get"}
```
