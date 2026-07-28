---
title: portal_customer_update
page_id: schema-portal-customer-update-16f61eee
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_customer_update

```yaml
{"title": "PortalCustomerUpdate", "required": ["allowed_updates", "enabled"], "type": "object", "properties": {"allowed_updates": {"type": "array", "description": "The types of customer updates that are supported. When empty, customers are not updateable.", "items": {"type": "string", "enum": ["address", "email", "name", "phone", "shipping", "tax_id"]}}, "enabled": {"type": "boolean", "description": "Whether the feature is enabled."}}, "description": "", "x-expandableFields": []}
```
