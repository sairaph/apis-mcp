---
title: List group members
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-groups-group-id-me-e1ab2894
path: operations/groups
description: List group members
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}/members
operation_ids:
    - get_GroupMemberList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List group members

`GET /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}/members`

Operation ID: `get_GroupMemberList`

List group members

## Definition

```yaml
{"operationId": "get_GroupMemberList", "summary": "List group members", "description": "List group members", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}], "responses": {"200": {"description": "Returns the group members.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"accountId": {"type": "string", "example": "123"}, "accountTag": {"type": "string", "example": "5e785ae7a11e4c3c8df4bce13f1b9c02"}, "createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "accountId"], "type": "object"}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
