---
title: issuing_cardholder_requirements
page_id: schema-issuing-cardholder-requirements-ade24aa9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_cardholder_requirements

```yaml
{"title": "IssuingCardholderRequirements", "type": "object", "properties": {"disabled_reason": {"type": "string", "description": "If `disabled_reason` is present, all cards will decline authorizations with `cardholder_verification_required` reason.", "nullable": true, "enum": ["listed", "rejected.listed", "requirements.past_due", "under_review"]}, "past_due": {"type": "array", "description": "Array of fields that need to be collected in order to verify and re-enable the cardholder.", "nullable": true, "items": {"type": "string", "enum": ["company.tax_id", "individual.card_issuing.user_terms_acceptance.date", "individual.card_issuing.user_terms_acceptance.ip", "individual.dob.day", "individual.dob.month", "individual.dob.year", "individual.first_name", "individual.last_name", "individual.verification.document"], "x-stripeBypassValidation": true}}}, "description": "", "x-expandableFields": []}
```
