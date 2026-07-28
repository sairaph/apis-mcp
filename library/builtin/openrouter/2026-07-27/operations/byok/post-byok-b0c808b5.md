---
title: Create a BYOK provider credential
page_id: operation-post-byok-c9c6c07f
path: operations/byok
description: Create a new bring-your-own-key (BYOK) provider credential. The raw key is encrypted at rest and never returned in API responses. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. Treat the raw key as write-only; it is never returned after creation. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /byok
operation_ids:
    - createBYOKKey
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create a BYOK provider credential

`POST /byok`

Operation ID: `createBYOKKey`

Create a new bring-your-own-key (BYOK) provider credential. The raw key is encrypted at rest and never returned in API responses. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. Treat the raw key as write-only; it is never returned after creation. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Create a new bring-your-own-key (BYOK) provider credential. The raw key is encrypted at rest and never returned in API responses. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. Treat the raw key as write-only; it is never returned after creation. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "createBYOKKey", "requestBody": {"content": {"application/json": {"example": {"key": "sk-proj-abc123...", "name": "Production OpenAI Key", "provider": "openai"}, "schema": {"$ref": "#/components/schemas/CreateBYOKKeyRequest"}}}, "required": true}, "responses": {"201": {"content": {"application/json": {"example": {"data": {"allowed_api_key_hashes": null, "allowed_models": null, "allowed_user_ids": null, "created_at": "2025-08-24T10:30:00Z", "disabled": false, "id": "11111111-2222-3333-4444-555555555555", "is_fallback": false, "label": "sk-...AbCd", "name": "Production OpenAI Key", "provider": "openai", "sort_order": 0, "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "schema": {"$ref": "#/components/schemas/CreateBYOKKeyResponse"}}}, "description": "BYOK credential created successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Create a BYOK provider credential", "tags": ["BYOK"], "x-speakeasy-name-override": "create"}
```
