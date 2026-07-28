---
title: treasury_inbound_transfers_resource_failure_details
page_id: schema-treasury-inbound-transfers-resource-failure-details-010f80a3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_inbound_transfers_resource_failure_details

```yaml
{"title": "TreasuryInboundTransfersResourceFailureDetails", "required": ["code"], "type": "object", "properties": {"code": {"type": "string", "description": "Reason for the failure.", "enum": ["account_closed", "account_frozen", "bank_account_restricted", "bank_ownership_changed", "debit_not_authorized", "incorrect_account_holder_address", "incorrect_account_holder_name", "incorrect_account_holder_tax_id", "insufficient_funds", "invalid_account_number", "invalid_currency", "no_account", "other"]}}, "description": "", "x-expandableFields": []}
```
