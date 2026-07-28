---
title: Create a preset from a chat-completions request body
page_id: operation-post-presets-slug-chat-completions-86e36ab8
path: operations/presets
description: Creates a preset (or a new version of an existing one) from an inference request body. Only fields that overlap with the preset config are persisted; other fields (e.g. `messages`, `stream`, `prompt`) are silently ignored.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /presets/{slug}/chat/completions
operation_ids:
    - createPresetsChatCompletions
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create a preset from a chat-completions request body

`POST /presets/{slug}/chat/completions`

Operation ID: `createPresetsChatCompletions`

Creates a preset (or a new version of an existing one) from an inference request body. Only fields that overlap with the preset config are persisted; other fields (e.g. `messages`, `stream`, `prompt`) are silently ignored.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Creates a preset (or a new version of an existing one) from an inference request body. Only fields that overlap with the preset config are persisted; other fields (e.g. `messages`, `stream`, `prompt`) are silently ignored.", "operationId": "createPresetsChatCompletions", "parameters": [{"description": "URL-safe slug identifying the preset. Created if it does not exist.", "in": "path", "name": "slug", "required": true, "schema": {"description": "URL-safe slug identifying the preset. Created if it does not exist.", "example": "my-preset", "minLength": 1, "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"messages": [{"content": "You are a helpful assistant.", "role": "system"}, {"content": "Hello!", "role": "user"}], "model": "openai/gpt-5.4", "temperature": 0.7}, "schema": {"$ref": "#/components/schemas/ChatRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"created_at": "2026-04-20T10:00:00Z", "creator_user_id": "user_2dHFtVWx2n56w6HkM0000000000", "description": null, "designated_version": {"config": {"model": "openai/gpt-5.4", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}, "designated_version_id": "550e8400-e29b-41d4-a716-446655440000", "id": "650e8400-e29b-41d4-a716-446655440001", "name": "my-preset", "slug": "my-preset", "status": "active", "status_updated_at": null, "updated_at": "2026-04-20T10:00:00Z", "workspace_id": "750e8400-e29b-41d4-a716-446655440002"}}, "schema": {"$ref": "#/components/schemas/CreatePresetFromInferenceResponse"}}}, "description": "Preset created or updated successfully."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "409": {"content": {"application/json": {"example": {"error": {"code": 409, "message": "Resource conflict. Please try again later."}}, "schema": {"$ref": "#/components/schemas/ConflictResponse"}}}, "description": "Conflict - Resource conflict or concurrent modification"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "security": [{"apiKey": []}], "summary": "Create a preset from a chat-completions request body", "tags": ["Presets"]}
```
