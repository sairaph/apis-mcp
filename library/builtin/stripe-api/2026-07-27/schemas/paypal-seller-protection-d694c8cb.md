---
title: paypal_seller_protection
page_id: schema-paypal-seller-protection-d694c8cb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# paypal_seller_protection

```yaml
{"title": "paypal_seller_protection", "required": ["status"], "type": "object", "properties": {"dispute_categories": {"type": "array", "description": "An array of conditions that are covered for the transaction, if applicable.", "nullable": true, "items": {"type": "string", "enum": ["fraudulent", "product_not_received"], "x-stripeBypassValidation": true}}, "status": {"type": "string", "description": "Indicates whether the transaction is eligible for PayPal's seller protection.", "enum": ["eligible", "not_eligible", "partially_eligible"]}}, "description": "", "x-expandableFields": []}
```
