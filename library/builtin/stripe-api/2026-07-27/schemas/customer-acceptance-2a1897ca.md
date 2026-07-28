---
title: customer_acceptance
page_id: schema-customer-acceptance-2a1897ca
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_acceptance

```yaml
{"title": "customer_acceptance", "required": ["type"], "type": "object", "properties": {"accepted_at": {"type": "integer", "description": "The time that the customer accepts the mandate.", "format": "unix-time", "nullable": true}, "offline": {"$ref": "#/components/schemas/offline_acceptance"}, "online": {"$ref": "#/components/schemas/online_acceptance"}, "type": {"type": "string", "description": "The mandate includes the type of customer acceptance information, such as: `online` or `offline`.", "enum": ["offline", "online"]}}, "description": "", "x-expandableFields": ["offline", "online"]}
```
