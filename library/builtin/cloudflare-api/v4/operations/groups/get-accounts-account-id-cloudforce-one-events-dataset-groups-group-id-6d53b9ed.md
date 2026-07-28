---
title: Read a group for an account
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-groups-group-id-a930926e
path: operations/groups
description: Read a group for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}
operation_ids:
    - get_GroupRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read a group for an account

`GET /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}`

Operation ID: `get_GroupRead`

Read a group for an account

## Definition

```yaml
{"operationId": "get_GroupRead", "summary": "Read a group for an account", "description": "Read a group for an account", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}], "responses": {"200": {"description": "Return the group.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "description": {"type": "string", "example": "Cloudforce subscribers"}, "members": {"type": "array", "items": {"properties": {"accountId": {"type": "string", "example": "123"}, "accountTag": {"type": "string", "example": "5e785ae7a11e4c3c8df4bce13f1b9c02"}, "createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "accountId"], "type": "object"}}, "name": {"type": "string", "example": "loudforce-subscribers"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "description", "createdAt", "updatedAt", "members"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
