---
title: Cancel a quote
page_id: operation-post-v1-quotes-quote-cancel-794bb5dc
path: operations/untagged
description: <p>Cancels the quote.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/quotes/{quote}/cancel
operation_ids:
    - PostQuotesQuoteCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel a quote

`POST /v1/quotes/{quote}/cancel`

Operation ID: `PostQuotesQuoteCancel`

<p>Cancels the quote.</p>

## Definition

```yaml
{"summary": "Cancel a quote", "description": "<p>Cancels the quote.</p>", "operationId": "PostQuotesQuoteCancel", "parameters": [{"name": "quote", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/quote"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
