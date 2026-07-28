---
title: Deletes one or more events
page_id: operation-delete-accounts-account-id-cloudforce-one-events-dataset-id-delete-10238964
path: operations/event
description: Delete one or more events from a dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/{dataset_id}/delete
operation_ids:
    - delete_EventDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes one or more events

`DELETE /accounts/{account_id}/cloudforce-one/events/{dataset_id}/delete`

Operation ID: `delete_EventDelete`

Delete one or more events from a dataset.

## Definition

```yaml
{"operationId": "delete_EventDelete", "summary": "Deletes one or more events", "description": "Delete one or more events from a dataset.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}, {"name": "eventIds", "in": "query", "description": "Array of Event IDs to delete.", "required": true, "schema": {"description": "Array of Event IDs to delete.", "type": "array", "items": {"minLength": 1, "type": "string"}}}], "responses": {"200": {"description": "Returns the number of deleted events.", "content": {"application/json": {"schema": {"description": "Number of deleted events", "type": "number"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
