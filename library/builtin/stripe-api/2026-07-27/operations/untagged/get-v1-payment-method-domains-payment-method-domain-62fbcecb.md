---
title: Retrieve a payment method domain
page_id: operation-get-v1-payment-method-domains-payment-method-domain-3ab1369b
path: operations/untagged
description: <p>Retrieves the details of an existing payment method domain.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/payment_method_domains/{payment_method_domain}
operation_ids:
    - GetPaymentMethodDomainsPaymentMethodDomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a payment method domain

`GET /v1/payment_method_domains/{payment_method_domain}`

Operation ID: `GetPaymentMethodDomainsPaymentMethodDomain`

<p>Retrieves the details of an existing payment method domain.</p>

## Definition

```yaml
{"summary": "Retrieve a payment method domain", "description": "<p>Retrieves the details of an existing payment method domain.</p>", "operationId": "GetPaymentMethodDomainsPaymentMethodDomain", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "payment_method_domain", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
