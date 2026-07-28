---
title: Generate an image
page_id: operation-post-images-d89f3c2e
path: operations/images
description: Generates an image from a text prompt via the image generation router
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /images
operation_ids:
    - createImages
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Generate an image

`POST /images`

Operation ID: `createImages`

Generates an image from a text prompt via the image generation router

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Generates an image from a text prompt via the image generation router", "operationId": "createImages", "requestBody": {"content": {"application/json": {"example": {"model": "bytedance-seed/seedream-4.5", "prompt": "a red panda astronaut floating in space, studio lighting"}, "schema": {"$ref": "#/components/schemas/ImageGenerationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"created": 1748372400, "data": [{"b64_json": "<base64-encoded-image>"}], "usage": {"completion_tokens": 4175, "cost": 0.04, "prompt_tokens": 0, "total_tokens": 4175}}, "schema": {"$ref": "#/components/schemas/ImageGenerationResponse"}}, "text/event-stream": {"example": {"data": {"b64_json": "<base64-encoded-partial-image>", "partial_image_index": 0, "type": "image_generation.partial_image"}}, "schema": {"$ref": "#/components/schemas/ImageStreamingResponse"}, "x-speakeasy-sse-sentinel": "[DONE]"}}, "description": "Image generation response"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "402": {"content": {"application/json": {"example": {"error": {"code": 402, "message": "Insufficient credits. Add more using https://openrouter.ai/credits"}}, "schema": {"$ref": "#/components/schemas/PaymentRequiredResponse"}}}, "description": "Payment Required - Insufficient credits or quota to complete request"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "413": {"content": {"application/json": {"example": {"error": {"code": 413, "message": "Request payload too large"}}, "schema": {"$ref": "#/components/schemas/PayloadTooLargeResponse"}}}, "description": "Payload Too Large - Request payload exceeds size limits"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}, "502": {"content": {"application/json": {"example": {"error": {"code": 502, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/BadGatewayResponse"}}}, "description": "Bad Gateway - Provider/upstream API failure"}, "524": {"content": {"application/json": {"example": {"error": {"code": 524, "message": "Request timed out. Please try again later."}}, "schema": {"$ref": "#/components/schemas/EdgeNetworkTimeoutResponse"}}}, "description": "Infrastructure Timeout - Provider request timed out at edge network"}, "529": {"content": {"application/json": {"example": {"error": {"code": 529, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/ProviderOverloadedResponse"}}}, "description": "Provider Overloaded - Provider is temporarily overloaded"}}, "summary": "Generate an image", "tags": ["Images"], "x-speakeasy-name-override": "generate"}
```
