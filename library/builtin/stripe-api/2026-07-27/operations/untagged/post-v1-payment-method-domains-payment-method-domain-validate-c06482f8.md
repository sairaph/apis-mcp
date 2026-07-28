---
title: Validate an existing payment method domain
page_id: operation-post-v1-payment-method-domains-payment-method-domain-validate-580d4600
path: operations/untagged
description: |-
    <p>Some payment methods might require additional steps to register a domain. If the requirements weren’t satisfied when the domain was created, the payment method will be inactive on the domain.
    The payment method doesn’t appear in Elements or Embedded Checkout for this domain until it is active.</p>

    <p>To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.</p>

    <p>Related guides: <a href="/docs/payments/payment-methods/pmd-registration">Payment method domains</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_method_domains/{payment_method_domain}/validate
operation_ids:
    - PostPaymentMethodDomainsPaymentMethodDomainValidate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Validate an existing payment method domain

`POST /v1/payment_method_domains/{payment_method_domain}/validate`

Operation ID: `PostPaymentMethodDomainsPaymentMethodDomainValidate`

<p>Some payment methods might require additional steps to register a domain. If the requirements weren’t satisfied when the domain was created, the payment method will be inactive on the domain.
The payment method doesn’t appear in Elements or Embedded Checkout for this domain until it is active.</p>

<p>To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.</p>

<p>Related guides: <a href="/docs/payments/payment-methods/pmd-registration">Payment method domains</a>.</p>

## Definition

```yaml
{"summary": "Validate an existing payment method domain", "description": "<p>Some payment methods might require additional steps to register a domain. If the requirements weren’t satisfied when the domain was created, the payment method will be inactive on the domain.\nThe payment method doesn’t appear in Elements or Embedded Checkout for this domain until it is active.</p>\n\n<p>To activate a payment method on an existing payment method domain, complete the required registration steps specific to the payment method, and then validate the payment method domain with this endpoint.</p>\n\n<p>Related guides: <a href=\"/docs/payments/payment-methods/pmd-registration\">Payment method domains</a>.</p>", "operationId": "PostPaymentMethodDomainsPaymentMethodDomainValidate", "parameters": [{"name": "payment_method_domain", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
