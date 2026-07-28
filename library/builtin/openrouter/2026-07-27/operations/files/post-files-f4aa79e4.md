---
title: Upload a file
page_id: operation-post-files-95037ffa
path: operations/files
description: 'Uploads a file to be referenced in future API calls. The file is stored under the workspace of the authenticating API key. Maximum file size: 100 MB.'
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /files
operation_ids:
    - uploadFile
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Upload a file

`POST /files`

Operation ID: `uploadFile`

Uploads a file to be referenced in future API calls. The file is stored under the workspace of the authenticating API key. Maximum file size: 100 MB.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Uploads a file to be referenced in future API calls. The file is stored under the workspace of the authenticating API key. Maximum file size: 100 MB.", "operationId": "uploadFile", "parameters": [{"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Workspace to scope the request to. Defaults to the caller’s default workspace.", "example": "a103d8b6-42f0-4e50-9a3c-bf41e2c3c1a7", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"multipart/form-data": {"example": {"file": "document.pdf"}, "schema": {"properties": {"file": {"format": "binary", "type": "string"}}, "required": ["file"], "type": "object"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"created_at": "2025-01-01T00:00:00Z", "downloadable": false, "filename": "document.pdf", "id": "file_011CNha8iCJcU1wXNR6q4V8w", "mime_type": "application/pdf", "size_bytes": 1024000, "type": "file"}, "schema": {"$ref": "#/components/schemas/FileMetadata"}}}, "description": "The uploaded file metadata."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "413": {"content": {"application/json": {"example": {"error": {"code": 413, "message": "Request payload too large"}}, "schema": {"$ref": "#/components/schemas/PayloadTooLargeResponse"}}}, "description": "Payload Too Large - Request payload exceeds size limits"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Upload a file", "tags": ["Files"], "x-hidden": true, "x-speakeasy-name-override": "upload"}
```
