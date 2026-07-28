---
title: List versions of a preset
page_id: operation-get-presets-slug-versions-c267c528
path: operations/presets
description: Lists all versions of a preset, ordered by version number ascending (oldest first).
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /presets/{slug}/versions
operation_ids:
    - listPresetVersions
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List versions of a preset

`GET /presets/{slug}/versions`

Operation ID: `listPresetVersions`

Lists all versions of a preset, ordered by version number ascending (oldest first).

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Lists all versions of a preset, ordered by version number ascending (oldest first).", "operationId": "listPresetVersions", "parameters": [{"description": "URL-safe slug identifying the preset.", "in": "path", "name": "slug", "required": true, "schema": {"description": "URL-safe slug identifying the preset.", "example": "my-preset", "minLength": 1, "type": "string"}}, {"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"config": {"model": "openai/gpt-4o", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListPresetVersionsResponse"}}}, "description": "Paginated list of preset versions."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "security": [{"apiKey": []}], "summary": "List versions of a preset", "tags": ["Presets"], "x-speakeasy-name-override": "listVersions", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
