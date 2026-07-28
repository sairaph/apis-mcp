---
title: Create a payout
page_id: operation-post-v1-payouts-de833ec2
path: operations/untagged
description: |-
    <p>To send funds to your own bank account, create a new payout object. Your <a href="#balance">Stripe balance</a> must cover the payout amount. If it doesn’t, you receive an “Insufficient Funds” error.</p>

    <p>If your API key is in test mode, money won’t actually be sent, though every other action occurs as if you’re in live mode.</p>

    <p>If you create a manual payout on a Stripe account that uses multiple payment source types, you need to specify the source type balance that the payout draws from. The <a href="/api/balances/object">balance object</a> details available and pending amounts by source type.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payouts
operation_ids:
    - PostPayouts
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a payout

`POST /v1/payouts`

Operation ID: `PostPayouts`

<p>To send funds to your own bank account, create a new payout object. Your <a href="#balance">Stripe balance</a> must cover the payout amount. If it doesn’t, you receive an “Insufficient Funds” error.</p>

<p>If your API key is in test mode, money won’t actually be sent, though every other action occurs as if you’re in live mode.</p>

<p>If you create a manual payout on a Stripe account that uses multiple payment source types, you need to specify the source type balance that the payout draws from. The <a href="/api/balances/object">balance object</a> details available and pending amounts by source type.</p>

## Definition

```yaml
{"summary": "Create a payout", "description": "<p>To send funds to your own bank account, create a new payout object. Your <a href=\"#balance\">Stripe balance</a> must cover the payout amount. If it doesn’t, you receive an “Insufficient Funds” error.</p>\n\n<p>If your API key is in test mode, money won’t actually be sent, though every other action occurs as if you’re in live mode.</p>\n\n<p>If you create a manual payout on a Stripe account that uses multiple payment source types, you need to specify the source type balance that the payout draws from. The <a href=\"/api/balances/object\">balance object</a> details available and pending amounts by source type.</p>", "operationId": "PostPayouts", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer in cents representing how much to payout."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "destination": {"type": "string", "description": "The ID of a bank account or a card to send the payout to. If you don't provide a destination, we use the default external account for the specified currency."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "method": {"maxLength": 5000, "type": "string", "description": "The method used to send this payout, which is `standard` or `instant`. We support `instant` for payouts to debit cards and bank accounts in certain countries. Learn more about [bank support for Instant Payouts](https://stripe.com/docs/payouts/instant-payouts-banks).", "enum": ["instant", "standard"], "x-stripeBypassValidation": true}, "payout_method": {"type": "string", "description": "The ID of a v2 FinancialAccount to send funds to."}, "source_type": {"maxLength": 5000, "type": "string", "description": "The balance type of your Stripe balance to draw this payout from. Balances for different payment sources are kept separately. You can find the amounts with the Balances API. One of `bank_account`, `card`, or `fpx`.", "enum": ["bank_account", "card", "fpx"], "x-stripeBypassValidation": true}, "statement_descriptor": {"maxLength": 22, "type": "string", "description": "A string that displays on the recipient's bank or card statement (up to 22 characters). A `statement_descriptor` that's longer than 22 characters return an error. Most banks truncate this information and display it inconsistently. Some banks might not display it at all. For US ACH payouts, this maps to the ACH Company Entry Description field, which the NACHA standard limits to 10 characters. Stripe truncates descriptors longer than 10 characters for US ACH payouts.", "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payout"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
