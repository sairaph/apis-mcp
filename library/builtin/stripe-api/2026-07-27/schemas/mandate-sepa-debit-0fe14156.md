---
title: mandate_sepa_debit
page_id: schema-mandate-sepa-debit-0fe14156
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_sepa_debit

```yaml
{"title": "mandate_sepa_debit", "required": ["reference", "url"], "type": "object", "properties": {"reference": {"maxLength": 5000, "type": "string", "description": "The unique reference of the mandate."}, "url": {"maxLength": 5000, "type": "string", "description": "The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively."}}, "description": "", "x-expandableFields": []}
```
