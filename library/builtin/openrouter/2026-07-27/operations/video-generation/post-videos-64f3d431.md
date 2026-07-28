---
title: Submit a video generation request
page_id: operation-post-videos-768863b0
path: operations/video-generation
description: Submits a video generation request and returns a polling URL to check status
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /videos
operation_ids:
    - createVideos
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Submit a video generation request

`POST /videos`

Operation ID: `createVideos`

Submits a video generation request and returns a polling URL to check status

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Submits a video generation request and returns a polling URL to check status", "operationId": "createVideos", "requestBody": {"content": {"application/json": {"example": {"aspect_ratio": "16:9", "duration": 8, "model": "google/veo-3.1", "prompt": "A serene mountain landscape at sunset", "resolution": "720p"}, "schema": {"$ref": "#/components/schemas/VideoGenerationRequest"}}}, "required": true}, "responses": {"202": {"content": {"application/json": {"example": {"generation_id": "gen-xyz789", "id": "job-abc123", "polling_url": "/api/v1/videos/job-abc123", "status": "pending"}, "schema": {"$ref": "#/components/schemas/VideoGenerationResponse"}}}, "description": "Video generation request accepted"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "402": {"content": {"application/json": {"example": {"error": {"code": 402, "message": "Insufficient credits. Add more using https://openrouter.ai/credits"}}, "schema": {"$ref": "#/components/schemas/PaymentRequiredResponse"}}}, "description": "Payment Required - Insufficient credits or quota to complete request"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Submit a video generation request", "tags": ["Video Generation"], "x-speakeasy-name-override": "generate"}
```
