---
title: received_payment_method_details_financial_account
page_id: schema-received-payment-method-details-financial-account-e6b58f73
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# received_payment_method_details_financial_account

```yaml
{"title": "received_payment_method_details_financial_account", "required": ["id", "network"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "The FinancialAccount ID."}, "network": {"type": "string", "description": "The rails the ReceivedCredit was sent over. A FinancialAccount can only send funds over `stripe`.", "enum": ["stripe"]}}, "description": "", "x-expandableFields": []}
```
