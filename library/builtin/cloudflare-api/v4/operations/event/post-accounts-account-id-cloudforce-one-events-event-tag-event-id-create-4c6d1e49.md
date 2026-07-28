---
title: Adds a tag to an event
page_id: operation-post-accounts-account-id-cloudforce-one-events-event-tag-event-id-create-3c8e7e06
path: operations/event
description: Add one or more tags to an event.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/event_tag/{event_id}/create
operation_ids:
    - post_EventTagCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Adds a tag to an event

`POST /accounts/{account_id}/cloudforce-one/events/event_tag/{event_id}/create`

Operation ID: `post_EventTagCreate`

Add one or more tags to an event.

## Definition

```yaml
{"operationId": "post_EventTagCreate", "summary": "Adds a tag to an event", "description": "Add one or more tags to an event.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"tags": {"type": "array", "items": {"example": "botnet", "type": "string"}}}, "required": ["tags"]}}}}, "responses": {"200": {"description": "Returns success if operation succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"success": {"type": "boolean", "example": true}}, "required": ["success"]}, "success": {"type": "boolean", "example": true}}, "required": ["success", "result"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
