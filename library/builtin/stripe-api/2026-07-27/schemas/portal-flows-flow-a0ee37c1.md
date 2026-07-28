---
title: portal_flows_flow
page_id: schema-portal-flows-flow-a0ee37c1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_flows_flow

```yaml
{"title": "PortalFlowsFlow", "required": ["after_completion", "type"], "type": "object", "properties": {"after_completion": {"$ref": "#/components/schemas/portal_flows_flow_after_completion"}, "subscription_cancel": {"description": "Configuration when `flow.type=subscription_cancel`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_flow_subscription_cancel"}]}, "subscription_update": {"description": "Configuration when `flow.type=subscription_update`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_flow_subscription_update"}]}, "subscription_update_confirm": {"description": "Configuration when `flow.type=subscription_update_confirm`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_flow_subscription_update_confirm"}]}, "type": {"type": "string", "description": "Type of flow that the customer will go through.", "enum": ["payment_method_update", "subscription_cancel", "subscription_update", "subscription_update_confirm"]}}, "description": "", "x-expandableFields": ["after_completion", "subscription_cancel", "subscription_update", "subscription_update_confirm"]}
```
