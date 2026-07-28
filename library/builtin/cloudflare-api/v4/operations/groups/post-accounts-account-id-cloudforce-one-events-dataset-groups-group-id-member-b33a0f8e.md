---
title: Create a group member
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-groups-group-id-m-50423bfc
path: operations/groups
description: Create a group member
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}/members
operation_ids:
    - post_GroupMemberCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a group member

`POST /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}/members`

Operation ID: `post_GroupMemberCreate`

Create a group member

## Definition

```yaml
{"operationId": "post_GroupMemberCreate", "summary": "Create a group member", "description": "Create a group member", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"accountId": {"type": "string", "example": "123"}, "accountTag": {"type": "string", "example": "5e785ae7a11e4c3c8df4bce13f1b9c02"}}}}}}, "responses": {"200": {"description": "Returns the created group member.", "content": {"application/json": {"schema": {"type": "object", "properties": {"accountId": {"type": "string", "example": "123"}, "accountTag": {"type": "string", "example": "5e785ae7a11e4c3c8df4bce13f1b9c02"}, "createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "accountId"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
