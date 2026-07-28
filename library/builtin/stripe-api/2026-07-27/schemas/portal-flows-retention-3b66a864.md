---
title: portal_flows_retention
page_id: schema-portal-flows-retention-3b66a864
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_flows_retention

```yaml
{"title": "PortalFlowsRetention", "required": ["type"], "type": "object", "properties": {"coupon_offer": {"description": "Configuration when `retention.type=coupon_offer`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_coupon_offer"}]}, "type": {"type": "string", "description": "Type of retention strategy that will be used.", "enum": ["coupon_offer"]}}, "description": "", "x-expandableFields": ["coupon_offer"]}
```
