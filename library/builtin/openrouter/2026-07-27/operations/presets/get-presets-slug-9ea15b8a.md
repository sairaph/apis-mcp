---
title: Get a preset
page_id: operation-get-presets-slug-e2c54600
path: operations/presets
description: Retrieves a preset by its slug with its currently designated version inline.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /presets/{slug}
operation_ids:
    - getPreset
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get a preset

`GET /presets/{slug}`

Operation ID: `getPreset`

Retrieves a preset by its slug with its currently designated version inline.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Retrieves a preset by its slug with its currently designated version inline.", "operationId": "getPreset", "parameters": [{"description": "URL-safe slug identifying the preset.", "in": "path", "name": "slug", "required": true, "schema": {"description": "URL-safe slug identifying the preset.", "example": "my-preset", "minLength": 1, "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"created_at": "2026-04-20T10:00:00Z", "creator_user_id": "user_2dHFtVWx2n56w6HkM0000000000", "description": null, "designated_version": {"config": {"model": "openai/gpt-4o", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}, "designated_version_id": "550e8400-e29b-41d4-a716-446655440000", "id": "650e8400-e29b-41d4-a716-446655440001", "name": "my-preset", "slug": "my-preset", "status": "active", "status_updated_at": null, "updated_at": "2026-04-20T10:00:00Z", "workspace_id": "750e8400-e29b-41d4-a716-446655440002"}}, "schema": {"$ref": "#/components/schemas/GetPresetResponse"}}}, "description": "Preset with its designated version."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "security": [{"apiKey": []}], "summary": "Get a preset", "tags": ["Presets"], "x-speakeasy-name-override": "get"}
```
