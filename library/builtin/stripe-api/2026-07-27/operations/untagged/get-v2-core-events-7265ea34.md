---
title: List Events
page_id: operation-get-v2-core-events-f479e953
path: operations/untagged
description: List events, going back up to 30 days.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/events
operation_ids:
    - GetV2CoreEvents
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List Events

`GET /v2/core/events`

Operation ID: `GetV2CoreEvents`

List events, going back up to 30 days.

## Definition

```yaml
{"summary": "List Events", "description": "List events, going back up to 30 days.", "operationId": "GetV2CoreEvents", "parameters": [{"name": "created", "in": "query", "description": "Set of filters to query events within a range of `created` timestamps.", "required": false, "style": "deepObject", "schema": {"type": "object", "properties": {"gt": {"type": "string", "description": "Filter for events created after the specified timestamp.", "format": "date-time"}, "gte": {"type": "string", "description": "Filter for events created at or after the specified timestamp.", "format": "date-time"}, "lt": {"type": "string", "description": "Filter for events created before the specified timestamp.", "format": "date-time"}, "lte": {"type": "string", "description": "Filter for events created at or before the specified timestamp.", "format": "date-time"}}}}, {"name": "limit", "in": "query", "description": "The page size.", "required": false, "style": "form", "schema": {"type": "integer"}}, {"name": "object_id", "in": "query", "description": "Primary object ID used to retrieve related events.", "required": false, "style": "form", "schema": {"type": "string"}}, {"name": "page", "in": "query", "description": "The requested page.", "required": false, "style": "form", "schema": {"type": "string"}}, {"name": "types", "in": "query", "description": "An array of up to 20 strings containing specific event names.", "required": false, "style": "deepObject", "schema": {"type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"required": ["data", "next_page_url", "previous_page_url"], "type": "object", "properties": {"data": {"type": "array", "description": "List of events.", "items": {"$ref": "#/components/schemas/v2.core.event"}}, "next_page_url": {"type": "string", "description": "URL to fetch the next page of the list. If there are no more pages, the value is null.", "nullable": true}, "previous_page_url": {"type": "string", "description": "URL to fetch the previous page of the list. If there are no previous pages, the value is null.", "nullable": true}}}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
