---
title: Delete collection
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-collections-collection-id-edeec1c9
path: operations/collections
description: Deletes a collection and all its items. This action cannot be undone. The Durable Object storage is deleted asynchronously.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}
operation_ids:
    - delete_CollectionDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete collection

`DELETE /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}`

Operation ID: `delete_CollectionDelete`

Deletes a collection and all its items. This action cannot be undone. The Durable Object storage is deleted asynchronously.

## Definition

```yaml
{"operationId": "delete_CollectionDelete", "summary": "Delete collection", "description": "Deletes a collection and all its items. This action cannot be undone. The Durable Object storage is deleted asynchronously.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string"}}], "responses": {"200": {"description": "Collection deleted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}}, "messages": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages"]}}}}, "404": {"description": "Collection not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
