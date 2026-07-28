---
title: quotes_resource_status_transitions
page_id: schema-quotes-resource-status-transitions-0bace310
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_status_transitions

```yaml
{"title": "QuotesResourceStatusTransitions", "type": "object", "properties": {"accepted_at": {"type": "integer", "description": "The time that the quote was accepted. Measured in seconds since Unix epoch.", "format": "unix-time", "nullable": true}, "canceled_at": {"type": "integer", "description": "The time that the quote was canceled. Measured in seconds since Unix epoch.", "format": "unix-time", "nullable": true}, "finalized_at": {"type": "integer", "description": "The time that the quote was finalized. Measured in seconds since Unix epoch.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
