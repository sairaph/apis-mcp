---
title: Delete a Location
page_id: operation-delete-v1-terminal-locations-location-297f5347
path: operations/untagged
description: <p>Deletes a <code>Location</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/terminal/locations/{location}
operation_ids:
    - DeleteTerminalLocationsLocation
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a Location

`DELETE /v1/terminal/locations/{location}`

Operation ID: `DeleteTerminalLocationsLocation`

<p>Deletes a <code>Location</code> object.</p>

## Definition

```yaml
{"summary": "Delete a Location", "description": "<p>Deletes a <code>Location</code> object.</p>", "operationId": "DeleteTerminalLocationsLocation", "parameters": [{"name": "location", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_terminal.location"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
