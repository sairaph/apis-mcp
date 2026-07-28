---
title: Removes a tag from an event
page_id: operation-delete-accounts-account-id-cloudforce-one-events-event-tag-event-id-1e94a400
path: operations/event
description: Remove one or more tags from an event.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/event_tag/{event_id}
operation_ids:
    - delete_EventTagDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Removes a tag from an event

`DELETE /accounts/{account_id}/cloudforce-one/events/event_tag/{event_id}`

Operation ID: `delete_EventTagDelete`

Remove one or more tags from an event.

## Definition

```yaml
{"operationId": "delete_EventTagDelete", "summary": "Removes a tag from an event", "description": "Remove one or more tags from an event.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"tags": {"type": "array", "items": {"example": "malware", "type": "string"}}}, "required": ["tags"]}}}}, "responses": {"200": {"description": "Returns success if operation succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"success": {"type": "boolean", "example": true}}, "required": ["success"]}, "success": {"type": "boolean", "example": true}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
