---
title: api-shield_old_validation_override_mitigation_action_write
page_id: schema-api-shield-old-validation-override-mitigation-action-write-0346bf46
path: schemas
description: |-
    When set, this overrides both zone level and operation level mitigation actions.

      - `none` will skip running schema validation entirely for the request
      - `null` indicates that no override is in place

    To clear any override, use the special value `disable_override` or `null`
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_validation_override_mitigation_action_write

When set, this overrides both zone level and operation level mitigation actions.

  - `none` will skip running schema validation entirely for the request
  - `null` indicates that no override is in place

To clear any override, use the special value `disable_override` or `null`

```yaml
{"description": "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `none` will skip running schema validation entirely for the request\n  - `null` indicates that no override is in place\n\nTo clear any override, use the special value `disable_override` or `null`\n", "type": "string", "example": "none", "enum": ["none", "disable_override", null], "nullable": true, "x-auditable": true}
```
