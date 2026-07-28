---
title: insights_resources_payment_evaluation_rejected_card
page_id: schema-insights-resources-payment-evaluation-rejected-card-d9af8784
path: schemas
description: Details of an rejected card outcome attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_rejected_card

Details of an rejected card outcome attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationRejectedCard", "required": ["address_line1_check", "address_postal_code_check", "cvc_check", "reason"], "type": "object", "properties": {"address_line1_check": {"type": "string", "description": "Result of the address line 1 check.", "enum": ["fail", "pass", "unavailable", "unchecked"]}, "address_postal_code_check": {"type": "string", "description": "Indicates whether the cardholder provided a postal code and if it matched the cardholder’s billing address.", "enum": ["fail", "pass", "unavailable", "unchecked"]}, "cvc_check": {"type": "string", "description": "Result of the CVC check.", "enum": ["fail", "pass", "unavailable", "unchecked"]}, "reason": {"type": "string", "description": "Card issuer's reason for the network decline.", "enum": ["authentication_failed", "do_not_honor", "expired", "incorrect_cvc", "incorrect_number", "incorrect_postal_code", "insufficient_funds", "invalid_account", "lost_card", "other", "processing_error", "reported_stolen", "try_again_later"]}}, "description": "Details of an rejected card outcome attached to this payment evaluation.", "x-expandableFields": []}
```
