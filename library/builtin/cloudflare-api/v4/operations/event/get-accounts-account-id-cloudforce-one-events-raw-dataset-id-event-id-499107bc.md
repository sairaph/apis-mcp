---
title: Reads raw data for an event by UUID
page_id: operation-get-accounts-account-id-cloudforce-one-events-raw-dataset-id-event-id-591bbf95
path: operations/event
description: Retrieves the raw data associated with an event. Searches across all shards in the dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/raw/{dataset_id}/{event_id}
operation_ids:
    - get_EventRawReadDS
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reads raw data for an event by UUID

`GET /accounts/{account_id}/cloudforce-one/events/raw/{dataset_id}/{event_id}`

Operation ID: `get_EventRawReadDS`

Retrieves the raw data associated with an event. Searches across all shards in the dataset.

## Definition

```yaml
{"operationId": "get_EventRawReadDS", "summary": "Reads raw data for an event by UUID", "description": "Retrieves the raw data associated with an event. Searches across all shards in the dataset.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event ID.", "required": true, "schema": {"description": "Event ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string"}}], "responses": {"200": {"description": "Returns the raw event data.", "content": {"application/json": {"schema": {"type": "object", "properties": {"accountId": {"type": "number", "example": 1234}, "created": {"type": "string", "example": "1970-01-01T00:00:00.000Z"}, "data": {"type": "string", "example": "{\"foo\": \"bar\"}"}, "id": {"type": "number", "example": 1}, "source": {"type": "string", "example": "https://example.com"}, "tlp": {"type": "string", "example": "amber"}}, "required": ["id", "accountId", "created", "data", "source", "tlp"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "500": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
