---
title: platform_earning_fee_source
page_id: schema-platform-earning-fee-source-51c26b6d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# platform_earning_fee_source

```yaml
{"title": "PlatformEarningFeeSource", "required": ["type"], "type": "object", "properties": {"charge": {"maxLength": 5000, "type": "string", "description": "Charge ID that created this application fee."}, "payout": {"maxLength": 5000, "type": "string", "description": "Payout ID that created this application fee."}, "type": {"type": "string", "description": "Type of object that created the application fee.", "enum": ["charge", "payout"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
