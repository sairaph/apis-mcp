---
title: refund_destination_details_swish
page_id: schema-refund-destination-details-swish-7c80d7f2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# refund_destination_details_swish

```yaml
{"title": "refund_destination_details_swish", "type": "object", "properties": {"network_decline_code": {"maxLength": 5000, "type": "string", "description": "For refunds declined by the network, a decline code provided by the network which indicates the reason the refund failed.", "nullable": true}, "reference": {"maxLength": 5000, "type": "string", "description": "The reference assigned to the refund.", "nullable": true}, "reference_status": {"maxLength": 5000, "type": "string", "description": "Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
