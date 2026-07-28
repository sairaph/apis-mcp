---
title: payment_links_resource_card_restrictions
page_id: schema-payment-links-resource-card-restrictions-74a672a9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_card_restrictions

```yaml
{"title": "PaymentLinksResourceCardRestrictions", "required": ["brands_blocked"], "type": "object", "properties": {"brands_blocked": {"type": "array", "description": "The card brands to block. If a customer enters or selects a card belonging to a blocked brand, they can't complete the payment.", "items": {"type": "string", "enum": ["american_express", "discover_global_network", "mastercard", "visa"]}}}, "description": "", "x-expandableFields": []}
```
