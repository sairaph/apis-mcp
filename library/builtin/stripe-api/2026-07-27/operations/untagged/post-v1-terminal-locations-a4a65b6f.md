---
title: Create a Location
page_id: operation-post-v1-terminal-locations-6fbf4e23
path: operations/untagged
description: |-
    <p>Creates a new <code>Location</code> object.
    For further details, including which address fields are required in each country, see the <a href="/docs/terminal/fleet/locations">Manage locations</a> guide.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/locations
operation_ids:
    - PostTerminalLocations
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Location

`POST /v1/terminal/locations`

Operation ID: `PostTerminalLocations`

<p>Creates a new <code>Location</code> object.
For further details, including which address fields are required in each country, see the <a href="/docs/terminal/fleet/locations">Manage locations</a> guide.</p>

## Definition

```yaml
{"summary": "Create a Location", "description": "<p>Creates a new <code>Location</code> object.\nFor further details, including which address fields are required in each country, see the <a href=\"/docs/terminal/fleet/locations\">Manage locations</a> guide.</p>", "operationId": "PostTerminalLocations", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"address": {"title": "create_location_address_param", "required": ["country"], "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}, "description": "The full address of the location."}, "address_kana": {"title": "japan_address_kana_specs", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}, "town": {"maxLength": 5000, "type": "string"}}, "description": "The Kana variation of the full address of the location (Japan only)."}, "address_kanji": {"title": "japan_address_kanji_specs", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}, "town": {"maxLength": 5000, "type": "string"}}, "description": "The Kanji variation of the full address of the location (Japan only)."}, "configuration_overrides": {"maxLength": 500, "type": "string", "description": "The ID of a configuration that will be used to customize all readers in this location."}, "display_name": {"maxLength": 1000, "type": "string", "description": "A name for the location. Maximum length is 1000 characters."}, "display_name_kana": {"maxLength": 1000, "type": "string", "description": "The Kana variation of the name for the location (Japan only). Maximum length is 1000 characters."}, "display_name_kanji": {"maxLength": 1000, "type": "string", "description": "The Kanji variation of the name for the location (Japan only). Maximum length is 1000 characters."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "phone": {"type": "string", "description": "The phone number for the location."}}, "additionalProperties": false}, "encoding": {"address": {"style": "deepObject", "explode": true}, "address_kana": {"style": "deepObject", "explode": true}, "address_kanji": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.location"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
