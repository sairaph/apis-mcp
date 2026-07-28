---
title: Delete a file
page_id: operation-delete-files-file-id-1da0c2da
path: operations/files
description: Deletes a file owned by the requesting workspace. Deletion is irreversible.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /files/{file_id}
operation_ids:
    - deleteFile
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete a file

`DELETE /files/{file_id}`

Operation ID: `deleteFile`

Deletes a file owned by the requesting workspace. Deletion is irreversible.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Deletes a file owned by the requesting workspace. Deletion is irreversible.", "operationId": "deleteFile", "parameters": [{"in": "path", "name": "file_id", "required": true, "schema": {"example": "file_011CNha8iCJcU1wXNR6q4V8w", "type": "string"}}, {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "example": "a103d8b6-42f0-4e50-9a3c-bf41e2c3c1a7", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"id": "file_011CNha8iCJcU1wXNR6q4V8w", "type": "file_deleted"}, "schema": {"$ref": "#/components/schemas/FileDeleteResponse"}}}, "description": "The file was deleted."}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete a file", "tags": ["Files"], "x-hidden": true, "x-speakeasy-name-override": "delete"}
```
