---
title: Delete an API key
page_id: operation-delete-keys-hash-0e96d819
path: operations/api-keys
description: Delete an existing API key. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /keys/{hash}
operation_ids:
    - deleteKeys
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Delete an API key

`DELETE /keys/{hash}`

Operation ID: `deleteKeys`

Delete an existing API key. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Delete an existing API key. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "deleteKeys", "parameters": [{"description": "The hash identifier of the API key to delete", "in": "path", "name": "hash", "required": true, "schema": {"description": "The hash identifier of the API key to delete", "example": "f01d52606dc8f0a8303a7b5cc3fa07109c2e346cec7c0a16b40de462992ce943", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"deleted": true}, "schema": {"example": {"deleted": true}, "properties": {"deleted": {"const": true, "description": "Confirmation that the API key was deleted", "example": true, "type": "boolean"}}, "required": ["deleted"], "type": "object"}}}, "description": "API key deleted successfully"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "429": {"content": {"application/json": {"example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "schema": {"$ref": "#/components/schemas/TooManyRequestsResponse"}}}, "description": "Too Many Requests - Rate limit exceeded"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Delete an API key", "tags": ["API Keys"], "x-speakeasy-name-override": "delete"}
```
