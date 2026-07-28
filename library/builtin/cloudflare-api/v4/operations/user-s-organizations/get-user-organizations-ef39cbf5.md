---
title: List Organizations
page_id: operation-get-user-organizations-28152392
path: operations/user-s-organizations
description: Lists organizations the user is associated with.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/organizations
operation_ids:
    - user'-s-organizations-list-organizations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Organizations

`GET /user/organizations`

Operation ID: `user'-s-organizations-list-organizations`

Lists organizations the user is associated with.

## Definition

```yaml
{"operationId": "user'-s-organizations-list-organizations", "summary": "List Organizations", "description": "Lists organizations the user is associated with.", "parameters": [{"name": "name", "in": "query", "schema": {"$ref": "#/components/schemas/iam_schemas-name"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of organizations per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Field to order organizations by.", "type": "string", "example": "status", "enum": ["id", "name", "status"]}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order organizations.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "match", "in": "query", "schema": {"description": "Whether to match all search requirements or at least one (any).", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "status", "in": "query", "schema": {"description": "Whether the user is a member of the organization or has an inivitation pending.", "type": "string", "example": "member", "enum": ["member", "invited"]}}], "responses": {"200": {"description": "List Organizations response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_organization_response"}}}}, "4XX": {"description": "List Organizations response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Organizations"], "x-api-token-group": ["Memberships Write", "Memberships Read"], "x-cfDeprecation": {"description": "This endpoint and its related APIs are deprecated in favor of the `/accounts` equivalent APIs, which have a broader range of features and are backwards compatible with these API.", "display": true, "eol": "2020-02-04", "id": "org_deprecation"}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
