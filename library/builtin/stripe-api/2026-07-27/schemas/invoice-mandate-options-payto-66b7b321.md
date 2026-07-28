---
title: invoice_mandate_options_payto
page_id: schema-invoice-mandate-options-payto-66b7b321
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_mandate_options_payto

```yaml
{"title": "invoice_mandate_options_payto", "type": "object", "properties": {"amount": {"type": "integer", "description": "The maximum amount that can be collected in a single invoice. If you don't specify a maximum, then there is no limit.", "nullable": true}, "amount_type": {"type": "string", "description": "Only `maximum` is supported.", "nullable": true, "enum": ["fixed", "maximum"]}, "purpose": {"type": "string", "description": "The purpose for which payments are made. Has a default value based on your merchant category code.", "nullable": true, "enum": ["dependant_support", "government", "loan", "mortgage", "other", "pension", "personal", "retail", "salary", "tax", "utility"]}}, "description": "", "x-expandableFields": []}
```
