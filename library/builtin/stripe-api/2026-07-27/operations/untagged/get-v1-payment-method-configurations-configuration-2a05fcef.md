---
title: Retrieve payment method configuration
page_id: operation-get-v1-payment-method-configurations-configuration-eb75939e
path: operations/untagged
description: <p>Retrieve payment method configuration</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/payment_method_configurations/{configuration}
operation_ids:
    - GetPaymentMethodConfigurationsConfiguration
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve payment method configuration

`GET /v1/payment_method_configurations/{configuration}`

Operation ID: `GetPaymentMethodConfigurationsConfiguration`

<p>Retrieve payment method configuration</p>

## Definition

```yaml
{"summary": "Retrieve payment method configuration", "description": "<p>Retrieve payment method configuration</p>", "operationId": "GetPaymentMethodConfigurationsConfiguration", "parameters": [{"name": "configuration", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method_configuration"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
