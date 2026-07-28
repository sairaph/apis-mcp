---
title: zaraz_form-submission-rule
page_id: schema-zaraz-form-submission-rule-88606c1f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_form-submission-rule

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["formSubmission"], "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"selector": {"type": "string", "x-auditable": true}, "validate": {"type": "boolean", "x-auditable": true}}, "required": ["selector", "validate"]}}, "required": ["id", "action", "settings"]}
```
