---
title: List Zones
page_id: operation-get-zones-00ec851a
path: operations/zone
description: |-
    Lists, searches, sorts, and filters your zones. Listing zones across more than 500 accounts
    is currently not allowed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones
operation_ids:
    - zones-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zones

`GET /zones`

Operation ID: `zones-get`

Lists, searches, sorts, and filters your zones. Listing zones across more than 500 accounts
is currently not allowed.

## Definition

```yaml
{"operationId": "zones-get", "summary": "List Zones", "description": "Lists, searches, sorts, and filters your zones. Listing zones across more than 500 accounts\nis currently not allowed.\n", "parameters": [{"name": "name", "in": "query", "schema": {"description": "A domain name. Optional filter operators can be provided to extend refine the search:\n  * `equal` (default)\n  * `not_equal`\n  * `starts_with`\n  * `ends_with`\n  * `contains`\n  * `starts_with_case_sensitive`\n  * `ends_with_case_sensitive`\n  * `contains_case_sensitive`\n", "type": "string", "maxLength": 253}, "examples": {"Basic Query": {"summary": "Simple Query", "value": "example.com"}, "Contains Query": {"summary": "Contains Query", "value": "contains:.org"}, "Ends With Query": {"summary": "Ends With Query", "value": "ends_with:arpa"}, "Starts With Query": {"summary": "Starts With Query", "value": "starts_with:dev"}}}, {"name": "status", "in": "query", "schema": {"description": "Specify a zone status to filter by.", "type": "string", "enum": ["initializing", "pending", "active", "moved"]}}, {"name": "type", "in": "query", "schema": {"description": "Zone types to filter by. Multiple types can be specified as a comma-separated list (e.g., ?type=full,partial,secondary). When this parameter is not provided, zones with type \"internal\" are excluded from the results.", "type": "array", "items": {"enum": ["full", "partial", "secondary", "internal"], "type": "string"}}, "explode": false, "style": "form"}, {"name": "account.id", "in": "query", "schema": {"description": "Filter by an account ID.", "type": "string"}}, {"name": "account.name", "in": "query", "schema": {"description": "An account Name. Optional filter operators can be provided to extend refine the search:\n  * `equal` (default)\n  * `not_equal`\n  * `starts_with`\n  * `ends_with`\n  * `contains`\n  * `starts_with_case_sensitive`\n  * `ends_with_case_sensitive`\n  * `contains_case_sensitive`\n", "type": "string", "maxLength": 253}, "examples": {"Basic Query": {"summary": "Simple Query", "value": "Dev Account"}, "Contains Query": {"summary": "Contains Query", "value": "contains:Test"}}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of zones per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Field to order zones by.", "type": "string", "example": "status", "enum": ["name", "status", "account.id", "account.name", "plan.id"]}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order zones.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "match", "in": "query", "schema": {"description": "Whether to match all search requirements or at least one (any).", "type": "string", "default": "all", "enum": ["any", "all"]}}], "responses": {"200": {"description": "List Zones response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common"}, {"properties": {"result_info": {"$ref": "#/components/schemas/zones_result_info"}}}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/zones_zone"}}}}]}}}}, "4XX": {"description": "List Zones response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Zone Zone Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones", "x-fern-sdk-method-name": "list"}
```
