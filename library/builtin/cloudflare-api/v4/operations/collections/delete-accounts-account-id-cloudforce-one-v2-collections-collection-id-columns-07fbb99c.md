---
title: Delete column
page_id: operation-delete-accounts-account-id-cloudforce-one-v2-collections-collection-id-c-fe75baca
path: operations/collections
description: Delete a column from the collection schema. Data is preserved as orphaned UUID keys (forensic safety) and filtered from API responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/columns/{column_id}
operation_ids:
    - delete_ColumnDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete column

`DELETE /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/columns/{column_id}`

Operation ID: `delete_ColumnDelete`

Delete a column from the collection schema. Data is preserved as orphaned UUID keys (forensic safety) and filtered from API responses.

## Definition

```yaml
{"operationId": "delete_ColumnDelete", "summary": "Delete column", "description": "Delete a column from the collection schema. Data is preserved as orphaned UUID keys (forensic safety) and filtered from API responses.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "collection_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "column_id", "in": "path", "description": "Column UUID", "required": true, "schema": {"description": "Column UUID", "type": "string"}}], "responses": {"200": {"description": "Column deleted successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"type": "string"}, "name": {"type": "string"}, "position": {"type": "number"}, "required": {"type": "boolean"}, "type": {"type": "string"}}, "required": ["id", "name", "type", "required", "position"]}}, "required": ["result"]}}}}, "400": {"description": "Cannot delete imported column"}, "404": {"description": "Collection or column not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
