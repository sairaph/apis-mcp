---
title: payment_method_card_present_networks
page_id: schema-payment-method-card-present-networks-ac7482ea
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_card_present_networks

```yaml
{"title": "payment_method_card_present_networks", "required": ["available"], "type": "object", "properties": {"available": {"type": "array", "description": "All networks available for selection via [payment_method_options.card.network](/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).", "items": {"maxLength": 5000, "type": "string"}}, "preferred": {"maxLength": 5000, "type": "string", "description": "The preferred network for the card.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
