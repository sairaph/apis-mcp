---
title: Update a BYOK provider credential
page_id: operation-patch-byok-id-9592d68d
path: operations/byok
description: Update an existing bring-your-own-key (BYOK) provider credential by its `id`. Include the `key` field to rotate the raw provider API key in-place (the previous key material is overwritten). [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /byok/{id}
operation_ids:
    - updateBYOKKey
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Update a BYOK provider credential

`PATCH /byok/{id}`

Operation ID: `updateBYOKKey`

Update an existing bring-your-own-key (BYOK) provider credential by its `id`. Include the `key` field to rotate the raw provider API key in-place (the previous key material is overwritten). [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Update an existing bring-your-own-key (BYOK) provider credential by its `id`. Include the `key` field to rotate the raw provider API key in-place (the previous key material is overwritten). [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "updateBYOKKey", "parameters": [{"description": "The BYOK credential ID (UUID).", "in": "path", "name": "id", "required": true, "schema": {"description": "The BYOK credential ID (UUID).", "example": "11111111-2222-3333-4444-555555555555", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"disabled": false, "name": "Updated OpenAI Key"}, "schema": {"$ref": "#/components/schemas/UpdateBYOKKeyRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"allowed_api_key_hashes": null, "allowed_models": null, "allowed_user_ids": null, "created_at": "2025-08-24T10:30:00Z", "disabled": false, "id": "11111111-2222-3333-4444-555555555555", "is_fallback": false, "label": "sk-...AbCd", "name": "Updated OpenAI Key", "provider": "openai", "sort_order": 0, "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "schema": {"$ref": "#/components/schemas/UpdateBYOKKeyResponse"}}}, "description": "BYOK credential updated successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Update a BYOK provider credential", "tags": ["BYOK"], "x-speakeasy-name-override": "update"}
```
