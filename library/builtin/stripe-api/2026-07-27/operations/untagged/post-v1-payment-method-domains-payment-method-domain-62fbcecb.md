---
title: Update a payment method domain
page_id: operation-post-v1-payment-method-domains-payment-method-domain-362f9cb6
path: operations/untagged
description: <p>Updates an existing payment method domain.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_method_domains/{payment_method_domain}
operation_ids:
    - PostPaymentMethodDomainsPaymentMethodDomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a payment method domain

`POST /v1/payment_method_domains/{payment_method_domain}`

Operation ID: `PostPaymentMethodDomainsPaymentMethodDomain`

<p>Updates an existing payment method domain.</p>

## Definition

```yaml
{"summary": "Update a payment method domain", "description": "<p>Updates an existing payment method domain.</p>", "operationId": "PostPaymentMethodDomainsPaymentMethodDomain", "parameters": [{"name": "payment_method_domain", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether this payment method domain is enabled. If the domain is not enabled, payment methods that require a payment method domain will not appear in Elements or Embedded Checkout."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
