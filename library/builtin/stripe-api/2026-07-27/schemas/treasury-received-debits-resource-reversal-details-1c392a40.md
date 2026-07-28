---
title: treasury_received_debits_resource_reversal_details
page_id: schema-treasury-received-debits-resource-reversal-details-1c392a40
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_received_debits_resource_reversal_details

```yaml
{"title": "TreasuryReceivedDebitsResourceReversalDetails", "type": "object", "properties": {"deadline": {"type": "integer", "description": "Time before which a ReceivedDebit can be reversed.", "format": "unix-time", "nullable": true}, "restricted_reason": {"type": "string", "description": "Set if a ReceivedDebit can't be reversed.", "nullable": true, "enum": ["already_reversed", "deadline_passed", "network_restricted", "other", "source_flow_restricted"]}}, "description": "", "x-expandableFields": []}
```
