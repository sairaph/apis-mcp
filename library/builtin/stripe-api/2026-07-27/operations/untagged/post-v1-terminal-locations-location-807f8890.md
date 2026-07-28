---
title: Update a Location
page_id: operation-post-v1-terminal-locations-location-f216402b
path: operations/untagged
description: <p>Updates a <code>Location</code> object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/locations/{location}
operation_ids:
    - PostTerminalLocationsLocation
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a Location

`POST /v1/terminal/locations/{location}`

Operation ID: `PostTerminalLocationsLocation`

<p>Updates a <code>Location</code> object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.</p>

## Definition

```yaml
{"summary": "Update a Location", "description": "<p>Updates a <code>Location</code> object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.</p>", "operationId": "PostTerminalLocationsLocation", "parameters": [{"name": "location", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"address": {"title": "optional_fields_address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}, "description": "The full address of the location. You can't change the location's `country`. If you need to modify the `country` field, create a new `Location` object and re-register any existing readers to that location."}, "address_kana": {"title": "japan_address_kana_specs", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}, "town": {"maxLength": 5000, "type": "string"}}, "description": "The Kana variation of the full address of the location (Japan only)."}, "address_kanji": {"title": "japan_address_kanji_specs", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}, "town": {"maxLength": 5000, "type": "string"}}, "description": "The Kanji variation of the full address of the location (Japan only)."}, "configuration_overrides": {"description": "The ID of a configuration that will be used to customize all readers in this location.", "anyOf": [{"maxLength": 1000, "type": "string"}, {"type": "string", "enum": [""]}]}, "display_name": {"description": "A name for the location.", "anyOf": [{"maxLength": 1000, "type": "string"}, {"type": "string", "enum": [""]}]}, "display_name_kana": {"description": "The Kana variation of the name for the location (Japan only).", "anyOf": [{"maxLength": 1000, "type": "string"}, {"type": "string", "enum": [""]}]}, "display_name_kanji": {"description": "The Kanji variation of the name for the location (Japan only).", "anyOf": [{"maxLength": 1000, "type": "string"}, {"type": "string", "enum": [""]}]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "phone": {"description": "The phone number for the location.", "anyOf": [{"type": "string"}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"address": {"style": "deepObject", "explode": true}, "address_kana": {"style": "deepObject", "explode": true}, "address_kanji": {"style": "deepObject", "explode": true}, "configuration_overrides": {"style": "deepObject", "explode": true}, "display_name": {"style": "deepObject", "explode": true}, "display_name_kana": {"style": "deepObject", "explode": true}, "display_name_kanji": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "phone": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/terminal.location"}, {"$ref": "#/components/schemas/deleted_terminal.location"}]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
