---
title: zones_priority
page_id: schema-zones-priority-bccbf978
path: schemas
description: |-
    The priority of the rule, used to define which Page Rule is processed
    over another. A higher number indicates a higher priority. For example,
    if you have a catch-all Page Rule (rule A: `/images/*`) but want a more
    specific Page Rule to take precedence (rule B: `/images/special/*`),
    specify a higher priority for rule B so it overrides rule A.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_priority

The priority of the rule, used to define which Page Rule is processed
over another. A higher number indicates a higher priority. For example,
if you have a catch-all Page Rule (rule A: `/images/*`) but want a more
specific Page Rule to take precedence (rule B: `/images/special/*`),
specify a higher priority for rule B so it overrides rule A.

```yaml
{"description": "The priority of the rule, used to define which Page Rule is processed\nover another. A higher number indicates a higher priority. For example,\nif you have a catch-all Page Rule (rule A: `/images/*`) but want a more\nspecific Page Rule to take precedence (rule B: `/images/special/*`),\nspecify a higher priority for rule B so it overrides rule A.\n", "type": "integer", "default": 1, "x-auditable": true}
```
