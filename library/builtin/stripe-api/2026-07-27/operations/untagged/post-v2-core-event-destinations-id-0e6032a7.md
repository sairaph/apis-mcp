---
title: Update an Event Destination
page_id: operation-post-v2-core-event-destinations-id-5c11372c
path: operations/untagged
description: Update the details of an event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/core/event_destinations/{id}
operation_ids:
    - PostV2CoreEventDestinationsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update an Event Destination

`POST /v2/core/event_destinations/{id}`

Operation ID: `PostV2CoreEventDestinationsId`

Update the details of an event destination.

## Definition

```yaml
{"summary": "Update an Event Destination", "description": "Update the details of an event destination.", "operationId": "PostV2CoreEventDestinationsId", "parameters": [{"name": "id", "in": "path", "description": "Identifier for the event destination to update.", "required": true, "style": "simple", "schema": {"type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "description": "An optional description of what the event destination is used for."}, "enabled_events": {"type": "array", "description": "The list of events to enable for this endpoint.", "items": {"type": "string"}}, "include": {"type": "array", "description": "Additional fields to include in the response. Currently supports `webhook_endpoint.url`.", "items": {"type": "string", "enum": ["webhook_endpoint.url"]}}, "metadata": {"type": "object", "additionalProperties": {"type": "string", "nullable": true}, "description": "Metadata."}, "name": {"type": "string", "description": "Event destination name."}, "webhook_endpoint": {"required": ["url"], "type": "object", "properties": {"url": {"type": "string", "description": "The URL of the webhook endpoint."}}, "description": "Webhook endpoint configuration."}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event_destination"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.idempotency_error"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
