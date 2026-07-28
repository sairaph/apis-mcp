---
title: Update collection item
page_id: operation-patch-accounts-account-id-cloudforce-one-v2-collections-collection-id-it-209934bc
path: operations/collections
description: Update an item in a collection with partial data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items/{item_id}
operation_ids:
    - patch_ItemUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update collection item

`PATCH /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items/{item_id}`

Operation ID: `patch_ItemUpdate`

Update an item in a collection with partial data.

## Definition

```yaml
{"operationId": "patch_ItemUpdate", "summary": "Update collection item", "description": "Update an item in a collection with partial data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "collection_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object", "additionalProperties": true}}, "required": ["data"]}}}}, "responses": {"200": {"description": "Item updated", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string"}, "data": {"type": "object", "additionalProperties": true}, "id": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "data", "created_at", "updated_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "400": {"description": "Validation error"}, "404": {"description": "Collection or item not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
