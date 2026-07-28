---
title: List groups for an account
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-groups-ddb62ae2
path: operations/groups
description: List groups for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups
operation_ids:
    - get_GroupList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List groups for an account

`GET /accounts/{account_id}/cloudforce-one/events/dataset/-/groups`

Operation ID: `get_GroupList`

List groups for an account

## Definition

```yaml
{"operationId": "get_GroupList", "summary": "List groups for an account", "description": "List groups for an account", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns the list of groups.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "description": {"type": "string", "example": "Cloudforce subscribers"}, "name": {"type": "string", "example": "loudforce-subscribers"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "description", "createdAt", "updatedAt"], "type": "object"}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
