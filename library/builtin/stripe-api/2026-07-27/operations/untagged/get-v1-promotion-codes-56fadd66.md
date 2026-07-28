---
title: List all promotion codes
page_id: operation-get-v1-promotion-codes-988d668a
path: operations/untagged
description: <p>Returns a list of your promotion codes.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/promotion_codes
operation_ids:
    - GetPromotionCodes
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List all promotion codes

`GET /v1/promotion_codes`

Operation ID: `GetPromotionCodes`

<p>Returns a list of your promotion codes.</p>

## Definition

```yaml
{"summary": "List all promotion codes", "description": "<p>Returns a list of your promotion codes.</p>", "operationId": "GetPromotionCodes", "parameters": [{"name": "active", "in": "query", "description": "Filter promotion codes by whether they are active.", "required": false, "style": "form", "explode": true, "schema": {"type": "boolean"}}, {"name": "code", "in": "query", "description": "Only return promotion codes that have this case-insensitive code.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "coupon", "in": "query", "description": "Only return promotion codes for this coupon.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "created", "in": "query", "description": "A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.", "required": false, "style": "deepObject", "explode": true, "schema": {"anyOf": [{"title": "range_query_specs", "type": "object", "properties": {"gt": {"type": "integer"}, "gte": {"type": "integer"}, "lt": {"type": "integer"}, "lte": {"type": "integer"}}}, {"type": "integer"}]}}, {"name": "customer", "in": "query", "description": "Only return promotion codes that are restricted to this customer.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "customer_account", "in": "query", "description": "Only return promotion codes that are restricted to this account representing the customer.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "ending_before", "in": "query", "description": "A cursor for use in pagination. `ending_before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with `obj_bar`, your subsequent call can include `ending_before=obj_bar` in order to fetch the previous page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "limit", "in": "query", "description": "A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "starting_after", "in": "query", "description": "A cursor for use in pagination. `starting_after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with `obj_foo`, your subsequent call can include `starting_after=obj_foo` in order to fetch the next page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "PromotionCodesResourcePromotionCodeList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "items": {"$ref": "#/components/schemas/promotion_code"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "pattern": "^/v1/promotion_codes", "type": "string", "description": "The URL where this list can be accessed."}}, "description": "", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
