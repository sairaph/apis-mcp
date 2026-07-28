---
title: api-shield_old_validation_override_mitigation_action
page_id: schema-api-shield-old-validation-override-mitigation-action-16f1b73b
path: schemas
description: |-
    When set, this overrides both zone level and operation level mitigation actions.

      - `none` will skip running schema validation entirely for the request
      - `null` indicates that no override is in place
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_validation_override_mitigation_action

When set, this overrides both zone level and operation level mitigation actions.

  - `none` will skip running schema validation entirely for the request
  - `null` indicates that no override is in place

```yaml
{"description": "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `none` will skip running schema validation entirely for the request\n  - `null` indicates that no override is in place\n", "type": "string", "example": "disable_override", "enum": ["none", null], "nullable": true, "x-auditable": true}
```
