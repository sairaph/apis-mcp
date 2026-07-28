---
title: insights_resources_payment_evaluation_dispute_opened
page_id: schema-insights-resources-payment-evaluation-dispute-opened-46780ea6
path: schemas
description: Dispute opened event details attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_dispute_opened

Dispute opened event details attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationDisputeOpened", "required": ["amount", "currency", "reason"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to dispute for this payment. A positive integer representing how much to charge in [the smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency)."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "reason": {"type": "string", "description": "Reason given by cardholder for dispute.", "enum": ["account_not_available", "credit_not_processed", "customer_initiated", "duplicate", "fraudulent", "general", "noncompliant", "product_not_received", "product_unacceptable", "subscription_canceled", "unrecognized"]}}, "description": "Dispute opened event details attached to this payment evaluation.", "x-expandableFields": []}
```
