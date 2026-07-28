---
title: Create a guardrail
page_id: operation-post-guardrails-8c300ed1
path: operations/guardrails
description: Create a new guardrail for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /guardrails
operation_ids:
    - createGuardrail
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create a guardrail

`POST /guardrails`

Operation ID: `createGuardrail`

Create a new guardrail for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Create a new guardrail for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "createGuardrail", "requestBody": {"content": {"application/json": {"example": {"allowed_models": null, "allowed_providers": ["openai", "anthropic", "deepseek"], "description": "A guardrail for limiting API usage", "enforce_zdr_anthropic": true, "enforce_zdr_google": false, "enforce_zdr_openai": true, "enforce_zdr_other": false, "enforce_zdr_xai": false, "ignored_models": null, "ignored_providers": null, "limit_usd": 50, "name": "My New Guardrail", "reset_interval": "monthly"}, "schema": {"$ref": "#/components/schemas/CreateGuardrailRequest"}}}, "required": true}, "responses": {"201": {"content": {"application/json": {"example": {"data": {"allowed_models": null, "allowed_providers": ["openai", "anthropic", "google"], "created_at": "2025-08-24T10:30:00Z", "description": "A guardrail for limiting API usage", "enforce_zdr": null, "enforce_zdr_anthropic": true, "enforce_zdr_google": false, "enforce_zdr_openai": true, "enforce_zdr_other": false, "enforce_zdr_xai": false, "id": "550e8400-e29b-41d4-a716-446655440000", "ignored_models": null, "ignored_providers": null, "include_byok_in_budgets": false, "limit_usd": 50, "name": "My New Guardrail", "reset_interval": "monthly", "updated_at": null, "workspace_id": "0df9e665-d932-5740-b2c7-b52af166bc11"}}, "schema": {"$ref": "#/components/schemas/CreateGuardrailResponse"}}}, "description": "Guardrail created successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Create a guardrail", "tags": ["Guardrails"], "x-speakeasy-name-override": "create"}
```
