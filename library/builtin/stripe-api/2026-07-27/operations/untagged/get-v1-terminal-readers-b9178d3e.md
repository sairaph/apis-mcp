---
title: List all Readers
page_id: operation-get-v1-terminal-readers-30afd1f8
path: operations/untagged
description: <p>Returns a list of <code>Reader</code> objects.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/terminal/readers
operation_ids:
    - GetTerminalReaders
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List all Readers

`GET /v1/terminal/readers`

Operation ID: `GetTerminalReaders`

<p>Returns a list of <code>Reader</code> objects.</p>

## Definition

```yaml
{"summary": "List all Readers", "description": "<p>Returns a list of <code>Reader</code> objects.</p>", "operationId": "GetTerminalReaders", "parameters": [{"name": "device_type", "in": "query", "description": "Filters readers by device type", "required": false, "style": "form", "explode": true, "schema": {"type": "string", "enum": ["bbpos_chipper2x", "bbpos_wisepad3", "bbpos_wisepos_e", "mobile_phone_reader", "simulated_stripe_s700", "simulated_stripe_s710", "simulated_verifone_m425", "simulated_verifone_p630", "simulated_verifone_ux700", "simulated_verifone_v660p", "simulated_wisepos_e", "stripe_m2", "stripe_s700", "stripe_s710", "verifone_P400", "verifone_m425", "verifone_p630", "verifone_ux700", "verifone_v660p"], "x-stripeBypassValidation": true}}, {"name": "ending_before", "in": "query", "description": "A cursor for use in pagination. `ending_before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with `obj_bar`, your subsequent call can include `ending_before=obj_bar` in order to fetch the previous page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "limit", "in": "query", "description": "A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "location", "in": "query", "description": "A location ID to filter the response list to only readers at the specific location", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "serial_number", "in": "query", "description": "Filters readers by serial number", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "starting_after", "in": "query", "description": "A cursor for use in pagination. `starting_after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with `obj_foo`, your subsequent call can include `starting_after=obj_foo` in order to fetch the next page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "status", "in": "query", "description": "A status filter to filter readers to only offline or online readers", "required": false, "style": "form", "explode": true, "schema": {"type": "string", "enum": ["offline", "online"]}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "TerminalReaderRetrieveReader", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "A list of readers", "items": {"$ref": "#/components/schemas/terminal.reader"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
