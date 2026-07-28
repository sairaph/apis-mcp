---
title: country_spec_verification_field_details
page_id: schema-country-spec-verification-field-details-be1f5271
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# country_spec_verification_field_details

```yaml
{"title": "CountrySpecVerificationFieldDetails", "required": ["additional", "minimum"], "type": "object", "properties": {"additional": {"type": "array", "description": "Additional fields which are only required for some users.", "items": {"maxLength": 5000, "type": "string"}}, "minimum": {"type": "array", "description": "Fields which every account must eventually provide.", "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": []}
```
