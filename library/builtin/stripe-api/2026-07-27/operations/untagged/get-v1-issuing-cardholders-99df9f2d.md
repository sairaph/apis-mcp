---
title: List all cardholders
page_id: operation-get-v1-issuing-cardholders-e3e461dc
path: operations/untagged
description: <p>Returns a list of Issuing <code>Cardholder</code> objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/issuing/cardholders
operation_ids:
    - GetIssuingCardholders
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List all cardholders

`GET /v1/issuing/cardholders`

Operation ID: `GetIssuingCardholders`

<p>Returns a list of Issuing <code>Cardholder</code> objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.</p>

## Definition

```yaml
{"summary": "List all cardholders", "description": "<p>Returns a list of Issuing <code>Cardholder</code> objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.</p>", "operationId": "GetIssuingCardholders", "parameters": [{"name": "created", "in": "query", "description": "Only return cardholders that were created during the given date interval.", "required": false, "style": "deepObject", "explode": true, "schema": {"anyOf": [{"title": "range_query_specs", "type": "object", "properties": {"gt": {"type": "integer"}, "gte": {"type": "integer"}, "lt": {"type": "integer"}, "lte": {"type": "integer"}}}, {"type": "integer"}]}}, {"name": "email", "in": "query", "description": "Only return cardholders that have the given email address.", "required": false, "style": "form", "explode": true, "schema": {"type": "string"}}, {"name": "ending_before", "in": "query", "description": "A cursor for use in pagination. `ending_before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with `obj_bar`, your subsequent call can include `ending_before=obj_bar` in order to fetch the previous page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "limit", "in": "query", "description": "A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "phone_number", "in": "query", "description": "Only return cardholders that have the given phone number.", "required": false, "style": "form", "explode": true, "schema": {"type": "string"}}, {"name": "starting_after", "in": "query", "description": "A cursor for use in pagination. `starting_after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with `obj_foo`, your subsequent call can include `starting_after=obj_foo` in order to fetch the next page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "status", "in": "query", "description": "Only return cardholders that have the given status. One of `active`, `inactive`, or `blocked`.", "required": false, "style": "form", "explode": true, "schema": {"type": "string", "enum": ["active", "blocked", "inactive"]}}, {"name": "type", "in": "query", "description": "Only return cardholders that have the given type. One of `individual` or `company`.", "required": false, "style": "form", "explode": true, "schema": {"type": "string", "enum": ["company", "individual"], "x-stripeBypassValidation": true}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "IssuingCardholderList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "items": {"$ref": "#/components/schemas/issuing.cardholder"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "pattern": "^/v1/issuing/cardholders", "type": "string", "description": "The URL where this list can be accessed."}}, "description": "", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
