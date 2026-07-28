---
title: issuing_cardholder_individual
page_id: schema-issuing-cardholder-individual-48489e75
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_cardholder_individual

```yaml
{"title": "IssuingCardholderIndividual", "type": "object", "properties": {"card_issuing": {"description": "Information related to the card_issuing program for this cardholder.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_cardholder_card_issuing"}]}, "dob": {"description": "The date of birth of this cardholder.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_cardholder_individual_dob"}]}, "first_name": {"maxLength": 5000, "type": "string", "description": "The first name of this cardholder. Required before activating Cards. This field cannot contain any numbers, special characters (except periods, commas, hyphens, spaces and apostrophes) or non-latin letters.", "nullable": true}, "last_name": {"maxLength": 5000, "type": "string", "description": "The last name of this cardholder. Required before activating Cards. This field cannot contain any numbers, special characters (except periods, commas, hyphens, spaces and apostrophes) or non-latin letters.", "nullable": true}, "verification": {"description": "Government-issued ID document for this cardholder.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_cardholder_verification"}]}}, "description": "", "x-expandableFields": ["card_issuing", "dob", "verification"]}
```
