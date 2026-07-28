---
title: source_mandate_notification_sepa_debit_data
page_id: schema-source-mandate-notification-sepa-debit-data-63e7f04e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_mandate_notification_sepa_debit_data

```yaml
{"title": "SourceMandateNotificationSepaDebitData", "type": "object", "properties": {"creditor_identifier": {"maxLength": 5000, "type": "string", "description": "SEPA creditor ID."}, "last4": {"maxLength": 5000, "type": "string", "description": "Last 4 digits of the account number associated with the debit."}, "mandate_reference": {"maxLength": 5000, "type": "string", "description": "Mandate reference associated with the debit."}}, "description": "", "x-expandableFields": []}
```
