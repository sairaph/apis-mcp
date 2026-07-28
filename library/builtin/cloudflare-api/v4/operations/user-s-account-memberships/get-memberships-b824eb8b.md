---
title: List Memberships
page_id: operation-get-memberships-f9ab28ee
path: operations/user-s-account-memberships
description: List memberships of accounts the user can access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /memberships
operation_ids:
    - user'-s-account-memberships-list-memberships
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Memberships

`GET /memberships`

Operation ID: `user'-s-account-memberships-list-memberships`

List memberships of accounts the user can access.

## Definition

```yaml
{"operationId": "user'-s-account-memberships-list-memberships", "summary": "List Memberships", "description": "List memberships of accounts the user can access.", "parameters": [{"name": "account.name", "in": "query", "schema": {"$ref": "#/components/schemas/iam_properties-name"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of memberships per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Field to order memberships by.", "type": "string", "example": "status", "enum": ["id", "account.name", "status"]}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order memberships.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"name": "name", "in": "query", "schema": {"$ref": "#/components/schemas/iam_properties-name"}}, {"name": "status", "in": "query", "schema": {"description": "Status of this membership.", "type": "string", "example": "accepted", "enum": ["accepted", "pending", "rejected"]}}], "responses": {"200": {"description": "List Memberships response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/iam_collection_membership_response"}, {"$ref": "#/components/schemas/iam_collection_membership_response_with_policies"}]}}}}, "4XX": {"description": "List Memberships response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User's Account Memberships"], "x-api-token-group": ["Memberships Write", "Memberships Read"]}
```
