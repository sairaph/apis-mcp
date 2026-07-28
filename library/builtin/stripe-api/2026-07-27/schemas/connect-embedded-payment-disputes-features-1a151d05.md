---
title: connect_embedded_payment_disputes_features
page_id: schema-connect-embedded-payment-disputes-features-1a151d05
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_payment_disputes_features

```yaml
{"title": "ConnectEmbeddedPaymentDisputesFeatures", "required": ["destination_on_behalf_of_charge_management", "dispute_management", "refund_management"], "type": "object", "properties": {"destination_on_behalf_of_charge_management": {"type": "boolean", "description": "Whether connected accounts can manage destination charges that are created on behalf of them. This is `false` by default."}, "dispute_management": {"type": "boolean", "description": "Whether responding to disputes is enabled, including submitting evidence and accepting disputes. This is `true` by default."}, "refund_management": {"type": "boolean", "description": "Whether sending refunds is enabled. This is `true` by default."}}, "description": "", "x-expandableFields": []}
```
