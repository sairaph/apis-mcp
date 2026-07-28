---
title: source_receiver_flow
page_id: schema-source-receiver-flow-a40600b7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_receiver_flow

```yaml
{"title": "SourceReceiverFlow", "required": ["amount_charged", "amount_received", "amount_returned", "refund_attributes_method", "refund_attributes_status"], "type": "object", "properties": {"address": {"maxLength": 5000, "type": "string", "description": "The address of the receiver source. This is the value that should be communicated to the customer to send their funds to.", "nullable": true}, "amount_charged": {"type": "integer", "description": "The total amount that was moved to your balance. This is almost always equal to the amount charged. In rare cases when customers deposit excess funds and we are unable to refund those, those funds get moved to your balance and show up in amount_charged as well. The amount charged is expressed in the source's currency."}, "amount_received": {"type": "integer", "description": "The total amount received by the receiver source. `amount_received = amount_returned + amount_charged` should be true for consumed sources unless customers deposit excess funds. The amount received is expressed in the source's currency."}, "amount_returned": {"type": "integer", "description": "The total amount that was returned to the customer. The amount returned is expressed in the source's currency."}, "refund_attributes_method": {"maxLength": 5000, "type": "string", "description": "Type of refund attribute method, one of `email`, `manual`, or `none`."}, "refund_attributes_status": {"maxLength": 5000, "type": "string", "description": "Type of refund attribute status, one of `missing`, `requested`, or `available`."}}, "description": "", "x-expandableFields": []}
```
