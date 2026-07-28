---
title: precursor_default_mode
page_id: schema-precursor-default-mode-647a8c04
path: schemas
description: |-
    The zone-level Precursor enforcement mode applied to requests that do
    not match a more specific enforcement rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# precursor_default_mode

The zone-level Precursor enforcement mode applied to requests that do
not match a more specific enforcement rule.

```yaml
{"description": "The zone-level Precursor enforcement mode applied to requests that do\nnot match a more specific enforcement rule.\n", "type": "string", "example": "off", "default": "off", "enum": ["off", "min-friction", "max-security"], "x-auditable": true}
```
