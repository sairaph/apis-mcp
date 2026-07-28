---
title: portal_features
page_id: schema-portal-features-3e803367
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_features

```yaml
{"title": "PortalFeatures", "required": ["customer_update", "invoice_history", "payment_method_update", "subscription_cancel", "subscription_update"], "type": "object", "properties": {"customer_update": {"$ref": "#/components/schemas/portal_customer_update"}, "invoice_history": {"$ref": "#/components/schemas/portal_invoice_list"}, "payment_method_update": {"$ref": "#/components/schemas/portal_payment_method_update"}, "subscription_cancel": {"$ref": "#/components/schemas/portal_subscription_cancel"}, "subscription_update": {"$ref": "#/components/schemas/portal_subscription_update"}}, "description": "", "x-expandableFields": ["customer_update", "invoice_history", "payment_method_update", "subscription_cancel", "subscription_update"]}
```
