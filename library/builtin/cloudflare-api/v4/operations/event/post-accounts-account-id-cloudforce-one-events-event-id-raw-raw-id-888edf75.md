---
title: Updates a raw event
page_id: operation-post-accounts-account-id-cloudforce-one-events-event-id-raw-raw-id-b85c43cd
path: operations/event
description: Update raw data for a specific event.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/{event_id}/raw/{raw_id}
operation_ids:
    - post_EventRawUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates a raw event

`POST /accounts/{account_id}/cloudforce-one/events/{event_id}/raw/{raw_id}`

Operation ID: `post_EventRawUpdate`

Update raw data for a specific event.

## Definition

```yaml
{"operationId": "post_EventRawUpdate", "summary": "Updates a raw event", "description": "Update raw data for a specific event.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}, {"name": "raw_id", "in": "path", "description": "Raw Event UUID.", "required": true, "schema": {"description": "Raw Event UUID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object"}, "source": {"type": "string", "example": "example.com"}, "tlp": {"type": "string", "example": "amber"}}}}}}, "responses": {"200": {"description": "Returns the uuid of the updated raw event and its data.", "content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "object"}, "id": {"type": "string", "example": "1234"}}, "required": ["id", "data"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
