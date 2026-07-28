---
title: tax.settings
page_id: schema-tax-settings-44ad7693
path: schemas
description: |-
    You can use Tax `Settings` to manage configurations used by Stripe Tax calculations.

    Related guide: [Using the Settings API](https://docs.stripe.com/tax/settings-api)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax.settings

You can use Tax `Settings` to manage configurations used by Stripe Tax calculations.

Related guide: [Using the Settings API](https://docs.stripe.com/tax/settings-api)

```yaml
{"title": "TaxProductResourceTaxSettings", "required": ["defaults", "livemode", "object", "status", "status_details"], "type": "object", "properties": {"defaults": {"$ref": "#/components/schemas/tax_product_resource_tax_settings_defaults"}, "head_office": {"description": "The place where your business is located.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/tax_product_resource_tax_settings_head_office"}]}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["tax.settings"]}, "status": {"type": "string", "description": "The status of the Tax `Settings`.", "enum": ["active", "pending"]}, "status_details": {"$ref": "#/components/schemas/tax_product_resource_tax_settings_status_details"}}, "description": "You can use Tax `Settings` to manage configurations used by Stripe Tax calculations.\n\nRelated guide: [Using the Settings API](https://docs.stripe.com/tax/settings-api)", "x-expandableFields": ["defaults", "head_office", "status_details"], "x-resourceId": "tax.settings"}
```
