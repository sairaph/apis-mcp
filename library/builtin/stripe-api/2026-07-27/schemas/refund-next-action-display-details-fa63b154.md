---
title: refund_next_action_display_details
page_id: schema-refund-next-action-display-details-fa63b154
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# refund_next_action_display_details

```yaml
{"title": "RefundNextActionDisplayDetails", "required": ["email_sent", "expires_at"], "type": "object", "properties": {"email_sent": {"$ref": "#/components/schemas/email_sent"}, "expires_at": {"type": "integer", "description": "The expiry timestamp.", "format": "unix-time"}}, "description": "", "x-expandableFields": ["email_sent"]}
```
