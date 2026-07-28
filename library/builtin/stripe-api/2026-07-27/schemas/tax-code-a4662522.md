---
title: tax_code
page_id: schema-tax-code-a4662522
path: schemas
description: '[Tax codes](https://stripe.com/docs/tax/tax-categories) classify goods and services for tax purposes.'
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_code

[Tax codes](https://stripe.com/docs/tax/tax-categories) classify goods and services for tax purposes.

```yaml
{"title": "TaxProductResourceTaxCode", "required": ["description", "id", "name", "object"], "type": "object", "properties": {"description": {"maxLength": 5000, "type": "string", "description": "A detailed description of which types of products the tax code represents."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "name": {"maxLength": 5000, "type": "string", "description": "A short name for the tax code."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["tax_code"]}}, "description": "[Tax codes](https://stripe.com/docs/tax/tax-categories) classify goods and services for tax purposes.", "x-expandableFields": [], "x-resourceId": "tax_code"}
```
