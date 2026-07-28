---
title: Delete a dataset
page_id: operation-delete-accounts-account-id-cloudforce-one-events-dataset-dataset-id-45502374
path: operations/dataset
description: Soft-deletes a dataset given a datasetId.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}
operation_ids:
    - delete_DatasetDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a dataset

`DELETE /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}`

Operation ID: `delete_DatasetDelete`

Soft-deletes a dataset given a datasetId.

## Definition

```yaml
{"operationId": "delete_DatasetDelete", "summary": "Delete a dataset", "description": "Soft-deletes a dataset given a datasetId.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID to delete", "required": true, "schema": {"description": "Dataset ID to delete", "type": "string"}}], "responses": {"200": {"description": "Returns the uuid and name of the deleted dataset.", "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"type": "string", "example": "friendly dataset name"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Dataset"], "x-api-token-group": ["Cloudforce One Write"]}
```
