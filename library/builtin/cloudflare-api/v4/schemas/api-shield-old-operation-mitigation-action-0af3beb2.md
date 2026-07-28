---
title: api-shield_old_operation_mitigation_action
page_id: schema-api-shield-old-operation-mitigation-action-0af3beb2
path: schemas
description: |-
    When set, this applies a mitigation action to this operation

      - `log` log request when request does not conform to schema for this operation
      - `block` deny access to the site when request does not conform to schema for this operation
      - `none` will skip mitigation for this operation
      - `null` indicates that no operation level mitigation is in place, see Zone Level Schema Validation Settings for mitigation action that will be applied
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_operation_mitigation_action

When set, this applies a mitigation action to this operation

  - `log` log request when request does not conform to schema for this operation
  - `block` deny access to the site when request does not conform to schema for this operation
  - `none` will skip mitigation for this operation
  - `null` indicates that no operation level mitigation is in place, see Zone Level Schema Validation Settings for mitigation action that will be applied

```yaml
{"description": "When set, this applies a mitigation action to this operation\n\n  - `log` log request when request does not conform to schema for this operation\n  - `block` deny access to the site when request does not conform to schema for this operation\n  - `none` will skip mitigation for this operation\n  - `null` indicates that no operation level mitigation is in place, see Zone Level Schema Validation Settings for mitigation action that will be applied\n", "type": "string", "example": "block", "enum": ["log", "block", "none", null], "nullable": true, "x-auditable": true}
```
