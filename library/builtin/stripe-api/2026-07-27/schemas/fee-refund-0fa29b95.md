---
title: fee_refund
page_id: schema-fee-refund-0fa29b95
path: schemas
description: |-
    `Application Fee Refund` objects allow you to refund an application fee that
    has previously been created but not yet refunded. Funds will be refunded to
    the Stripe account from which the fee was originally collected.

    Related guide: [Refunding application fees](https://docs.stripe.com/connect/destination-charges#refunding-app-fee)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# fee_refund

`Application Fee Refund` objects allow you to refund an application fee that
has previously been created but not yet refunded. Funds will be refunded to
the Stripe account from which the fee was originally collected.

Related guide: [Refunding application fees](https://docs.stripe.com/connect/destination-charges#refunding-app-fee)

```yaml
{"title": "FeeRefund", "required": ["amount", "created", "currency", "fee", "id", "object"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount, in cents (or local equivalent)."}, "balance_transaction": {"description": "Balance transaction that describes the impact on your account balance.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/balance_transaction"}]}}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "fee": {"description": "ID of the application fee that was refunded.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/application_fee"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/application_fee"}]}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["fee_refund"]}}, "description": "`Application Fee Refund` objects allow you to refund an application fee that\nhas previously been created but not yet refunded. Funds will be refunded to\nthe Stripe account from which the fee was originally collected.\n\nRelated guide: [Refunding application fees](https://docs.stripe.com/connect/destination-charges#refunding-app-fee)", "x-expandableFields": ["balance_transaction", "fee"], "x-resourceId": "fee_refund"}
```
