---
title: payment_intent_next_action_display_bank_transfer_instructions
page_id: schema-payment-intent-next-action-display-bank-transfer-instructions-a418bbba
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_display_bank_transfer_instructions

```yaml
{"title": "PaymentIntentNextActionDisplayBankTransferInstructions", "required": ["type"], "type": "object", "properties": {"amount_remaining": {"type": "integer", "description": "The remaining amount that needs to be transferred to complete the payment.", "nullable": true}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency", "nullable": true}, "financial_addresses": {"type": "array", "description": "A list of financial addresses that can be used to fund the customer balance", "items": {"$ref": "#/components/schemas/funding_instructions_bank_transfer_financial_address"}}, "hosted_instructions_url": {"maxLength": 5000, "type": "string", "description": "A link to a hosted page that guides your customer through completing the transfer.", "nullable": true}, "reference": {"maxLength": 5000, "type": "string", "description": "A string identifying this payment. Instruct your customer to include this code in the reference or memo field of their bank transfer.", "nullable": true}, "type": {"type": "string", "description": "Type of bank transfer", "enum": ["eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["financial_addresses"]}
```
