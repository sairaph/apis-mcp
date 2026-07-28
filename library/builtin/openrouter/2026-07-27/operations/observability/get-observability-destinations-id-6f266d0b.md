---
title: Get an observability destination
page_id: operation-get-observability-destinations-id-e315a142
path: operations/observability
description: Fetch a single observability destination by its UUID. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /observability/destinations/{id}
operation_ids:
    - getObservabilityDestination
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get an observability destination

`GET /observability/destinations/{id}`

Operation ID: `getObservabilityDestination`

Fetch a single observability destination by its UUID. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Fetch a single observability destination by its UUID. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "getObservabilityDestination", "parameters": [{"description": "The destination ID (UUID).", "in": "path", "name": "id", "required": true, "schema": {"description": "The destination ID (UUID).", "example": "99999999-aaaa-bbbb-cccc-dddddddddddd", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": {"api_key_hashes": null, "config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "created_at": "2025-08-24T10:30:00Z", "enabled": true, "filter_rules": null, "id": "99999999-aaaa-bbbb-cccc-dddddddddddd", "name": "Production Langfuse", "privacy_mode": false, "sampling_rate": 1, "type": "langfuse", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "schema": {"$ref": "#/components/schemas/GetObservabilityDestinationResponse"}}}, "description": "The observability destination"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get an observability destination", "tags": ["Observability"], "x-speakeasy-name-override": "get"}
```
