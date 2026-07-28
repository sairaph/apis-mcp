---
title: Accept a quote
page_id: operation-post-v1-quotes-quote-accept-6186d996
path: operations/untagged
description: <p>Accepts the specified quote.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/quotes/{quote}/accept
operation_ids:
    - PostQuotesQuoteAccept
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Accept a quote

`POST /v1/quotes/{quote}/accept`

Operation ID: `PostQuotesQuoteAccept`

<p>Accepts the specified quote.</p>

## Definition

```yaml
{"summary": "Accept a quote", "description": "<p>Accepts the specified quote.</p>", "operationId": "PostQuotesQuoteAccept", "parameters": [{"name": "quote", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/quote"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
