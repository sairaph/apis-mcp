---
title: Update a group
page_id: operation-put-accounts-account-id-cloudforce-one-events-dataset-groups-group-id-08385eb2
path: operations/groups
description: Update a group
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}
operation_ids:
    - put_GroupUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a group

`PUT /accounts/{account_id}/cloudforce-one/events/dataset/-/groups/{group_id}`

Operation ID: `put_GroupUpdate`

Update a group

## Definition

```yaml
{"operationId": "put_GroupUpdate", "summary": "Update a group", "description": "Update a group", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "example": "Cloudforce subscribers"}, "name": {"type": "string", "example": "cloudforce-subscribers"}}, "required": ["name", "description"]}}}}, "responses": {"200": {"description": "Returns the updated group.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "description": {"type": "string", "example": "Cloudforce subscribers"}, "name": {"type": "string", "example": "loudforce-subscribers"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "description", "createdAt", "updatedAt"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write"]}
```
