---
title: Delete an observability destination
page_id: operation-delete-observability-destinations-id-2be8592d
path: operations/observability
description: Delete an existing observability destination. This performs a soft delete. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /observability/destinations/{id}
operation_ids:
    - deleteObservabilityDestination
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete an observability destination

`DELETE /observability/destinations/{id}`

Operation ID: `deleteObservabilityDestination`

Delete an existing observability destination. This performs a soft delete. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Delete an existing observability destination. This performs a soft delete. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "deleteObservabilityDestination", "parameters": [{"description": "The destination ID (UUID).", "in": "path", "name": "id", "required": true, "schema": {"description": "The destination ID (UUID).", "example": "99999999-aaaa-bbbb-cccc-dddddddddddd", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"deleted": true}, "schema": {"$ref": "#/components/schemas/DeleteObservabilityDestinationResponse"}}}, "description": "Destination deleted successfully"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete an observability destination", "tags": ["Observability"], "x-speakeasy-name-override": "delete"}
```
