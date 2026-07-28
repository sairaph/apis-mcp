---
title: Create item(s)
page_id: operation-post-accounts-account-id-cloudforce-one-v2-collections-collection-id-ite-3774b3c8
path: operations/collections-items
description: Create one or more items in a collection. Supports single item (data field) or bulk creation (data array). Item data is validated against the collection schema.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items
operation_ids:
    - post_ItemCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create item(s)

`POST /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items`

Operation ID: `post_ItemCreate`

Create one or more items in a collection. Supports single item (data field) or bulk creation (data array). Item data is validated against the collection schema.

## Definition

```yaml
{"operationId": "post_ItemCreate", "summary": "Create item(s)", "description": "Create one or more items in a collection. Supports single item (data field) or bulk creation (data array). Item data is validated against the collection schema.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "collection_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"anyOf": [{"properties": {"data": {"description": "Single item data matching collection schema", "type": "object", "additionalProperties": true}}, "required": ["data"], "type": "object"}, {"properties": {"data": {"description": "Array of items to create (max 1000)", "type": "array", "items": {"additionalProperties": true, "type": "object"}, "maxItems": 1000, "minItems": 1}}, "required": ["data"], "type": "object"}]}}}}, "responses": {"201": {"description": "Item(s) created successfully", "content": {"application/json": {"schema": {"anyOf": [{"properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string"}, "data": {"type": "object", "additionalProperties": true}, "id": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "data", "created_at", "updated_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"], "type": "object"}, {"properties": {"result": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"data": {"type": "object", "additionalProperties": true}, "error": {"type": "string"}, "index": {"type": "number"}}, "required": ["index", "error", "data"], "type": "object"}}, "failed": {"type": "number"}, "inserted": {"type": "number"}}, "required": ["inserted"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"], "type": "object"}]}}}}, "400": {"description": "Validation error"}, "404": {"description": "Collection not found"}}, "security": [{"api_token": []}], "tags": ["Collections - Items"], "x-api-token-group": ["Cloudforce One Write"]}
```
