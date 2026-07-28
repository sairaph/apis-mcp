---
title: Delete a saved event query
page_id: operation-delete-accounts-account-id-cloudforce-one-events-queries-query-id-cc82295e
path: operations/event
description: Delete a saved event query by its ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/queries/{query_id}
operation_ids:
    - delete_EventQueryDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a saved event query

`DELETE /accounts/{account_id}/cloudforce-one/events/queries/{query_id}`

Operation ID: `delete_EventQueryDelete`

Delete a saved event query by its ID

## Definition

```yaml
{"operationId": "delete_EventQueryDelete", "summary": "Delete a saved event query", "description": "Delete a saved event query by its ID", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "query_id", "in": "path", "description": "Event query ID", "required": true, "schema": {"description": "Event query ID", "type": "integer"}}], "responses": {"200": {"description": "Event query deleted successfully."}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
