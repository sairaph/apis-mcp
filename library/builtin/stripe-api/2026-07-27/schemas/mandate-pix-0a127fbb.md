---
title: mandate_pix
page_id: schema-mandate-pix-0a127fbb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_pix

```yaml
{"title": "mandate_pix", "type": "object", "properties": {"amount_includes_iof": {"type": "string", "description": "Determines if the amount includes the IOF tax.", "enum": ["always", "never"]}, "amount_type": {"type": "string", "description": "Type of amount.", "enum": ["fixed", "maximum"]}, "end_date": {"maxLength": 5000, "type": "string", "description": "Date when the mandate expires and no further payments will be charged, in `YYYY-MM-DD`."}, "payment_schedule": {"type": "string", "description": "Schedule at which the future payments will be charged.", "enum": ["halfyearly", "monthly", "quarterly", "weekly", "yearly"]}, "reference": {"maxLength": 5000, "type": "string", "description": "Subscription name displayed to buyers in their bank app."}, "start_date": {"maxLength": 5000, "type": "string", "description": "Start date of the mandate, in `YYYY-MM-DD`."}}, "description": "", "x-expandableFields": []}
```
