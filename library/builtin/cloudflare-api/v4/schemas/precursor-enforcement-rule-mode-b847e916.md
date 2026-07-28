---
title: precursor_enforcement_rule_mode
page_id: schema-precursor-enforcement-rule-mode-b847e916
path: schemas
description: |-
    The override mode Precursor applies to requests matching an enforcement
    rule. Unlike `default_mode`, this cannot be `off`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# precursor_enforcement_rule_mode

The override mode Precursor applies to requests matching an enforcement
rule. Unlike `default_mode`, this cannot be `off`.

```yaml
{"description": "The override mode Precursor applies to requests matching an enforcement\nrule. Unlike `default_mode`, this cannot be `off`.\n", "type": "string", "example": "max-security", "enum": ["min-friction", "max-security"], "x-auditable": true}
```
