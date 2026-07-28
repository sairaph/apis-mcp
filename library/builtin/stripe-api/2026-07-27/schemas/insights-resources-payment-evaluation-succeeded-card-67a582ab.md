---
title: insights_resources_payment_evaluation_succeeded_card
page_id: schema-insights-resources-payment-evaluation-succeeded-card-67a582ab
path: schemas
description: Details of an succeeded card outcome attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_succeeded_card

Details of an succeeded card outcome attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationSucceededCard", "required": ["address_line1_check", "address_postal_code_check", "cvc_check"], "type": "object", "properties": {"address_line1_check": {"type": "string", "description": "Result of the address line 1 check.", "enum": ["fail", "pass", "unavailable", "unchecked"]}, "address_postal_code_check": {"type": "string", "description": "Indicates whether the cardholder provided a postal code and if it matched the cardholder’s billing address.", "enum": ["fail", "pass", "unavailable", "unchecked"]}, "cvc_check": {"type": "string", "description": "Result of the CVC check.", "enum": ["fail", "pass", "unavailable", "unchecked"]}}, "description": "Details of an succeeded card outcome attached to this payment evaluation.", "x-expandableFields": []}
```
