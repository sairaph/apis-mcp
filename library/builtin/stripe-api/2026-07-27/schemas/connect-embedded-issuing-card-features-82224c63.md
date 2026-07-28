---
title: connect_embedded_issuing_card_features
page_id: schema-connect-embedded-issuing-card-features-82224c63
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_issuing_card_features

```yaml
{"title": "ConnectEmbeddedIssuingCardFeatures", "required": ["card_management", "card_spend_dispute_management", "cardholder_management", "spend_control_management"], "type": "object", "properties": {"card_management": {"type": "boolean", "description": "Whether to allow card management features."}, "card_spend_dispute_management": {"type": "boolean", "description": "Whether to allow card spend dispute management features."}, "cardholder_management": {"type": "boolean", "description": "Whether to allow cardholder management features."}, "spend_control_management": {"type": "boolean", "description": "Whether to allow spend control management features."}}, "description": "", "x-expandableFields": []}
```
