---
title: Removes an event reference
page_id: operation-delete-accounts-account-id-cloudforce-one-events-relate-event-id-e219cad5
path: operations/event
description: Remove one or more references from an event.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/relate/{event_id}
operation_ids:
    - delete_EventReferenceDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Removes an event reference

`DELETE /accounts/{account_id}/cloudforce-one/events/relate/{event_id}`

Operation ID: `delete_EventReferenceDelete`

Remove one or more references from an event.

## Definition

```yaml
{"operationId": "delete_EventReferenceDelete", "summary": "Removes an event reference", "description": "Remove one or more references from an event.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"events": {"type": "array", "items": {"example": "88888888-4444-4444-4444-121212121212", "type": "string"}}}, "required": ["events"]}}}}, "responses": {"200": {"description": "Returns success if operation succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"success": {"type": "boolean", "example": true}}, "required": ["success"]}, "success": {"type": "boolean", "example": true}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
