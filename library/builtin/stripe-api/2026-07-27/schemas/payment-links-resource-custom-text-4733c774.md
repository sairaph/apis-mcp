---
title: payment_links_resource_custom_text
page_id: schema-payment-links-resource-custom-text-4733c774
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_custom_text

```yaml
{"title": "PaymentLinksResourceCustomText", "type": "object", "properties": {"after_submit": {"description": "Custom text that should be displayed after the payment confirmation button.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_custom_text_position"}]}, "shipping_address": {"description": "Custom text that should be displayed alongside shipping address collection.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_custom_text_position"}]}, "submit": {"description": "Custom text that should be displayed alongside the payment confirmation button.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_custom_text_position"}]}, "terms_of_service_acceptance": {"description": "Custom text that should be displayed in place of the default terms of service agreement text.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_custom_text_position"}]}}, "description": "", "x-expandableFields": ["after_submit", "shipping_address", "submit", "terms_of_service_acceptance"]}
```
