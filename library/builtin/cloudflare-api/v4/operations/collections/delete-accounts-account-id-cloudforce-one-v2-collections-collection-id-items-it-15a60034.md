---
title: Delete collection item
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-collections-collection-id-i-a4ce5b60
path: operations/collections
description: Delete an item from a collection by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items/{item_id}
operation_ids:
    - delete_ItemDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete collection item

`DELETE /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items/{item_id}`

Operation ID: `delete_ItemDelete`

Delete an item from a collection by its identifier.

## Definition

```yaml
{"operationId": "delete_ItemDelete", "summary": "Delete collection item", "description": "Delete an item from a collection by its identifier.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "collection_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Item deleted", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"deleted": {"type": "boolean"}}, "required": ["deleted"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Collection or item not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
