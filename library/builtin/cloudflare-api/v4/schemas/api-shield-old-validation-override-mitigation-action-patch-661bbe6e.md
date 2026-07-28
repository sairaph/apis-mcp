---
title: api-shield_old_validation_override_mitigation_action_patch
page_id: schema-api-shield-old-validation-override-mitigation-action-patch-661bbe6e
path: schemas
description: |-
    When set, this overrides both zone level and operation level mitigation actions.

      - `none` will skip running schema validation entirely for the request

    To clear any override, use the special value `disable_override`

    `null` will have no effect.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_validation_override_mitigation_action_patch

When set, this overrides both zone level and operation level mitigation actions.

  - `none` will skip running schema validation entirely for the request

To clear any override, use the special value `disable_override`

`null` will have no effect.

```yaml
{"description": "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `none` will skip running schema validation entirely for the request\n\nTo clear any override, use the special value `disable_override`\n\n`null` will have no effect.\n", "type": "string", "example": "none", "enum": ["none", "disable_override", null], "nullable": true, "x-auditable": true}
```
