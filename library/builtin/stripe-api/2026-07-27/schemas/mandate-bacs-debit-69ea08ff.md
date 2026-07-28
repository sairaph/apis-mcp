---
title: mandate_bacs_debit
page_id: schema-mandate-bacs-debit-69ea08ff
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_bacs_debit

```yaml
{"title": "mandate_bacs_debit", "required": ["network_status", "reference", "url"], "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string", "description": "The display name for the account on this mandate.", "nullable": true}, "network_status": {"type": "string", "description": "The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.", "enum": ["accepted", "pending", "refused", "revoked"]}, "reference": {"maxLength": 5000, "type": "string", "description": "The unique reference identifying the mandate on the Bacs network."}, "revocation_reason": {"type": "string", "description": "When the mandate is revoked on the Bacs network this field displays the reason for the revocation.", "nullable": true, "enum": ["account_closed", "bank_account_restricted", "bank_ownership_changed", "could_not_process", "debit_not_authorized"]}, "service_user_number": {"maxLength": 5000, "type": "string", "description": "The service user number for the account on this mandate.", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "The URL that will contain the mandate that the customer has signed."}}, "description": "", "x-expandableFields": []}
```
