---
title: portal_flows_flow_subscription_cancel
page_id: schema-portal-flows-flow-subscription-cancel-0c393f02
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_flows_flow_subscription_cancel

```yaml
{"title": "PortalFlowsFlowSubscriptionCancel", "required": ["subscription"], "type": "object", "properties": {"retention": {"description": "Specify a retention strategy to be used in the cancellation flow.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_retention"}]}, "subscription": {"maxLength": 5000, "type": "string", "description": "The ID of the subscription to be canceled."}}, "description": "", "x-expandableFields": ["retention"]}
```
