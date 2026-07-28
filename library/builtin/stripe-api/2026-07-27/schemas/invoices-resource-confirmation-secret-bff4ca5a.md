---
title: invoices_resource_confirmation_secret
page_id: schema-invoices-resource-confirmation-secret-bff4ca5a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_resource_confirmation_secret

```yaml
{"title": "InvoicesResourceConfirmationSecret", "required": ["client_secret", "type"], "type": "object", "properties": {"client_secret": {"maxLength": 5000, "type": "string", "description": "The client_secret of the payment that Stripe creates for the invoice after finalization."}, "type": {"maxLength": 5000, "type": "string", "description": "The type of client_secret. Currently this is always payment_intent, referencing the default payment_intent that Stripe creates during invoice finalization"}}, "description": "", "x-expandableFields": []}
```
