---
title: payment_method_details_alma
page_id: schema-payment-method-details-alma-2871a608
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_alma

```yaml
{"title": "payment_method_details_alma", "type": "object", "properties": {"installments": {"$ref": "#/components/schemas/alma_installments"}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "The Alma transaction ID associated with this payment.", "nullable": true}}, "description": "", "x-expandableFields": ["installments"]}
```
