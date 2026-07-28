---
title: Reads data for a raw event
page_id: operation-get-accounts-account-id-cloudforce-one-events-event-id-raw-raw-id-ebd09573
path: operations/event
description: Retrieve raw data for a specific event.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/{event_id}/raw/{raw_id}
operation_ids:
    - get_EventRawRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reads data for a raw event

`GET /accounts/{account_id}/cloudforce-one/events/{event_id}/raw/{raw_id}`

Operation ID: `get_EventRawRead`

Retrieve raw data for a specific event.

## Definition

```yaml
{"operationId": "get_EventRawRead", "summary": "Reads data for a raw event", "description": "Retrieve raw data for a specific event.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}, {"name": "raw_id", "in": "path", "description": "Raw Event UUID.", "required": true, "schema": {"description": "Raw Event UUID.", "type": "string"}}], "responses": {"200": {"description": "Returns the raw event.", "content": {"application/json": {"schema": {"type": "object", "properties": {"accountId": {"type": "number", "example": 1234}, "created": {"type": "string", "example": "1970-01-01"}, "data": {"type": "object"}, "id": {"type": "string", "example": "1234"}, "source": {"type": "string", "example": "https://example.com"}, "tlp": {"type": "string", "example": "amber"}}, "required": ["accountId", "created", "data", "id", "source", "tlp"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
