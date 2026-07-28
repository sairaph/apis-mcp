---
title: List Event Destinations
page_id: operation-get-v2-core-event-destinations-68c3d7f8
path: operations/untagged
description: Lists all event destinations.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/event_destinations
operation_ids:
    - GetV2CoreEventDestinations
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List Event Destinations

`GET /v2/core/event_destinations`

Operation ID: `GetV2CoreEventDestinations`

Lists all event destinations.

## Definition

```yaml
{"summary": "List Event Destinations", "description": "Lists all event destinations.", "operationId": "GetV2CoreEventDestinations", "parameters": [{"name": "include", "in": "query", "description": "Additional fields to include in the response. Currently supports `webhook_endpoint.url`.", "required": false, "style": "deepObject", "schema": {"type": "array", "items": {"type": "string", "enum": ["webhook_endpoint.url"]}}}, {"name": "limit", "in": "query", "description": "The page size.", "required": false, "style": "form", "schema": {"type": "integer"}}, {"name": "page", "in": "query", "description": "The requested page.", "required": false, "style": "form", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"required": ["data", "next_page_url", "previous_page_url"], "type": "object", "properties": {"data": {"type": "array", "description": "List of event destinations.", "items": {"$ref": "#/components/schemas/v2.core.event_destination"}}, "next_page_url": {"type": "string", "description": "URL to fetch the next page of the list. If there are no more pages, the value is null.", "nullable": true}, "previous_page_url": {"type": "string", "description": "URL to fetch the previous page of the list. If there are no previous pages, the value is null.", "nullable": true}}}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
