---
title: email_rule_catchall-action
page_id: schema-email-rule-catchall-action-01a39b25
path: schemas
description: Action for the catch-all routing rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_rule_catchall-action

Action for the catch-all routing rule.

```yaml
{"description": "Action for the catch-all routing rule.", "type": "object", "properties": {"type": {"description": "Type of action for catch-all rule.", "type": "string", "example": "forward", "enum": ["drop", "forward", "worker"], "x-auditable": true}, "value": {"type": "array", "items": {"description": "Input value for action.", "example": "destinationaddress@example.net", "maxLength": 90, "type": "string", "x-auditable": true}, "maxItems": 1}}, "required": ["type"]}
```
