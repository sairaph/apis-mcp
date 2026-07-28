---
title: payment_links_resource_consent_collection
page_id: schema-payment-links-resource-consent-collection-d02bc0cf
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_consent_collection

```yaml
{"title": "PaymentLinksResourceConsentCollection", "type": "object", "properties": {"payment_method_reuse_agreement": {"description": "Settings related to the payment method reuse text shown in the Checkout UI.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_payment_method_reuse_agreement"}]}, "promotions": {"type": "string", "description": "If set to `auto`, enables the collection of customer consent for promotional communications.", "nullable": true, "enum": ["auto", "none"]}, "terms_of_service": {"type": "string", "description": "If set to `required`, it requires cutomers to accept the terms of service before being able to pay. If set to `none`, customers won't be shown a checkbox to accept the terms of service.", "nullable": true, "enum": ["none", "required"]}}, "description": "", "x-expandableFields": ["payment_method_reuse_agreement"]}
```
