---
title: Capture a payment
page_id: operation-post-v1-charges-charge-capture-1711019b
path: operations/untagged
description: |-
    <p>Capture the payment of an existing, uncaptured charge that was created with the <code>capture</code> option set to false.</p>

    <p>Uncaptured payments expire a set number of days after they are created (<a href="/docs/charges/placing-a-hold">7 by default</a>), after which they are marked as refunded and capture attempts will fail.</p>

    <p>Don’t use this method to capture a PaymentIntent-initiated charge. Use <a href="/docs/api/payment_intents/capture">Capture a PaymentIntent</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/charges/{charge}/capture
operation_ids:
    - PostChargesChargeCapture
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Capture a payment

`POST /v1/charges/{charge}/capture`

Operation ID: `PostChargesChargeCapture`

<p>Capture the payment of an existing, uncaptured charge that was created with the <code>capture</code> option set to false.</p>

<p>Uncaptured payments expire a set number of days after they are created (<a href="/docs/charges/placing-a-hold">7 by default</a>), after which they are marked as refunded and capture attempts will fail.</p>

<p>Don’t use this method to capture a PaymentIntent-initiated charge. Use <a href="/docs/api/payment_intents/capture">Capture a PaymentIntent</a>.</p>

## Definition

```yaml
{"summary": "Capture a payment", "description": "<p>Capture the payment of an existing, uncaptured charge that was created with the <code>capture</code> option set to false.</p>\n\n<p>Uncaptured payments expire a set number of days after they are created (<a href=\"/docs/charges/placing-a-hold\">7 by default</a>), after which they are marked as refunded and capture attempts will fail.</p>\n\n<p>Don’t use this method to capture a PaymentIntent-initiated charge. Use <a href=\"/docs/api/payment_intents/capture\">Capture a PaymentIntent</a>.</p>", "operationId": "PostChargesChargeCapture", "parameters": [{"name": "charge", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer", "description": "The amount to capture, which must be less than or equal to the original amount."}, "application_fee": {"type": "integer", "description": "An application fee to add on to this charge."}, "application_fee_amount": {"type": "integer", "description": "An application fee amount to add on to this charge, which must be less than or equal to the original amount."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "receipt_email": {"type": "string", "description": "The email address to send this charge's receipt to. This will override the previously-specified email address for this charge, if one was set. Receipts will not be sent in test mode."}, "statement_descriptor": {"maxLength": 22, "type": "string", "description": "For a non-card charge, text that appears on the customer's statement as the statement descriptor. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).\n\nFor a card charge, this value is ignored unless you don't specify a `statement_descriptor_suffix`, in which case this value is used as the suffix."}, "statement_descriptor_suffix": {"maxLength": 22, "type": "string", "description": "Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement. If the account has no prefix value, the suffix is concatenated to the account's statement descriptor."}, "transfer_data": {"title": "transfer_data_specs", "type": "object", "properties": {"amount": {"type": "integer"}}, "description": "An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details."}, "transfer_group": {"type": "string", "description": "A string that identifies this transaction as part of a group. `transfer_group` may only be provided if it has not been set. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "transfer_data": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/charge"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
