---
title: payment_method_domain
page_id: schema-payment-method-domain-e2e4c27c
path: schemas
description: |-
    A payment method domain represents a web domain that you have registered with Stripe.
    Stripe Elements use registered payment method domains to control where certain payment methods are shown.

    Related guide: [Payment method domains](https://docs.stripe.com/payments/payment-methods/pmd-registration).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_domain

A payment method domain represents a web domain that you have registered with Stripe.
Stripe Elements use registered payment method domains to control where certain payment methods are shown.

Related guide: [Payment method domains](https://docs.stripe.com/payments/payment-methods/pmd-registration).

```yaml
{"title": "PaymentMethodDomainResourcePaymentMethodDomain", "required": ["amazon_pay", "apple_pay", "created", "domain_name", "enabled", "google_pay", "id", "klarna", "link", "livemode", "object", "paypal"], "type": "object", "properties": {"amazon_pay": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status"}, "apple_pay": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status"}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "domain_name": {"maxLength": 5000, "type": "string", "description": "The domain name that this payment method domain object represents."}, "enabled": {"type": "boolean", "description": "Whether this payment method domain is enabled. If the domain is not enabled, payment methods that require a payment method domain will not appear in Elements."}, "google_pay": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "klarna": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status"}, "link": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status"}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["payment_method_domain"]}, "paypal": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status"}}, "description": "A payment method domain represents a web domain that you have registered with Stripe.\nStripe Elements use registered payment method domains to control where certain payment methods are shown.\n\nRelated guide: [Payment method domains](https://docs.stripe.com/payments/payment-methods/pmd-registration).", "x-expandableFields": ["amazon_pay", "apple_pay", "google_pay", "klarna", "link", "paypal"], "x-resourceId": "payment_method_domain"}
```
