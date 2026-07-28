---
title: Create a payment method domain
page_id: operation-post-v1-payment-method-domains-a25a2cd4
path: operations/untagged
description: <p>Creates a payment method domain.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_method_domains
operation_ids:
    - PostPaymentMethodDomains
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a payment method domain

`POST /v1/payment_method_domains`

Operation ID: `PostPaymentMethodDomains`

<p>Creates a payment method domain.</p>

## Definition

```yaml
{"summary": "Create a payment method domain", "description": "<p>Creates a payment method domain.</p>", "operationId": "PostPaymentMethodDomains", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["domain_name"], "type": "object", "properties": {"domain_name": {"maxLength": 5000, "type": "string", "description": "The domain name that this payment method domain object represents."}, "enabled": {"type": "boolean", "description": "Whether this payment method domain is enabled. If the domain is not enabled, payment methods that require a payment method domain will not appear in Elements or Embedded Checkout."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
