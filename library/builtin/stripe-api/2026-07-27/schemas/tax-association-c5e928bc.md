---
title: tax.association
page_id: schema-tax-association-c5e928bc
path: schemas
description: A Tax Association exposes the Tax Transactions that Stripe attempted to create on your behalf based on the PaymentIntent input
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax.association

A Tax Association exposes the Tax Transactions that Stripe attempted to create on your behalf based on the PaymentIntent input

```yaml
{"title": "TaxProductResourceTaxAssociation", "required": ["calculation", "id", "object", "payment_intent"], "type": "object", "properties": {"calculation": {"maxLength": 5000, "type": "string", "description": "The [Tax Calculation](https://docs.stripe.com/api/tax/calculations/object) that was included in PaymentIntent."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["tax.association"]}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The [PaymentIntent](https://docs.stripe.com/api/payment_intents/object) that this Tax Association is tracking."}, "tax_transaction_attempts": {"type": "array", "description": "Information about the tax transactions linked to this payment intent", "nullable": true, "items": {"$ref": "#/components/schemas/tax_product_resource_tax_association_transaction_attempts"}}}, "description": "A Tax Association exposes the Tax Transactions that Stripe attempted to create on your behalf based on the PaymentIntent input", "x-expandableFields": ["tax_transaction_attempts"], "x-resourceId": "tax.association"}
```
