---
title: Download generated video content
page_id: operation-get-videos-jobid-content-b9a2c313
path: operations/video-generation
description: Streams the generated video content from the upstream provider
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /videos/{jobId}/content
operation_ids:
    - listVideosContent
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Download generated video content

`GET /videos/{jobId}/content`

Operation ID: `listVideosContent`

Streams the generated video content from the upstream provider

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Streams the generated video content from the upstream provider", "operationId": "listVideosContent", "parameters": [{"in": "path", "name": "jobId", "required": true, "schema": {"example": "job-abc123", "minLength": 1, "type": "string"}}, {"in": "query", "name": "index", "required": false, "schema": {"default": 0, "example": 0, "minimum": 0, "type": ["integer", "null"]}}], "responses": {"200": {"content": {"video/mp4": {"example": "<binary video data>", "schema": {"format": "binary", "type": "string"}}}, "description": "Video content stream. The body is the raw video bytes proxied from the upstream provider, and the Content-Type reflects the provider media type (video/mp4)."}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}, "502": {"content": {"application/json": {"example": {"error": {"code": 502, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/BadGatewayResponse"}}}, "description": "Bad Gateway - Provider/upstream API failure"}}, "summary": "Download generated video content", "tags": ["Video Generation"], "x-speakeasy-name-override": "getVideoContent"}
```
