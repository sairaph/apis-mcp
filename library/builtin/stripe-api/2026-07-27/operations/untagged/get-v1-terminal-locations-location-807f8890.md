---
title: Retrieve a Location
page_id: operation-get-v1-terminal-locations-location-ad316697
path: operations/untagged
description: <p>Retrieves a <code>Location</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/terminal/locations/{location}
operation_ids:
    - GetTerminalLocationsLocation
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Location

`GET /v1/terminal/locations/{location}`

Operation ID: `GetTerminalLocationsLocation`

<p>Retrieves a <code>Location</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a Location", "description": "<p>Retrieves a <code>Location</code> object.</p>", "operationId": "GetTerminalLocationsLocation", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "location", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/terminal.location"}, {"$ref": "#/components/schemas/deleted_terminal.location"}]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
