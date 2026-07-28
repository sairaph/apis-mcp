---
title: Reads a dataset
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-dataset-id-d3acf801
path: operations/dataset
description: Retrieve metadata for a specific dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}
operation_ids:
    - get_DatasetRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reads a dataset

`GET /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}`

Operation ID: `get_DatasetRead`

Retrieve metadata for a specific dataset.

## Definition

```yaml
{"operationId": "get_DatasetRead", "summary": "Reads a dataset", "description": "Retrieve metadata for a specific dataset.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns a dataset.", "content": {"application/json": {"schema": {"type": "object", "properties": {"deletedAt": {"type": "string"}, "isPublic": {"type": "boolean", "example": true}, "name": {"type": "string", "example": "friendly dataset name"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "isPublic"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Dataset"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
