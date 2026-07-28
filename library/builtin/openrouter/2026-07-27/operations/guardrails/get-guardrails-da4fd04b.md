---
title: List guardrails
page_id: operation-get-guardrails-be567751
path: operations/guardrails
description: List all guardrails for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /guardrails
operation_ids:
    - listGuardrails
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List guardrails

`GET /guardrails`

Operation ID: `listGuardrails`

List all guardrails for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all guardrails for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listGuardrails", "parameters": [{"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}, {"description": "Filter guardrails by workspace ID. By default, guardrails in the default workspace are returned.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Filter guardrails by workspace ID. By default, guardrails in the default workspace are returned.", "example": "0df9e665-d932-5740-b2c7-b52af166bc11", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"allowed_models": null, "allowed_providers": ["openai", "anthropic", "google"], "created_at": "2025-08-24T10:30:00Z", "description": "Guardrail for production environment", "enforce_zdr": false, "id": "550e8400-e29b-41d4-a716-446655440000", "ignored_models": null, "ignored_providers": null, "include_byok_in_budgets": false, "limit_usd": 100, "name": "Production Guardrail", "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "0df9e665-d932-5740-b2c7-b52af166bc11"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListGuardrailsResponse"}}}, "description": "List of guardrails"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List guardrails", "tags": ["Guardrails"], "x-speakeasy-name-override": "list", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
