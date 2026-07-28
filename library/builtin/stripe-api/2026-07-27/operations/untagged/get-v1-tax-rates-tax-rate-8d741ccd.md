---
title: Retrieve a tax rate
page_id: operation-get-v1-tax-rates-tax-rate-68f7b60b
path: operations/untagged
description: <p>Retrieves a tax rate with the given ID</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/tax_rates/{tax_rate}
operation_ids:
    - GetTaxRatesTaxRate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a tax rate

`GET /v1/tax_rates/{tax_rate}`

Operation ID: `GetTaxRatesTaxRate`

<p>Retrieves a tax rate with the given ID</p>

## Definition

```yaml
{"summary": "Retrieve a tax rate", "description": "<p>Retrieves a tax rate with the given ID</p>", "operationId": "GetTaxRatesTaxRate", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "tax_rate", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax_rate"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
