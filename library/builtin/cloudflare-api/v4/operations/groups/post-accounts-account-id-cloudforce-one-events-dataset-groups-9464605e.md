---
title: Create a group
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-groups-c8a37161
path: operations/groups
description: Create a group
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/-/groups
operation_ids:
    - post_GroupCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a group

`POST /accounts/{account_id}/cloudforce-one/events/dataset/-/groups`

Operation ID: `post_GroupCreate`

Create a group

## Definition

```yaml
{"operationId": "post_GroupCreate", "summary": "Create a group", "description": "Create a group", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "example": "Cloudforce subscribers"}, "name": {"type": "string", "example": "cloudforce-subscribers"}}, "required": ["name", "description"]}}}}, "responses": {"200": {"description": "Returns the created group.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "description": {"type": "string", "example": "Cloudforce subscribers"}, "name": {"type": "string", "example": "loudforce-subscribers"}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "description", "createdAt", "updatedAt"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Groups"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
