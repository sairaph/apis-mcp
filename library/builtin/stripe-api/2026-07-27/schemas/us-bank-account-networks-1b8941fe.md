---
title: us_bank_account_networks
page_id: schema-us-bank-account-networks-1b8941fe
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# us_bank_account_networks

```yaml
{"title": "us_bank_account_networks", "required": ["supported"], "type": "object", "properties": {"preferred": {"maxLength": 5000, "type": "string", "description": "The preferred network.", "nullable": true}, "supported": {"type": "array", "description": "All supported networks.", "items": {"type": "string", "enum": ["ach", "us_domestic_wire"]}}}, "description": "", "x-expandableFields": []}
```
