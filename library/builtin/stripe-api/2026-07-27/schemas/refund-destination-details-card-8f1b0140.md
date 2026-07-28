---
title: refund_destination_details_card
page_id: schema-refund-destination-details-card-8f1b0140
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# refund_destination_details_card

```yaml
{"title": "refund_destination_details_card", "required": ["type"], "type": "object", "properties": {"reference": {"maxLength": 5000, "type": "string", "description": "Value of the reference number assigned to the refund."}, "reference_status": {"maxLength": 5000, "type": "string", "description": "Status of the reference number on the refund. This can be `pending`, `available` or `unavailable`."}, "reference_type": {"maxLength": 5000, "type": "string", "description": "Type of the reference number assigned to the refund."}, "type": {"type": "string", "description": "The type of refund. This can be `refund`, `reversal`, or `pending`.", "enum": ["pending", "refund", "reversal"]}}, "description": "", "x-expandableFields": []}
```
