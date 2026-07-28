---
title: payment_method_details_payment_record_paynow
page_id: schema-payment-method-details-payment-record-paynow-dac2a74f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_payment_record_paynow

```yaml
{"title": "payment_method_details_payment_record_paynow", "type": "object", "properties": {"location": {"maxLength": 5000, "type": "string", "description": "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to."}, "reader": {"maxLength": 5000, "type": "string", "description": "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on."}, "reference": {"maxLength": 5000, "type": "string", "description": "Reference number associated with this PayNow payment", "nullable": true}}, "description": "", "x-expandableFields": []}
```
