---
title: payment_method_options_card_mandate_options
page_id: schema-payment-method-options-card-mandate-options-d6701855
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_options_card_mandate_options

```yaml
{"title": "payment_method_options_card_mandate_options", "required": ["amount", "amount_type", "interval", "reference", "start_date"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to be charged for future payments, specified in the presentment currency."}, "amount_type": {"type": "string", "description": "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.", "enum": ["fixed", "maximum"]}, "description": {"maxLength": 200, "type": "string", "description": "A description of the mandate or subscription that is meant to be displayed to the customer.", "nullable": true}, "end_date": {"type": "integer", "description": "End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.", "format": "unix-time", "nullable": true}, "interval": {"type": "string", "description": "Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.", "enum": ["day", "month", "sporadic", "week", "year"]}, "interval_count": {"type": "integer", "description": "The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.", "nullable": true}, "reference": {"maxLength": 80, "type": "string", "description": "Unique identifier for the mandate or subscription."}, "start_date": {"type": "integer", "description": "Start date of the mandate or subscription. Start date should not be lesser than yesterday.", "format": "unix-time"}, "supported_types": {"type": "array", "description": "Specifies the type of mandates supported. Possible values are `india`.", "nullable": true, "items": {"type": "string", "enum": ["india"]}}}, "description": "", "x-expandableFields": []}
```
