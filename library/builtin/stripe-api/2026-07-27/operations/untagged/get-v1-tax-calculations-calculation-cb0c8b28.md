---
title: Retrieve a Calculation
page_id: operation-get-v1-tax-calculations-calculation-4ab362f0
path: operations/untagged
description: <p>Retrieves a Tax <code>Calculation</code> object, if the calculation hasn’t expired.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/tax/calculations/{calculation}
operation_ids:
    - GetTaxCalculationsCalculation
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Calculation

`GET /v1/tax/calculations/{calculation}`

Operation ID: `GetTaxCalculationsCalculation`

<p>Retrieves a Tax <code>Calculation</code> object, if the calculation hasn’t expired.</p>

## Definition

```yaml
{"summary": "Retrieve a Calculation", "description": "<p>Retrieves a Tax <code>Calculation</code> object, if the calculation hasn’t expired.</p>", "operationId": "GetTaxCalculationsCalculation", "parameters": [{"name": "calculation", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.calculation"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
