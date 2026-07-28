---
title: outbound_transfers_payment_method_details_financial_account
page_id: schema-outbound-transfers-payment-method-details-financial-account-a829f1f2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# outbound_transfers_payment_method_details_financial_account

```yaml
{"title": "outbound_transfers_payment_method_details_financial_account", "required": ["id", "network"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "Token of the FinancialAccount."}, "network": {"type": "string", "description": "The rails used to send funds.", "enum": ["stripe"]}}, "description": "", "x-expandableFields": []}
```
