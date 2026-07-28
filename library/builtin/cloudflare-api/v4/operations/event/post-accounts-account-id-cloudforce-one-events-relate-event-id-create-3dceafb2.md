---
title: Creates event references for a event
page_id: operation-post-accounts-account-id-cloudforce-one-events-relate-event-id-create-e8ee6e82
path: operations/event
description: Create one or more references between events.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/relate/{event_id}/create
operation_ids:
    - post_EventReferenceCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates event references for a event

`POST /accounts/{account_id}/cloudforce-one/events/relate/{event_id}/create`

Operation ID: `post_EventReferenceCreate`

Create one or more references between events.

## Definition

```yaml
{"operationId": "post_EventReferenceCreate", "summary": "Creates event references for a event", "description": "Create one or more references between events.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"events": {"type": "array", "items": {"example": "88888888-4444-4444-4444-121212121212", "type": "string"}}}, "required": ["events"]}}}}, "responses": {"200": {"description": "Returns success if operation succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"success": {"type": "boolean", "example": true}}, "required": ["success"]}, "success": {"type": "boolean", "example": true}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
