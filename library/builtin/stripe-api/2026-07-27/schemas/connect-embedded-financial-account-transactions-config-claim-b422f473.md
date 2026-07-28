---
title: connect_embedded_financial_account_transactions_config_claim
page_id: schema-connect-embedded-financial-account-transactions-config-claim-b422f473
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_financial_account_transactions_config_claim

```yaml
{"title": "ConnectEmbeddedFinancialAccountTransactionsConfigClaim", "required": ["enabled", "features"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the embedded component is enabled."}, "features": {"$ref": "#/components/schemas/connect_embedded_financial_account_transactions_features"}}, "description": "", "x-expandableFields": ["features"]}
```
