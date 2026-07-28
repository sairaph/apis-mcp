---
title: List all ForwardingRequests
page_id: operation-get-v1-forwarding-requests-b003e3c4
path: operations/untagged
description: <p>Lists all ForwardingRequest objects.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/forwarding/requests
operation_ids:
    - GetForwardingRequests
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List all ForwardingRequests

`GET /v1/forwarding/requests`

Operation ID: `GetForwardingRequests`

<p>Lists all ForwardingRequest objects.</p>

## Definition

```yaml
{"summary": "List all ForwardingRequests", "description": "<p>Lists all ForwardingRequest objects.</p>", "operationId": "GetForwardingRequests", "parameters": [{"name": "created", "in": "query", "description": "Similar to other List endpoints, filters results based on created timestamp. You can pass gt, gte, lt, and lte timestamp values.", "required": false, "style": "deepObject", "explode": true, "schema": {"title": "created_param", "type": "object", "properties": {"gt": {"type": "integer"}, "gte": {"type": "integer"}, "lt": {"type": "integer"}, "lte": {"type": "integer"}}}}, {"name": "ending_before", "in": "query", "description": "A pagination cursor to fetch the previous page of the list. The value must be a ForwardingRequest ID.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "limit", "in": "query", "description": "A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "starting_after", "in": "query", "description": "A pagination cursor to fetch the next page of the list. The value must be a ForwardingRequest ID.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "ForwardingRequestList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "items": {"$ref": "#/components/schemas/forwarding.request"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "List of ForwardingRequest data.", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
