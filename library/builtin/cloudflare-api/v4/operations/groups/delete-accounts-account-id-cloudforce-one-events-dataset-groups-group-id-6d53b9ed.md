---
title: Delete a group for an account
page_id: operation-delete-accounts-account-id-cloudforce-one-events-dataset-groups-group-id-96a46f28
path: operations/groups
description: Delete a group for an account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}
operation_ids:
    - delete_GroupDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a group for an account

`DELETE /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}`

Operation ID: `delete_GroupDelete`

Delete a group for an account

## Definition

```yaml
{"operationId": "delete_GroupDelete", "summary": "Delete a group for an account", "description": "Delete a group for an account", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}], "responses": {"200": {"description": "Group deleted successfully.", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write"]}
```
