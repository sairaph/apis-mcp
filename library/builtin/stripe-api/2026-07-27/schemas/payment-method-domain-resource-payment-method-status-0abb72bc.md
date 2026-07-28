---
title: payment_method_domain_resource_payment_method_status
page_id: schema-payment-method-domain-resource-payment-method-status-0abb72bc
path: schemas
description: Indicates the status of a specific payment method on a payment method domain.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_domain_resource_payment_method_status

Indicates the status of a specific payment method on a payment method domain.

```yaml
{"title": "PaymentMethodDomainResourcePaymentMethodStatus", "required": ["status"], "type": "object", "properties": {"status": {"type": "string", "description": "The status of the payment method on the domain.", "enum": ["active", "inactive"]}, "status_details": {"$ref": "#/components/schemas/payment_method_domain_resource_payment_method_status_details"}}, "description": "Indicates the status of a specific payment method on a payment method domain.", "x-expandableFields": ["status_details"]}
```
