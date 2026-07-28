---
title: payment_method_details_passthrough_card
page_id: schema-payment-method-details-passthrough-card-c3907b47
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_passthrough_card

```yaml
{"title": "payment_method_details_passthrough_card", "type": "object", "properties": {"brand": {"maxLength": 5000, "type": "string", "description": "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.", "nullable": true}, "exp_month": {"type": "integer", "description": "Two-digit number representing the card's expiration month.", "nullable": true}, "exp_year": {"type": "integer", "description": "Four-digit number representing the card's expiration year.", "nullable": true}, "funding": {"maxLength": 5000, "type": "string", "description": "Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "The last four digits of the card.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
