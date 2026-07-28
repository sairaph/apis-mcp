---
title: zaraz_load-rule
page_id: schema-zaraz-load-rule-9a578830
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_load-rule

```yaml
{"type": "object", "properties": {"id": {"type": "string", "x-auditable": true}, "match": {"type": "string", "x-auditable": true}, "op": {"type": "string", "enum": ["CONTAINS", "EQUALS", "STARTS_WITH", "ENDS_WITH", "MATCH_REGEX", "NOT_MATCH_REGEX", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL"], "x-auditable": true}, "value": {"type": "string", "x-auditable": true}}, "required": ["id", "match", "op", "value"]}
```
