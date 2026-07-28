---
title: transfer_reversal
page_id: schema-transfer-reversal-d8bc5101
path: schemas
description: |-
    [Stripe Connect](https://docs.stripe.com/connect) platforms can reverse transfers made to a
    connected account, either entirely or partially, and can also specify whether
    to refund any related application fees. Transfer reversals add to the
    platform's balance and subtract from the destination account's balance.

    Reversing a transfer that was made for a [destination
    charge](/docs/connect/destination-charges) is allowed only up to the amount of
    the charge. It is possible to reverse a
    [transfer_group](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options)
    transfer only if the destination account has enough balance to cover the
    reversal.

    Related guide: [Reverse transfers](https://docs.stripe.com/connect/separate-charges-and-transfers#reverse-transfers)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# transfer_reversal

[Stripe Connect](https://docs.stripe.com/connect) platforms can reverse transfers made to a
connected account, either entirely or partially, and can also specify whether
to refund any related application fees. Transfer reversals add to the
platform's balance and subtract from the destination account's balance.

Reversing a transfer that was made for a [destination
charge](/docs/connect/destination-charges) is allowed only up to the amount of
the charge. It is possible to reverse a
[transfer_group](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options)
transfer only if the destination account has enough balance to cover the
reversal.

Related guide: [Reverse transfers](https://docs.stripe.com/connect/separate-charges-and-transfers#reverse-transfers)

```yaml
{"title": "TransferReversal", "required": ["amount", "created", "currency", "id", "object", "transfer"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount, in cents (or local equivalent)."}, "balance_transaction": {"description": "Balance transaction that describes the impact on your account balance.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/balance_transaction"}]}}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "destination_payment_refund": {"description": "Linked payment refund for the transfer reversal.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/refund"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/refund"}]}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["transfer_reversal"]}, "source_refund": {"description": "ID of the refund responsible for the transfer reversal.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/refund"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/refund"}]}}, "transfer": {"description": "ID of the transfer that was reversed.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/transfer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/transfer"}]}}}, "description": "[Stripe Connect](https://docs.stripe.com/connect) platforms can reverse transfers made to a\nconnected account, either entirely or partially, and can also specify whether\nto refund any related application fees. Transfer reversals add to the\nplatform's balance and subtract from the destination account's balance.\n\nReversing a transfer that was made for a [destination\ncharge](/docs/connect/destination-charges) is allowed only up to the amount of\nthe charge. It is possible to reverse a\n[transfer_group](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options)\ntransfer only if the destination account has enough balance to cover the\nreversal.\n\nRelated guide: [Reverse transfers](https://docs.stripe.com/connect/separate-charges-and-transfers#reverse-transfers)", "x-expandableFields": ["balance_transaction", "destination_payment_refund", "source_refund", "transfer"], "x-resourceId": "transfer_reversal"}
```
