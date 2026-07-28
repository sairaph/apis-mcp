---
title: Enable an Event Destination
page_id: operation-post-v2-core-event-destinations-id-enable-fb7d2bd4
path: operations/untagged
description: Enable an event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/core/event_destinations/{id}/enable
operation_ids:
    - PostV2CoreEventDestinationsIdEnable
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Enable an Event Destination

`POST /v2/core/event_destinations/{id}/enable`

Operation ID: `PostV2CoreEventDestinationsIdEnable`

Enable an event destination.

## Definition

```yaml
{"summary": "Enable an Event Destination", "description": "Enable an event destination.", "operationId": "PostV2CoreEventDestinationsIdEnable", "parameters": [{"name": "id", "in": "path", "description": "Identifier for the event destination to enable.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event_destination"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.idempotency_error"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
