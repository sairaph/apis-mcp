---
title: Get stored prompt and completion content for a generation
page_id: operation-get-generation-content-f11b8ebc
path: operations/generations
description: Get stored prompt and completion content for a generation
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /generation/content
operation_ids:
    - listGenerationContent
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get stored prompt and completion content for a generation

`GET /generation/content`

Operation ID: `listGenerationContent`

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"operationId": "listGenerationContent", "parameters": [{"description": "The generation ID", "in": "query", "name": "id", "required": true, "schema": {"description": "The generation ID", "example": "gen-1234567890", "minLength": 1, "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"input": {"messages": [{"content": "What is the meaning of life?", "role": "user"}]}, "output": {"completion": "The meaning of life is a philosophical question...", "reasoning": null}}}, "schema": {"$ref": "#/components/schemas/GenerationContentResponse"}}}, "description": "Returns the stored prompt and completion content"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}, "502": {"content": {"application/json": {"example": {"error": {"code": 502, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/BadGatewayResponse"}}}, "description": "Bad Gateway - Provider/upstream API failure"}, "524": {"content": {"application/json": {"example": {"error": {"code": 524, "message": "Request timed out. Please try again later."}}, "schema": {"$ref": "#/components/schemas/EdgeNetworkTimeoutResponse"}}}, "description": "Infrastructure Timeout - Provider request timed out at edge network"}, "529": {"content": {"application/json": {"example": {"error": {"code": 529, "message": "Provider returned error"}}, "schema": {"$ref": "#/components/schemas/ProviderOverloadedResponse"}}}, "description": "Provider Overloaded - Provider is temporarily overloaded"}}, "summary": "Get stored prompt and completion content for a generation", "tags": ["Generations"]}
```
