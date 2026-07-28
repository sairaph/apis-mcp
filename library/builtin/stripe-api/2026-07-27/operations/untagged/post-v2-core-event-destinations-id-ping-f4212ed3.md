---
title: Ping an Event Destination
page_id: operation-post-v2-core-event-destinations-id-ping-29b9bfb3
path: operations/untagged
description: Send a `ping` event to an event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/core/event_destinations/{id}/ping
operation_ids:
    - PostV2CoreEventDestinationsIdPing
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Ping an Event Destination

`POST /v2/core/event_destinations/{id}/ping`

Operation ID: `PostV2CoreEventDestinationsIdPing`

Send a `ping` event to an event destination.

## Definition

```yaml
{"summary": "Ping an Event Destination", "description": "Send a `ping` event to an event destination.", "operationId": "PostV2CoreEventDestinationsIdPing", "parameters": [{"name": "id", "in": "path", "description": "Identifier for the event destination to ping.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
