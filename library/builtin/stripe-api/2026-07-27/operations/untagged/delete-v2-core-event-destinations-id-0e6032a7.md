---
title: Delete an Event Destination
page_id: operation-delete-v2-core-event-destinations-id-433c3454
path: operations/untagged
description: Delete an event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v2/core/event_destinations/{id}
operation_ids:
    - DeleteV2CoreEventDestinationsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete an Event Destination

`DELETE /v2/core/event_destinations/{id}`

Operation ID: `DeleteV2CoreEventDestinationsId`

Delete an event destination.

## Definition

```yaml
{"summary": "Delete an Event Destination", "description": "Delete an event destination.", "operationId": "DeleteV2CoreEventDestinationsId", "parameters": [{"name": "id", "in": "path", "description": "Identifier for the event destination to delete.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.deleted_object"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.idempotency_error"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
