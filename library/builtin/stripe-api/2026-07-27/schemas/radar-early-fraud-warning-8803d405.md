---
title: radar.early_fraud_warning
page_id: schema-radar-early-fraud-warning-8803d405
path: schemas
description: |-
    An early fraud warning indicates that the card issuer has notified us that a
    charge may be fraudulent.

    Related guide: [Early fraud warnings](https://docs.stripe.com/disputes/measuring#early-fraud-warnings)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# radar.early_fraud_warning

An early fraud warning indicates that the card issuer has notified us that a
charge may be fraudulent.

Related guide: [Early fraud warnings](https://docs.stripe.com/disputes/measuring#early-fraud-warnings)

```yaml
{"title": "RadarEarlyFraudWarning", "required": ["actionable", "charge", "created", "fraud_type", "id", "livemode", "object"], "type": "object", "properties": {"actionable": {"type": "boolean", "description": "An EFW is actionable if it has not received a dispute and has not been fully refunded. You may wish to proactively refund a charge that receives an EFW, in order to avoid receiving a dispute later."}, "charge": {"description": "ID of the charge this early fraud warning is for, optionally expanded.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "fraud_type": {"maxLength": 5000, "type": "string", "description": "The type of fraud labelled by the issuer. One of `card_never_received`, `fraudulent_card_application`, `made_with_counterfeit_card`, `made_with_lost_card`, `made_with_stolen_card`, `misc`, `unauthorized_use_of_card`."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["radar.early_fraud_warning"]}, "payment_intent": {"description": "ID of the Payment Intent this early fraud warning is for, optionally expanded.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}}, "description": "An early fraud warning indicates that the card issuer has notified us that a\ncharge may be fraudulent.\n\nRelated guide: [Early fraud warnings](https://docs.stripe.com/disputes/measuring#early-fraud-warnings)", "x-expandableFields": ["charge", "payment_intent"], "x-resourceId": "radar.early_fraud_warning"}
```
