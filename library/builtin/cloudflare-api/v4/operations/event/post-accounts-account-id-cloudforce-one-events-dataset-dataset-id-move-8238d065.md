---
title: Moves specified events from one dataset to another dataset
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-dataset-id-move-d2288948
path: operations/event
description: Move one or more events from a source dataset to a destination dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/move
operation_ids:
    - post_EventMoveToNewDS
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Moves specified events from one dataset to another dataset

`POST /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/move`

Operation ID: `post_EventMoveToNewDS`

Move one or more events from a source dataset to a destination dataset.

## Definition

```yaml
{"operationId": "post_EventMoveToNewDS", "summary": "Moves specified events from one dataset to another dataset", "description": "Move one or more events from a source dataset to a destination dataset.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}, {"name": "keepRawData", "in": "query", "description": "If true, copies raw data to the destination dataset. Default is false (raw data is stripped/not copied). Raw data is always deleted from the source.", "schema": {"description": "If true, copies raw data to the destination dataset. Default is false (raw data is stripped/not copied). Raw data is always deleted from the source.", "type": "boolean"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"destDatasetId": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}, "eventIds": {"type": "array", "items": {"example": "7632a037-fdef-4899-9b12-148470aae772", "type": "string"}}}, "required": ["destDatasetId", "eventIds"]}}}}, "responses": {"200": {"description": "Returns the number of moved events ", "content": {"application/json": {"schema": {"type": "object", "properties": {"deletionFailures": {"description": "Array of source datasets where deletion failed", "type": "array", "items": {"properties": {"datasetId": {"description": "Dataset ID where deletion failed", "type": "string"}, "reason": {"description": "Reason for the deletion failure", "type": "string"}}, "required": ["datasetId", "reason"], "type": "object"}}, "indicatorsCopied": {"description": "Number of indicators successfully copied", "type": "number", "example": 5}, "insertFailures": {"description": "Array of events that failed to insert into destination", "type": "array", "items": {"properties": {"index": {"description": "Index of the event that failed to insert", "type": "number"}, "reason": {"description": "Reason for the failure", "type": "string"}}, "required": ["index", "reason"], "type": "object"}}, "moved": {"description": "Number of events successfully moved", "type": "number", "example": 2}, "relationshipsCopied": {"description": "Number of relationships successfully copied", "type": "number", "example": 3}}, "required": ["moved", "indicatorsCopied", "relationshipsCopied"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
