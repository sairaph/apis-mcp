---
title: List organization members
page_id: operation-get-organization-members-4ed00ff0
path: operations/organization
description: List all members of the organization associated with the authenticated management key. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organization/members
operation_ids:
    - listOrganizationMembers
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List organization members

`GET /organization/members`

Operation ID: `listOrganizationMembers`

List all members of the organization associated with the authenticated management key. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all members of the organization associated with the authenticated management key. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listOrganizationMembers", "parameters": [{"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"email": "jane.doe@example.com", "first_name": "Jane", "id": "user_2dHFtVWx2n56w6HkM0000000000", "last_name": "Doe", "role": "org:member"}], "total_count": 25}, "schema": {"properties": {"data": {"description": "List of organization members", "items": {"properties": {"email": {"description": "Email address of the member", "example": "jane.doe@example.com", "type": "string"}, "first_name": {"description": "First name of the member", "example": "Jane", "type": ["string", "null"]}, "id": {"description": "User ID of the organization member", "example": "user_2dHFtVWx2n56w6HkM0000000000", "type": "string"}, "last_name": {"description": "Last name of the member", "example": "Doe", "type": ["string", "null"]}, "role": {"description": "Role of the member in the organization", "enum": ["org:admin", "org:member"], "example": "org:member", "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["id", "first_name", "last_name", "email", "role"], "type": "object"}, "type": "array"}, "total_count": {"description": "Total number of members in the organization", "example": 25, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}}}, "description": "List of organization members"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "404": {"content": {"application/json": {"example": {"error": {"code": 404, "message": "Resource not found"}}, "schema": {"$ref": "#/components/schemas/NotFoundResponse"}}}, "description": "Not Found - Resource does not exist"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List organization members", "tags": ["Organization"], "x-speakeasy-name-override": "listMembers", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
