---
title: payment_flows_automatic_payment_methods_setup_intent
page_id: schema-payment-flows-automatic-payment-methods-setup-intent-694171b1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_automatic_payment_methods_setup_intent

```yaml
{"title": "PaymentFlowsAutomaticPaymentMethodsSetupIntent", "type": "object", "properties": {"allow_redirects": {"type": "string", "description": "Controls whether this SetupIntent will accept redirect-based payment methods.\n\nRedirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://docs.stripe.com/api/setup_intents/confirm) this SetupIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the setup.", "enum": ["always", "never"]}, "enabled": {"type": "boolean", "description": "Automatically calculates compatible payment methods", "nullable": true}}, "description": "", "x-expandableFields": []}
```
