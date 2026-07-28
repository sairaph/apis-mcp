---
title: Download quote PDF
page_id: operation-get-v1-quotes-quote-pdf-970568d1
path: operations/untagged
description: <p>Download the PDF for a finalized quote. Explanation for special handling can be found <a href="https://docs.stripe.com/quotes/overview#quote_pdf">here</a></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/quotes/{quote}/pdf
operation_ids:
    - GetQuotesQuotePdf
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Download quote PDF

`GET /v1/quotes/{quote}/pdf`

Operation ID: `GetQuotesQuotePdf`

<p>Download the PDF for a finalized quote. Explanation for special handling can be found <a href="https://docs.stripe.com/quotes/overview#quote_pdf">here</a></p>

## Definition

```yaml
{"summary": "Download quote PDF", "description": "<p>Download the PDF for a finalized quote. Explanation for special handling can be found <a href=\"https://docs.stripe.com/quotes/overview#quote_pdf\">here</a></p>", "operationId": "GetQuotesQuotePdf", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "quote", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/pdf": {"schema": {"type": "string", "format": "binary"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}, "servers": [{"url": "https://files.stripe.com/"}]}
```
