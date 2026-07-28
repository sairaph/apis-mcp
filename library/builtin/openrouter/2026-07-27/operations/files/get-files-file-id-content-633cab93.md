---
title: Download file content
page_id: operation-get-files-file-id-content-12bbf31f
path: operations/files
description: Downloads the raw bytes of a file. Only files created server-side are downloadable; uploaded files return 400.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /files/{file_id}/content
operation_ids:
    - downloadFileContent
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Download file content

`GET /files/{file_id}/content`

Operation ID: `downloadFileContent`

Downloads the raw bytes of a file. Only files created server-side are downloadable; uploaded files return 400.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Downloads the raw bytes of a file. Only files created server-side are downloadable; uploaded files return 400.", "operationId": "downloadFileContent", "parameters": [{"in": "path", "name": "file_id", "required": true, "schema": {"example": "file_011CNha8iCJcU1wXNR6q4V8w", "type": "string"}}, {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "example": "a103d8b6-42f0-4e50-9a3c-bf41e2c3c1a7", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/octet-stream": {"example": "binary file contents", "schema": {"format": "binary", "type": "string"}}}, "description": "The raw file content."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Download file content", "tags": ["Files"], "x-hidden": true, "x-speakeasy-name-override": "download"}
```
