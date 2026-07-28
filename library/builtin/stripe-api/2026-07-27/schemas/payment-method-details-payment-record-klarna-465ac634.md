---
title: payment_method_details_payment_record_klarna
page_id: schema-payment-method-details-payment-record-klarna-465ac634
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_payment_record_klarna

```yaml
{"title": "payment_method_details_payment_record_klarna", "type": "object", "properties": {"location": {"maxLength": 5000, "type": "string", "description": "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to."}, "payer_details": {"description": "The payer details for this transaction.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payments_primitives_payment_records_resource_payment_method_klarna_details_resource_payer_details"}]}, "payment_method_category": {"maxLength": 5000, "type": "string", "description": "The Klarna payment method used for this transaction. Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`", "nullable": true}, "preferred_locale": {"maxLength": 5000, "type": "string", "description": "Preferred language of the Klarna authorization page that the customer is redirected to. Can be one of `de-AT`, `en-AT`, `nl-BE`, `fr-BE`, `en-BE`, `de-DE`, `en-DE`, `da-DK`, `en-DK`, `es-ES`, `en-ES`, `fi-FI`, `sv-FI`, `en-FI`, `en-GB`, `en-IE`, `it-IT`, `en-IT`, `nl-NL`, `en-NL`, `nb-NO`, `en-NO`, `sv-SE`, `en-SE`, `en-US`, `es-US`, `fr-FR`, `en-FR`, `cs-CZ`, `en-CZ`, `ro-RO`, `en-RO`, `el-GR`, `en-GR`, `en-AU`, `en-NZ`, `en-CA`, `fr-CA`, `pl-PL`, `en-PL`, `pt-PT`, `en-PT`, `de-CH`, `fr-CH`, `it-CH`, or `en-CH`", "nullable": true}, "reader": {"maxLength": 5000, "type": "string", "description": "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on."}}, "description": "", "x-expandableFields": ["payer_details"]}
```
