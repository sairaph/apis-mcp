---
title: List observability destinations
page_id: operation-get-observability-destinations-4dfe476d
path: operations/observability
description: List the observability destinations configured for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace. Only destinations with stable release status are surfaced — destinations of other types are excluded. [Management key](/docs/guides/overview/auth/management-api-keys) required.
source: https://openrouter.ai/openapi.json
http_methods:
    - GET
api_endpoints:
    - /observability/destinations
operation_ids:
    - listObservabilityDestinations
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# List observability destinations

`GET /observability/destinations`

Operation ID: `listObservabilityDestinations`

List the observability destinations configured for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace. Only destinations with stable release status are surfaced — destinations of other types are excluded. [Management key](/docs/guides/overview/auth/management-api-keys) required.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "List the observability destinations configured for the authenticated entity's default workspace. Use the `workspace_id` query parameter to scope the result to a different workspace. Only destinations with stable release status are surfaced — destinations of other types are excluded. [Management key](/docs/guides/overview/auth/management-api-keys) required.", "operationId": "listObservabilityDestinations", "parameters": [{"description": "Number of records to skip for pagination", "in": "query", "name": "offset", "required": false, "schema": {"default": 0, "description": "Number of records to skip for pagination", "example": 0, "minimum": 0, "type": ["integer", "null"]}}, {"description": "Maximum number of records to return (max 100)", "in": "query", "name": "limit", "required": false, "schema": {"default": 50, "description": "Maximum number of records to return (max 100)", "example": 50, "maximum": 100, "minimum": 1, "type": "integer"}}, {"description": "Optional workspace ID to filter by. Defaults to the authenticated entity's default workspace.", "in": "query", "name": "workspace_id", "required": false, "schema": {"description": "Optional workspace ID to filter by. Defaults to the authenticated entity's default workspace.", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}], "responses": {"200": {"content": {"application/json": {"example": {"data": [{"api_key_hashes": null, "config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "created_at": "2025-08-24T10:30:00Z", "enabled": true, "filter_rules": null, "id": "99999999-aaaa-bbbb-cccc-dddddddddddd", "name": "Production Langfuse", "privacy_mode": false, "sampling_rate": 1, "type": "langfuse", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "total_count": 1}, "schema": {"$ref": "#/components/schemas/ListObservabilityDestinationsResponse"}}}, "description": "List of observability destinations"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "List observability destinations", "tags": ["Observability"], "x-speakeasy-name-override": "list", "x-speakeasy-pagination": {"inputs": [{"in": "parameters", "name": "offset", "type": "offset"}, {"in": "parameters", "name": "limit", "type": "limit"}], "outputs": {"results": "$.data"}, "type": "offsetLimit"}}
```
