---
title: tax.registration
page_id: schema-tax-registration-1b2168e5
path: schemas
description: |-
    A Tax `Registration` lets us know that your business is registered to collect tax on payments within a region, enabling you to [automatically collect tax](https://docs.stripe.com/tax).

    Stripe doesn't register on your behalf with the relevant authorities when you create a Tax `Registration` object. For more information on how to register to collect tax, see [our guide](https://docs.stripe.com/tax/registering).

    Related guide: [Using the Registrations API](https://docs.stripe.com/tax/registrations-api)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax.registration

A Tax `Registration` lets us know that your business is registered to collect tax on payments within a region, enabling you to [automatically collect tax](https://docs.stripe.com/tax).

Stripe doesn't register on your behalf with the relevant authorities when you create a Tax `Registration` object. For more information on how to register to collect tax, see [our guide](https://docs.stripe.com/tax/registering).

Related guide: [Using the Registrations API](https://docs.stripe.com/tax/registrations-api)

```yaml
{"title": "TaxProductRegistrationsResourceTaxRegistration", "required": ["active_from", "country", "country_options", "created", "id", "livemode", "object", "status"], "type": "object", "properties": {"active_from": {"type": "integer", "description": "Time at which the registration becomes active. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2))."}, "country_options": {"$ref": "#/components/schemas/tax_product_registrations_resource_country_options"}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expires_at": {"type": "integer", "description": "If set, the registration stops being active at this time. If not set, the registration will be active indefinitely. Measured in seconds since the Unix epoch.", "format": "unix-time", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["tax.registration"]}, "status": {"type": "string", "description": "The status of the registration. This field is present for convenience and can be deduced from `active_from` and `expires_at`.", "enum": ["active", "expired", "scheduled"]}}, "description": "A Tax `Registration` lets us know that your business is registered to collect tax on payments within a region, enabling you to [automatically collect tax](https://docs.stripe.com/tax).\n\nStripe doesn't register on your behalf with the relevant authorities when you create a Tax `Registration` object. For more information on how to register to collect tax, see [our guide](https://docs.stripe.com/tax/registering).\n\nRelated guide: [Using the Registrations API](https://docs.stripe.com/tax/registrations-api)", "x-expandableFields": ["country_options"], "x-resourceId": "tax.registration"}
```
