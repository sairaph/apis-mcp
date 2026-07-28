---
title: Poll video generation status
page_id: operation-get-videos-jobid-ef0fb724
path: operations/video-generation
description: Returns job status and content URLs when completed
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /videos/{jobId}
operation_ids:
    - getVideos
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Poll video generation status

`GET /videos/{jobId}`

Operation ID: `getVideos`

Returns job status and content URLs when completed

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Returns job status and content URLs when completed", "operationId": "getVideos", "parameters": [{"in": "path", "name": "jobId", "required": true, "schema": {"example": "job-abc123", "minLength": 1, "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"generation_id": "gen-xyz789", "id": "job-abc123", "polling_url": "/api/v1/videos/job-abc123", "status": "completed", "unsigned_urls": ["https://storage.example.com/video.mp4"], "usage": {"cost": 0.5}}, "schema": {"$ref": "#/components/schemas/VideoGenerationResponse"}}}, "description": "Video generation status"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Poll video generation status", "tags": ["Video Generation"], "x-speakeasy-name-override": "getGeneration"}
```
