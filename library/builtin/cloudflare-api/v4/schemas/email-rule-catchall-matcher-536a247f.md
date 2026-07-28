---
title: email_rule_catchall-matcher
page_id: schema-email-rule-catchall-matcher-536a247f
path: schemas
description: Matcher for catch-all routing rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_rule_catchall-matcher

Matcher for catch-all routing rule.

```yaml
{"description": "Matcher for catch-all routing rule.", "type": "object", "properties": {"type": {"description": "Type of matcher. Default is 'all'.", "type": "string", "example": "all", "enum": ["all"], "x-auditable": true}}, "required": ["type"]}
```
