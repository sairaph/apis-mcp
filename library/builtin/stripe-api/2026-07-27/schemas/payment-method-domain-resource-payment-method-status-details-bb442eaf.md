---
title: payment_method_domain_resource_payment_method_status_details
page_id: schema-payment-method-domain-resource-payment-method-status-details-bb442eaf
path: schemas
description: Contains additional details about the status of a payment method for a specific payment method domain.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_domain_resource_payment_method_status_details

Contains additional details about the status of a payment method for a specific payment method domain.

```yaml
{"title": "PaymentMethodDomainResourcePaymentMethodStatusDetails", "required": ["error_message"], "type": "object", "properties": {"error_message": {"maxLength": 5000, "type": "string", "description": "The error message associated with the status of the payment method on the domain."}}, "description": "Contains additional details about the status of a payment method for a specific payment method domain.", "x-expandableFields": []}
```
