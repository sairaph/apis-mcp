---
title: List files
page_id: operation-get-files-62814c17
path: operations/files
description: Lists files belonging to the workspace of the authenticating API key.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /files
operation_ids:
    - listFiles
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List files

`GET /files`

Operation ID: `listFiles`

Lists files belonging to the workspace of the authenticating API key.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Lists files belonging to the workspace of the authenticating API key.", "operationId": "listFiles", "parameters": [{"description": "Maximum number of files to return (1–1000).", "in": "query", "name": "limit", "required": false, "schema": {"description": "Maximum number of files to return (1–1000).", "example": 100, "maximum": 1000, "minimum": 1, "type": "integer"}}, {"description": "Opaque pagination cursor from a previous response.", "in": "query", "name": "cursor", "required": false, "schema": {"description": "Opaque pagination cursor from a previous response.", "example": "eyJjdXJzb3IiOiJmaWxlXzAxMUNOaGE4aUNKY1Uxd1hOUjZxNFY4dyJ9", "type": "string"}}, {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "example": "a103d8b6-42f0-4e50-9a3c-bf41e2c3c1a7", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"cursor": null, "data": [{"created_at": "2025-01-01T00:00:00Z", "downloadable": false, "filename": "document.pdf", "id": "file_011CNha8iCJcU1wXNR6q4V8w", "mime_type": "application/pdf", "size_bytes": 1024000, "type": "file"}], "first_id": "file_011CNha8iCJcU1wXNR6q4V8w", "has_more": false, "last_id": "file_011CNha8iCJcU1wXNR6q4V8w"}, "schema": {"$ref": "#/components/schemas/FileListResponse"}}}, "description": "A page of files."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List files", "tags": ["Files"], "x-hidden": true, "x-speakeasy-name-override": "list", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "cursor", "type": "cursor"}], "outputs": {"nextCursor": "$.cursor", "results": "$.data"}, "type": "cursor"}}
```
