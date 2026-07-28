---
title: issuing_cardholder_card_issuing
page_id: schema-issuing-cardholder-card-issuing-996498cd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_cardholder_card_issuing

```yaml
{"title": "IssuingCardholderCardIssuing", "type": "object", "properties": {"user_terms_acceptance": {"description": "Information about cardholder acceptance of Celtic [Authorized User Terms](https://stripe.com/docs/issuing/cards#accept-authorized-user-terms). Required for cards backed by a Celtic program.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_cardholder_user_terms_acceptance"}]}}, "description": "", "x-expandableFields": ["user_terms_acceptance"]}
```
