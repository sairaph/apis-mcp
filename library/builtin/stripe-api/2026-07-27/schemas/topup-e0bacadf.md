---
title: topup
page_id: schema-topup-e0bacadf
path: schemas
description: |-
    To top up your Stripe balance, you create a top-up object. You can retrieve
    individual top-ups, as well as list all top-ups. Top-ups are identified by a
    unique, random ID.

    Related guide: [Topping up your platform account](https://docs.stripe.com/connect/top-ups)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# topup

To top up your Stripe balance, you create a top-up object. You can retrieve
individual top-ups, as well as list all top-ups. Top-ups are identified by a
unique, random ID.

Related guide: [Topping up your platform account](https://docs.stripe.com/connect/top-ups)

```yaml
{"title": "Topup", "required": ["amount", "created", "currency", "id", "livemode", "metadata", "object", "status"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount transferred."}, "balance_transaction": {"description": "ID of the balance transaction that describes the impact of this top-up on your account balance. May not be specified depending on status of top-up.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/balance_transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/balance_transaction"}]}}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies)."}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users.", "nullable": true}, "expected_availability_date": {"type": "integer", "description": "Date the funds are expected to arrive in your Stripe account for payouts. This factors in delays like weekends or bank holidays. May not be specified depending on status of top-up.", "nullable": true}, "failure_code": {"maxLength": 5000, "type": "string", "description": "Error code explaining reason for top-up failure if available (see [the errors section](/api/errors) for a list of codes).", "nullable": true}, "failure_message": {"maxLength": 5000, "type": "string", "description": "Message to user further explaining reason for top-up failure if available.", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["topup"]}, "source": {"description": "The source field is deprecated. It might not always be present in the API response.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/source"}]}, "statement_descriptor": {"maxLength": 5000, "type": "string", "description": "Extra information about a top-up. This will appear on your source's bank statement. It must contain at least one letter.", "nullable": true}, "status": {"type": "string", "description": "The status of the top-up is either `canceled`, `failed`, `pending`, `reversed`, or `succeeded`.", "enum": ["canceled", "failed", "pending", "reversed", "succeeded"]}, "transfer_group": {"maxLength": 5000, "type": "string", "description": "A string that identifies this top-up as part of a group.", "nullable": true}}, "description": "To top up your Stripe balance, you create a top-up object. You can retrieve\nindividual top-ups, as well as list all top-ups. Top-ups are identified by a\nunique, random ID.\n\nRelated guide: [Topping up your platform account](https://docs.stripe.com/connect/top-ups)", "x-expandableFields": ["balance_transaction", "source"], "x-resourceId": "topup"}
```
