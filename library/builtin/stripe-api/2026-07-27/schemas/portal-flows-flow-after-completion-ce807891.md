---
title: portal_flows_flow_after_completion
page_id: schema-portal-flows-flow-after-completion-ce807891
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_flows_flow_after_completion

```yaml
{"title": "PortalFlowsFlowAfterCompletion", "required": ["type"], "type": "object", "properties": {"hosted_confirmation": {"description": "Configuration when `after_completion.type=hosted_confirmation`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_after_completion_hosted_confirmation"}]}, "redirect": {"description": "Configuration when `after_completion.type=redirect`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_after_completion_redirect"}]}, "type": {"type": "string", "description": "The specified type of behavior after the flow is completed.", "enum": ["hosted_confirmation", "portal_homepage", "redirect"]}}, "description": "", "x-expandableFields": ["hosted_confirmation", "redirect"]}
```
