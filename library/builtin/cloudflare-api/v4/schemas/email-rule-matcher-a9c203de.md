---
title: email_rule_matcher
page_id: schema-email-rule-matcher-a9c203de
path: schemas
description: Matching pattern to forward your actions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_rule_matcher

Matching pattern to forward your actions.

```yaml
{"description": "Matching pattern to forward your actions.", "type": "object", "properties": {"field": {"description": "Field for type matcher.", "type": "string", "example": "to", "enum": ["to"], "x-auditable": true}, "type": {"description": "Type of matcher.", "type": "string", "example": "literal", "enum": ["all", "literal"], "x-auditable": true}, "value": {"description": "Value for matcher.", "type": "string", "example": "test@example.com", "maxLength": 90, "x-auditable": true}}, "required": ["type"]}
```
