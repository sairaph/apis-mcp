---
title: Get user activity grouped by endpoint
page_id: operation-get-activity-a8919950
path: operations/analytics
description: Returns user activity data grouped by endpoint for the last 30 (completed) UTC days. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /activity
operation_ids:
    - getUserActivity
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get user activity grouped by endpoint

`GET /activity`

Operation ID: `getUserActivity`

Returns user activity data grouped by endpoint for the last 30 (completed) UTC days. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns user activity data grouped by endpoint for the last 30 (completed) UTC days. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "getUserActivity", "parameters": [{"description": "Filter by a single UTC date in the last 30 days (YYYY-MM-DD format).", "in": "query", "name": "date", "required": false, "schema": {"description": "Filter by a single UTC date in the last 30 days (YYYY-MM-DD format).", "example": "2025-08-24", "type": "string"}}, {"description": "Filter by API key hash (SHA-256 hex string, as returned by the keys API).", "in": "query", "name": "api_key_hash", "required": false, "schema": {"description": "Filter by API key hash (SHA-256 hex string, as returned by the keys API).", "example": "abc123def456...", "type": "string"}}, {"description": "Filter by org member user ID. Only applicable for organization accounts.", "in": "query", "name": "user_id", "required": false, "schema": {"description": "Filter by org member user ID. Only applicable for organization accounts.", "example": "user_abc123", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"byok_usage_inference": 0.012, "completion_tokens": 125, "date": "2025-08-24", "endpoint_id": "550e8400-e29b-41d4-a716-446655440000", "model": "openai/gpt-4.1", "model_permaslug": "openai/gpt-4.1-2025-04-14", "prompt_tokens": 50, "provider_name": "OpenAI", "reasoning_tokens": 25, "requests": 5, "usage": 0.015}]}, "schema": {"$ref": "#/components/schemas/ActivityResponse"}}}, "description": "Returns user activity data grouped by endpoint"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get user activity grouped by endpoint", "tags": ["Analytics"]}
```
