---
title: cloudforce-one_RuleMetaEntry
page_id: schema-cloudforce-one-rulemetaentry-60750af6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_RuleMetaEntry

```yaml
{"type": "object", "properties": {"key": {"type": "string"}, "type": {"type": "string", "enum": ["string", "bool", "int"]}, "value": {"anyOf": [{"type": "string"}, {"type": "boolean"}, {"type": "number"}]}}, "required": ["key", "value", "type"]}
```
