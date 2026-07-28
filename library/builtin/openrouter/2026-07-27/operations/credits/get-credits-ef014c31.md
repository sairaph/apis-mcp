---
title: Get remaining credits
page_id: operation-get-credits-1529373c
path: operations/credits
description: Get total credits purchased and used for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /credits
operation_ids:
    - getCredits
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Get remaining credits

`GET /credits`

Operation ID: `getCredits`

Get total credits purchased and used for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Get total credits purchased and used for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "getCredits", "responses": {"200": {"content": {"application/json": {"example": {"data": {"total_credits": 100.5, "total_usage": 25.75}}, "schema": {"description": "Total credits purchased and used", "example": {"data": {"total_credits": 100.5, "total_usage": 25.75}}, "properties": {"data": {"example": {"total_credits": 100.5, "total_usage": 25.75}, "properties": {"total_credits": {"description": "Total credits purchased", "example": 100.5, "format": "double", "type": "number"}, "total_usage": {"description": "Total credits used", "example": 25.75, "format": "double", "type": "number"}}, "required": ["total_credits", "total_usage"], "type": "object"}}, "required": ["data"], "type": "object"}}}, "description": "Returns the total credits purchased and used"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Get remaining credits", "tags": ["Credits"], "x-speakeasy-name-override": "getCredits"}
```
