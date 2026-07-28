---
title: issuing_card_shipping_address_validation
page_id: schema-issuing-card-shipping-address-validation-0ca0f30c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_card_shipping_address_validation

```yaml
{"title": "IssuingCardShippingAddressValidation", "required": ["mode"], "type": "object", "properties": {"mode": {"type": "string", "description": "The address validation capabilities to use.", "enum": ["disabled", "normalization_only", "validation_and_normalization"]}, "normalized_address": {"description": "The normalized shipping address.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "result": {"type": "string", "description": "The validation result for the shipping address.", "nullable": true, "enum": ["indeterminate", "likely_deliverable", "likely_undeliverable"]}}, "description": "", "x-expandableFields": ["normalized_address"]}
```
