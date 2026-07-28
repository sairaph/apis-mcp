---
title: email_rule_action
page_id: schema-email-rule-action-b5e8728d
path: schemas
description: Actions pattern.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_rule_action

Actions pattern.

```yaml
{"description": "Actions pattern.", "type": "object", "properties": {"type": {"description": "Type of supported action.", "type": "string", "example": "forward", "enum": ["drop", "forward", "worker"], "x-auditable": true}, "value": {"type": "array", "items": {"description": "Value for action.", "example": "destinationaddress@example.net", "maxLength": 90, "type": "string", "x-auditable": true}, "maxItems": 1}}, "required": ["type"]}
```
