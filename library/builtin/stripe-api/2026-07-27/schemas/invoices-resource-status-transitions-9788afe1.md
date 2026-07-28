---
title: invoices_resource_status_transitions
page_id: schema-invoices-resource-status-transitions-9788afe1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_resource_status_transitions

```yaml
{"title": "InvoicesResourceStatusTransitions", "type": "object", "properties": {"finalized_at": {"type": "integer", "description": "The time that the invoice draft was finalized.", "format": "unix-time", "nullable": true}, "marked_uncollectible_at": {"type": "integer", "description": "The time that the invoice was marked uncollectible.", "format": "unix-time", "nullable": true}, "paid_at": {"type": "integer", "description": "The time that the invoice was paid.", "format": "unix-time", "nullable": true}, "voided_at": {"type": "integer", "description": "The time that the invoice was voided.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
