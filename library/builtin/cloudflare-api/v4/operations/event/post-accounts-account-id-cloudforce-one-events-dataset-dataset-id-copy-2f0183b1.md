---
title: Copies specified events from one dataset to another dataset
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-dataset-id-copy-33b226d5
path: operations/event
description: Copy one or more events from a source dataset to a destination dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/copy
operation_ids:
    - post_EventCopyToNewDS
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Copies specified events from one dataset to another dataset

`POST /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/copy`

Operation ID: `post_EventCopyToNewDS`

Copy one or more events from a source dataset to a destination dataset.

## Definition

```yaml
{"operationId": "post_EventCopyToNewDS", "summary": "Copies specified events from one dataset to another dataset", "description": "Copy one or more events from a source dataset to a destination dataset.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}, {"name": "keepRawData", "in": "query", "description": "If true, copies raw data to the destination dataset. Default is false (raw data is stripped/not copied).", "schema": {"description": "If true, copies raw data to the destination dataset. Default is false (raw data is stripped/not copied).", "type": "boolean"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"destDatasetId": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}, "eventIds": {"type": "array", "items": {"example": "7632a037-fdef-4899-9b12-148470aae772", "type": "string"}}}, "required": ["destDatasetId", "eventIds"]}}}}, "responses": {"200": {"description": "Returns the number of copied events", "content": {"application/json": {"schema": {"type": "object", "properties": {"copied": {"description": "Number of events successfully copied", "type": "number", "example": 2}, "indicatorsCopied": {"description": "Number of indicators successfully copied", "type": "number", "example": 5}, "insertFailures": {"description": "Array of events that failed to insert into destination", "type": "array", "items": {"properties": {"index": {"description": "Index of the event that failed to insert", "type": "number"}, "reason": {"description": "Reason for the failure", "type": "string"}}, "required": ["index", "reason"], "type": "object"}}, "relationshipsCopied": {"description": "Number of relationships successfully copied", "type": "number", "example": 3}}, "required": ["copied", "indicatorsCopied", "relationshipsCopied"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
