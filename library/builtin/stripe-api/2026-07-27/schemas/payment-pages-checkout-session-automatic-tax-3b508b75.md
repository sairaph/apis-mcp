---
title: payment_pages_checkout_session_automatic_tax
page_id: schema-payment-pages-checkout-session-automatic-tax-3b508b75
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_automatic_tax

```yaml
{"title": "PaymentPagesCheckoutSessionAutomaticTax", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Indicates whether automatic tax is enabled for the session"}, "liability": {"description": "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/connect_account_reference"}]}, "provider": {"maxLength": 5000, "type": "string", "description": "The tax provider powering automatic tax.", "nullable": true}, "status": {"type": "string", "description": "The status of the most recent automated tax calculation for this session.", "nullable": true, "enum": ["complete", "failed", "requires_location_inputs"]}}, "description": "", "x-expandableFields": ["liability"]}
```
