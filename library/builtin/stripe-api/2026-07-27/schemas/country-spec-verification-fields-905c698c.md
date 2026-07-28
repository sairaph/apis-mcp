---
title: country_spec_verification_fields
page_id: schema-country-spec-verification-fields-905c698c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# country_spec_verification_fields

```yaml
{"title": "CountrySpecVerificationFields", "required": ["company", "individual"], "type": "object", "properties": {"company": {"$ref": "#/components/schemas/country_spec_verification_field_details"}, "individual": {"$ref": "#/components/schemas/country_spec_verification_field_details"}}, "description": "", "x-expandableFields": ["company", "individual"]}
```
