---
title: country_spec
page_id: schema-country-spec-f01f822a
path: schemas
description: |-
    Stripe needs to collect certain pieces of information about each account
    created. These requirements can differ depending on the account's country. The
    Country Specs API makes these rules available to your integration.

    You can also view the information from this API call as [an online
    guide](/docs/connect/required-verification-information).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# country_spec

Stripe needs to collect certain pieces of information about each account
created. These requirements can differ depending on the account's country. The
Country Specs API makes these rules available to your integration.

You can also view the information from this API call as [an online
guide](/docs/connect/required-verification-information).

```yaml
{"title": "CountrySpec", "required": ["default_currency", "id", "object", "supported_bank_account_currencies", "supported_payment_currencies", "supported_payment_methods", "supported_transfer_countries", "verification_fields"], "type": "object", "properties": {"default_currency": {"maxLength": 5000, "type": "string", "description": "The default currency for this country. This applies to both payment methods and bank accounts."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object. Represented as the ISO country code for this country."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["country_spec"]}, "supported_bank_account_currencies": {"type": "object", "additionalProperties": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}, "description": "Currencies that can be accepted in the specific country (for transfers)."}, "supported_payment_currencies": {"type": "array", "description": "Currencies that can be accepted in the specified country (for payments).", "items": {"maxLength": 5000, "type": "string"}}, "supported_payment_methods": {"type": "array", "description": "Payment methods available in the specified country. You may need to enable some payment methods (e.g., [ACH](https://stripe.com/docs/ach)) on your account before they appear in this list. The `stripe` payment method refers to [charging through your platform](https://stripe.com/docs/connect/destination-charges).", "items": {"maxLength": 5000, "type": "string"}}, "supported_transfer_countries": {"type": "array", "description": "Countries that can accept transfers from the specified country.", "items": {"maxLength": 5000, "type": "string"}}, "verification_fields": {"$ref": "#/components/schemas/country_spec_verification_fields"}}, "description": "Stripe needs to collect certain pieces of information about each account\ncreated. These requirements can differ depending on the account's country. The\nCountry Specs API makes these rules available to your integration.\n\nYou can also view the information from this API call as [an online\nguide](/docs/connect/required-verification-information).", "x-expandableFields": ["verification_fields"], "x-resourceId": "country_spec"}
```
