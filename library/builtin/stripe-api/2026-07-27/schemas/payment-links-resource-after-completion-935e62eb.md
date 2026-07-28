---
title: payment_links_resource_after_completion
page_id: schema-payment-links-resource-after-completion-935e62eb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_after_completion

```yaml
{"title": "PaymentLinksResourceAfterCompletion", "required": ["type"], "type": "object", "properties": {"hosted_confirmation": {"$ref": "#/components/schemas/payment_links_resource_completion_behavior_confirmation_page"}, "redirect": {"$ref": "#/components/schemas/payment_links_resource_completion_behavior_redirect"}, "type": {"type": "string", "description": "The specified behavior after the purchase is complete.", "enum": ["hosted_confirmation", "redirect"]}}, "description": "", "x-expandableFields": ["hosted_confirmation", "redirect"]}
```
