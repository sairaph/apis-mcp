---
title: Finalize a quote
page_id: operation-post-v1-quotes-quote-finalize-9c392cf0
path: operations/untagged
description: <p>Finalizes the quote.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/quotes/{quote}/finalize
operation_ids:
    - PostQuotesQuoteFinalize
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Finalize a quote

`POST /v1/quotes/{quote}/finalize`

Operation ID: `PostQuotesQuoteFinalize`

<p>Finalizes the quote.</p>

## Definition

```yaml
{"summary": "Finalize a quote", "description": "<p>Finalizes the quote.</p>", "operationId": "PostQuotesQuoteFinalize", "parameters": [{"name": "quote", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"type": "integer", "description": "A future timestamp on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch.", "format": "unix-time"}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/quote"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
