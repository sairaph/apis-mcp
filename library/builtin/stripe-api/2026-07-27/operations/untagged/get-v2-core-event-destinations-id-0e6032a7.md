---
title: Retrieve an Event Destination
page_id: operation-get-v2-core-event-destinations-id-9af8a9a0
path: operations/untagged
description: Retrieves the details of an event destination.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/event_destinations/{id}
operation_ids:
    - GetV2CoreEventDestinationsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an Event Destination

`GET /v2/core/event_destinations/{id}`

Operation ID: `GetV2CoreEventDestinationsId`

Retrieves the details of an event destination.

## Definition

```yaml
{"summary": "Retrieve an Event Destination", "description": "Retrieves the details of an event destination.", "operationId": "GetV2CoreEventDestinationsId", "parameters": [{"name": "id", "in": "path", "description": "Identifier for the event destination to retrieve.", "required": true, "style": "simple", "schema": {"type": "string"}}, {"name": "include", "in": "query", "description": "Additional fields to include in the response.", "required": false, "style": "deepObject", "schema": {"type": "array", "items": {"type": "string", "enum": ["webhook_endpoint.url"]}}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.event_destination"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
