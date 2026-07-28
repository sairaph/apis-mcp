---
title: refund_next_action
page_id: schema-refund-next-action-8b501f0b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# refund_next_action

```yaml
{"title": "RefundNextAction", "required": ["type"], "type": "object", "properties": {"display_details": {"$ref": "#/components/schemas/refund_next_action_display_details"}, "type": {"maxLength": 5000, "type": "string", "description": "Type of the next action to perform."}}, "description": "", "x-expandableFields": ["display_details"]}
```
