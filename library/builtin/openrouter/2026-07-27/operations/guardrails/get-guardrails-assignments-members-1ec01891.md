---
title: List all member assignments
page_id: operation-get-guardrails-assignments-members-86b1457f
path: operations/guardrails
description: List all organization member guardrail assignments for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /guardrails/assignments/members
operation_ids:
    - listMemberAssignments
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List all member assignments

`GET /guardrails/assignments/members`

Operation ID: `listMemberAssignments`

List all organization member guardrail assignments for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List all organization member guardrail assignments for the authenticated user. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listMemberAssignments", "parameters": [{"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "organization_id": "org_xyz789", "user_id": "user_abc123"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListMemberAssignmentsResponse"}}}, "description": "List of member assignments"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List all member assignments", "tags": ["Guardrails"], "x-speakeasy-name-override": "listMemberAssignments", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
