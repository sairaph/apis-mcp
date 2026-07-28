---
title: Retrieve payment link
page_id: operation-get-v1-payment-links-payment-link-13685231
path: operations/untagged
description: <p>Retrieve a payment link.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/payment_links/{payment_link}
operation_ids:
    - GetPaymentLinksPaymentLink
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve payment link

`GET /v1/payment_links/{payment_link}`

Operation ID: `GetPaymentLinksPaymentLink`

<p>Retrieve a payment link.</p>

## Definition

```yaml
{"summary": "Retrieve payment link", "description": "<p>Retrieve a payment link.</p>", "operationId": "GetPaymentLinksPaymentLink", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "payment_link", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_link"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
