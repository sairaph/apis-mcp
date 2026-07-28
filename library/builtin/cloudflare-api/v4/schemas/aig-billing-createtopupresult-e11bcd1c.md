---
title: aig-billing_CreateTopupResult
page_id: schema-aig-billing-createtopupresult-e11bcd1c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_CreateTopupResult

```yaml
{"type": "object", "properties": {"brand": {"description": "Card brand (visa, mastercard, etc.).", "type": "string"}, "client_secret": {"description": "Stripe PaymentIntent client secret.", "type": "string", "nullable": true}, "last4": {"description": "Last 4 digits of card.", "type": "string"}, "onboarding": {"description": "Whether the user was already onboarded.", "type": "boolean"}, "payment_intent_id": {"description": "Stripe invoice ID.", "type": "string"}}, "required": ["onboarding", "payment_intent_id", "client_secret"]}
```
