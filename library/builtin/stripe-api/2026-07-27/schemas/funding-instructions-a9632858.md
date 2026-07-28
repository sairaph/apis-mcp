---
title: funding_instructions
page_id: schema-funding-instructions-a9632858
path: schemas
description: |-
    Each customer has a [`balance`](https://docs.stripe.com/api/customers/object#customer_object-balance) that is
    automatically applied to future invoices and payments using the `customer_balance` payment method.
    Customers can fund this balance by initiating a bank transfer to any account in the
    `financial_addresses` field.
    Related guide: [Customer balance funding instructions](https://docs.stripe.com/payments/customer-balance/funding-instructions)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# funding_instructions

Each customer has a [`balance`](https://docs.stripe.com/api/customers/object#customer_object-balance) that is
automatically applied to future invoices and payments using the `customer_balance` payment method.
Customers can fund this balance by initiating a bank transfer to any account in the
`financial_addresses` field.
Related guide: [Customer balance funding instructions](https://docs.stripe.com/payments/customer-balance/funding-instructions)

```yaml
{"title": "CustomerBalanceFundingInstructionsCustomerBalanceFundingInstructions", "required": ["bank_transfer", "currency", "funding_type", "livemode", "object"], "type": "object", "properties": {"bank_transfer": {"$ref": "#/components/schemas/funding_instructions_bank_transfer"}, "currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies)."}, "funding_type": {"type": "string", "description": "The `funding_type` of the returned instructions", "enum": ["bank_transfer"]}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["funding_instructions"]}}, "description": "Each customer has a [`balance`](https://docs.stripe.com/api/customers/object#customer_object-balance) that is\nautomatically applied to future invoices and payments using the `customer_balance` payment method.\nCustomers can fund this balance by initiating a bank transfer to any account in the\n`financial_addresses` field.\nRelated guide: [Customer balance funding instructions](https://docs.stripe.com/payments/customer-balance/funding-instructions)", "x-expandableFields": ["bank_transfer"], "x-resourceId": "funding_instructions"}
```
