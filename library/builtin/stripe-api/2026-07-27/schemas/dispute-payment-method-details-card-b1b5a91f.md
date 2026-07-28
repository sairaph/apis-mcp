---
title: dispute_payment_method_details_card
page_id: schema-dispute-payment-method-details-card-b1b5a91f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_payment_method_details_card

```yaml
{"title": "DisputePaymentMethodDetailsCard", "required": ["brand", "case_type"], "type": "object", "properties": {"brand": {"maxLength": 5000, "type": "string", "description": "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`."}, "case_type": {"type": "string", "description": "The type of dispute opened. Different case types may have varying fees and financial impact.", "enum": ["block", "chargeback", "compliance", "inquiry", "resolution"]}, "network_reason_code": {"maxLength": 5000, "type": "string", "description": "The card network's specific dispute reason code, which maps to one of Stripe's primary dispute categories to simplify response guidance. The [Network code map](https://stripe.com/docs/disputes/categories#network-code-map) lists all available dispute reason codes by network.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
