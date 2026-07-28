---
title: Delete a BYOK provider credential
page_id: operation-delete-byok-id-31461f13
path: operations/byok
description: Delete (soft-delete) a bring-your-own-key (BYOK) provider credential by its `id`. The encrypted key material is wiped and the record is marked as deleted. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /byok/{id}
operation_ids:
    - deleteBYOKKey
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete a BYOK provider credential

`DELETE /byok/{id}`

Operation ID: `deleteBYOKKey`

Delete (soft-delete) a bring-your-own-key (BYOK) provider credential by its `id`. The encrypted key material is wiped and the record is marked as deleted. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Delete (soft-delete) a bring-your-own-key (BYOK) provider credential by its `id`. The encrypted key material is wiped and the record is marked as deleted. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "deleteBYOKKey", "parameters": [{"description": "The BYOK credential ID (UUID).", "in": "path", "name": "id", "required": true, "schema": {"description": "The BYOK credential ID (UUID).", "example": "11111111-2222-3333-4444-555555555555", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"deleted": true}, "schema": {"$ref": "#/components/schemas/DeleteBYOKKeyResponse"}}}, "description": "BYOK credential deleted successfully"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete a BYOK provider credential", "tags": ["BYOK"], "x-speakeasy-name-override": "delete"}
```
