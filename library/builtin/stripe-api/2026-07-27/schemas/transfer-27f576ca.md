---
title: transfer
page_id: schema-transfer-27f576ca
path: schemas
description: |-
    A `Transfer` object is created when you move funds between Stripe accounts as
    part of Connect.

    Before April 6, 2017, transfers also represented movement of funds from a
    Stripe account to a card or bank account. This behavior has since been split
    out into a [Payout](https://api.stripe.com#payout_object) object, with corresponding payout endpoints. For more
    information, read about the
    [transfer/payout split](https://docs.stripe.com/transfer-payout-split).

    Related guide: [Creating separate charges and transfers](https://docs.stripe.com/connect/separate-charges-and-transfers)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# transfer

A `Transfer` object is created when you move funds between Stripe accounts as
part of Connect.

Before April 6, 2017, transfers also represented movement of funds from a
Stripe account to a card or bank account. This behavior has since been split
out into a [Payout](https://api.stripe.com#payout_object) object, with corresponding payout endpoints. For more
information, read about the
[transfer/payout split](https://docs.stripe.com/transfer-payout-split).

Related guide: [Creating separate charges and transfers](https://docs.stripe.com/connect/separate-charges-and-transfers)

```yaml
{"title": "Transfer", "required": ["amount", "amount_reversed", "created", "currency", "id", "livemode", "metadata", "object", "reversals", "reversed"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount in cents (or local equivalent) to be transferred."}, "amount_reversed": {"type": "integer", "description": "Amount in cents (or local equivalent) reversed (can be less than the amount attribute on the transfer if a partial reversal was issued)."}, "balance_transaction": {"description": "Balance transaction that describes the impact of this transfer on your account balance.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/balance_transaction"}]}}, "created": {"type": "integer", "description": "Time that this record of the transfer was first created.", "format": "unix-time"}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users.", "nullable": true}, "destination": {"description": "ID of the Stripe account the transfer was sent to.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}, "destination_payment": {"description": "If the destination is a Stripe account, this will be the ID of the payment that the destination account received for the transfer.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["transfer"]}, "reversals": {"title": "TransferReversalList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/transfer_reversal"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "A list of reversals that have been applied to the transfer.", "x-expandableFields": ["data"]}, "reversed": {"type": "boolean", "description": "Whether the transfer has been fully reversed. If the transfer is only partially reversed, this attribute will still be false."}, "source_transaction": {"description": "ID of the charge that was used to fund the transfer. If null, the transfer was funded from the available balance.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "source_type": {"maxLength": 5000, "type": "string", "description": "The source balance this transfer came from. One of `card`, `fpx`, or `bank_account`."}, "transfer_group": {"maxLength": 5000, "type": "string", "description": "A string that identifies this transaction as part of a group. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.", "nullable": true}}, "description": "A `Transfer` object is created when you move funds between Stripe accounts as\npart of Connect.\n\nBefore April 6, 2017, transfers also represented movement of funds from a\nStripe account to a card or bank account. This behavior has since been split\nout into a [Payout](https://api.stripe.com#payout_object) object, with corresponding payout endpoints. For more\ninformation, read about the\n[transfer/payout split](https://docs.stripe.com/transfer-payout-split).\n\nRelated guide: [Creating separate charges and transfers](https://docs.stripe.com/connect/separate-charges-and-transfers)", "x-expandableFields": ["balance_transaction", "destination", "destination_payment", "reversals", "source_transaction"], "x-resourceId": "transfer"}
```
