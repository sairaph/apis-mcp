---
title: Get collection item
page_id: operation-get-accounts-account-id-cloudforce-one-v2-collections-collection-id-item-0bca47a2
path: operations/collections
description: Retrieve a single item from a collection by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items/{item_id}
operation_ids:
    - get_ItemGet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get collection item

`GET /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items/{item_id}`

Operation ID: `get_ItemGet`

Retrieve a single item from a collection by its identifier.

## Definition

```yaml
{"operationId": "get_ItemGet", "summary": "Get collection item", "description": "Retrieve a single item from a collection by its identifier.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "collection_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "item_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Item retrieved", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string"}, "data": {"type": "object", "additionalProperties": true}, "id": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "data", "created_at", "updated_at"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Collection or item not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
