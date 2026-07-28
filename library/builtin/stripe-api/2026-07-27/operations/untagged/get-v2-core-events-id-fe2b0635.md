---
title: Retrieve an Event
page_id: operation-get-v2-core-events-id-be679705
path: operations/untagged
description: |-
    Retrieves the details of an event if it was created in the last 30 days. Supply the unique
    identifier of the event, which might have been delivered to your event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/events/{id}
operation_ids:
    - GetV2CoreEventsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an Event

`GET /v2/core/events/{id}`

Operation ID: `GetV2CoreEventsId`

Retrieves the details of an event if it was created in the last 30 days. Supply the unique
identifier of the event, which might have been delivered to your event destination.

## Definition

```yaml
{"summary": "Retrieve an Event", "description": "Retrieves the details of an event if it was created in the last 30 days. Supply the unique\nidentifier of the event, which might have been delivered to your event destination.", "operationId": "GetV2CoreEventsId", "parameters": [{"name": "id", "in": "path", "description": "Unique identifier for the object.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
