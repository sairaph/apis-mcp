---
title: Retrieve an exchange rate
page_id: operation-get-v1-exchange-rates-rate-id-d2ba284e
path: operations/untagged
description: |-
    <p>[Deprecated] The <code>ExchangeRate</code> APIs are deprecated. Please use the <a href="https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api">FX Quotes API</a> instead.</p>

    <p>Retrieves the exchange rates from the given currency to every supported currency.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/exchange_rates/{rate_id}
operation_ids:
    - GetExchangeRatesRateId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an exchange rate

`GET /v1/exchange_rates/{rate_id}`

Operation ID: `GetExchangeRatesRateId`

<p>[Deprecated] The <code>ExchangeRate</code> APIs are deprecated. Please use the <a href="https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api">FX Quotes API</a> instead.</p>

<p>Retrieves the exchange rates from the given currency to every supported currency.</p>

## Definition

```yaml
{"summary": "Retrieve an exchange rate", "description": "<p>[Deprecated] The <code>ExchangeRate</code> APIs are deprecated. Please use the <a href=\"https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api\">FX Quotes API</a> instead.</p>\n\n<p>Retrieves the exchange rates from the given currency to every supported currency.</p>", "operationId": "GetExchangeRatesRateId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "rate_id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/exchange_rate"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
