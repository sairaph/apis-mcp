---
title: card_generated_from_payment_method_details
page_id: schema-card-generated-from-payment-method-details-7463bdf2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# card_generated_from_payment_method_details

```yaml
{"title": "card_generated_from_payment_method_details", "required": ["type"], "type": "object", "properties": {"card_present": {"$ref": "#/components/schemas/payment_method_details_card_present"}, "type": {"maxLength": 5000, "type": "string", "description": "The type of payment method transaction-specific details from the transaction that generated this `card` payment method. Always `card_present`."}}, "description": "", "x-expandableFields": ["card_present"]}
```
