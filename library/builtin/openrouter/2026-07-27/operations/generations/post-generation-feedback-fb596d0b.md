---
title: Submit feedback for a generation
page_id: operation-post-generation-feedback-3c9795d8
path: operations/generations
description: Submit structured feedback on a generation the authenticated user made. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /generation/feedback
operation_ids:
    - submitGenerationFeedback
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Submit feedback for a generation

`POST /generation/feedback`

Operation ID: `submitGenerationFeedback`

Submit structured feedback on a generation the authenticated user made. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Submit structured feedback on a generation the authenticated user made. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "submitGenerationFeedback", "requestBody": {"content": {"application/json": {"example": {"category": "incorrect_response", "comment": "The model repeated the same paragraph three times.", "generation_id": "gen-3bhGkxlo4XFrqiabUM7NDtwDzWwG"}, "schema": {"$ref": "#/components/schemas/SubmitGenerationFeedbackRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"success": true}}, "schema": {"$ref": "#/components/schemas/SubmitGenerationFeedbackResponse"}}}, "description": "Feedback recorded successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Submit feedback for a generation", "tags": ["Generations"], "x-speakeasy-name-override": "submitFeedback"}
```
