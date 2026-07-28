---
title: v2.error.attach_payment_method_to_customer
page_id: schema-v2-error-attach-payment-method-to-customer-57aadc34
path: schemas
description: Default payment method is added to the customer config before attaching it to the account using `/v1/payment_methods`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.attach_payment_method_to_customer

Default payment method is added to the customer config before attaching it to the account using `/v1/payment_methods`.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "doc_url", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["attach_payment_method_to_customer"]}, "doc_url": {"type": "string", "description": "A URL to more information about the error reported"}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "Default payment method is added to the customer config before attaching it to the account using `/v1/payment_methods`."}
```
