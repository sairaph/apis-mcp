---
title: Disable an Event Destination
page_id: operation-post-v2-core-event-destinations-id-disable-eba92f69
path: operations/untagged
description: Disable an event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/core/event_destinations/{id}/disable
operation_ids:
    - PostV2CoreEventDestinationsIdDisable
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Disable an Event Destination

`POST /v2/core/event_destinations/{id}/disable`

Operation ID: `PostV2CoreEventDestinationsIdDisable`

Disable an event destination.

## Definition

```yaml
{"summary": "Disable an Event Destination", "description": "Disable an event destination.", "operationId": "PostV2CoreEventDestinationsIdDisable", "parameters": [{"name": "id", "in": "path", "description": "Identifier for the event destination to disable.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event_destination"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.idempotency_error"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
