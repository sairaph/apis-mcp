---
title: Update an observability destination
page_id: operation-patch-observability-destinations-id-fbf17719
path: operations/observability
description: Update an existing observability destination. Only the fields provided in the request body are updated. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /observability/destinations/{id}
operation_ids:
    - updateObservabilityDestination
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Update an observability destination

`PATCH /observability/destinations/{id}`

Operation ID: `updateObservabilityDestination`

Update an existing observability destination. Only the fields provided in the request body are updated. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Update an existing observability destination. Only the fields provided in the request body are updated. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "updateObservabilityDestination", "parameters": [{"description": "The destination ID (UUID).", "in": "path", "name": "id", "required": true, "schema": {"description": "The destination ID (UUID).", "example": "99999999-aaaa-bbbb-cccc-dddddddddddd", "format": "uuid", "type": "string"}}], "requestBody": {"content": {"application/json": {"example": {"enabled": false, "name": "Updated Langfuse"}, "schema": {"$ref": "#/components/schemas/UpdateObservabilityDestinationRequest"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"api_key_hashes": null, "config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "created_at": "2025-08-24T10:30:00Z", "enabled": true, "filter_rules": null, "id": "99999999-aaaa-bbbb-cccc-dddddddddddd", "name": "Production Langfuse", "privacy_mode": false, "sampling_rate": 1, "type": "langfuse", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "schema": {"$ref": "#/components/schemas/UpdateObservabilityDestinationResponse"}}}, "description": "Destination updated successfully"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "409": {"content": {"application/json": {"example": {"error": {"code": 409, "message": "Resource conflict. Please try again later."}}, "schema": {"$ref": "#/components/schemas/ConflictResponse"}}}, "description": "Conflict - Resource conflict or concurrent modification"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Update an observability destination", "tags": ["Observability"], "x-speakeasy-name-override": "update"}
```
