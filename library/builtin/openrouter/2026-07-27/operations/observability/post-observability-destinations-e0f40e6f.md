---
title: Create an observability destination
page_id: operation-post-observability-destinations-bca1289d
path: operations/observability
description: Create a new observability destination. A maximum of 5 destinations per type is allowed. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /observability/destinations
operation_ids:
    - createObservabilityDestination
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create an observability destination

`POST /observability/destinations`

Operation ID: `createObservabilityDestination`

Create a new observability destination. A maximum of 5 destinations per type is allowed. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Create a new observability destination. A maximum of 5 destinations per type is allowed. Defaults to the authenticated entity's default workspace; use the `workspace_id` body field to scope to a different workspace. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "createObservabilityDestination", "requestBody": {"content": {"application/json": {"example": {"config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "name": "Production Langfuse", "type": "langfuse"}, "schema": {"$ref": "#/components/schemas/CreateObservabilityDestinationRequest"}}}, "required": true}, "responses": {"201": {"content": {"application/json": {"example": {"data": {"api_key_hashes": null, "config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "created_at": "2025-08-24T10:30:00Z", "enabled": true, "filter_rules": null, "id": "99999999-aaaa-bbbb-cccc-dddddddddddd", "name": "Production Langfuse", "privacy_mode": false, "sampling_rate": 1, "type": "langfuse", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "schema": {"$ref": "#/components/schemas/CreateObservabilityDestinationResponse"}}}, "description": "Destination created successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "409": {"content": {"application/json": {"example": {"error": {"code": 409, "message": "Resource conflict. Please try again later."}}, "schema": {"$ref": "#/components/schemas/ConflictResponse"}}}, "description": "Conflict - Resource conflict or concurrent modification"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Create an observability destination", "tags": ["Observability"], "x-speakeasy-name-override": "create"}
```
