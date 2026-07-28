---
title: api-shield_old_validation_default_mitigation_action_patch
page_id: schema-api-shield-old-validation-default-mitigation-action-patch-98cf8f90
path: schemas
description: |-
    The default mitigation action used when there is no mitigation action defined on the operation
    Mitigation actions are as follows:

      * `log` - log request when request does not conform to schema
      * `block` - deny access to the site when request does not conform to schema

    A special value of of `none` will skip running schema validation entirely for the request when there is no mitigation action defined on the operation

    `null` will have no effect.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_validation_default_mitigation_action_patch

The default mitigation action used when there is no mitigation action defined on the operation
Mitigation actions are as follows:

  * `log` - log request when request does not conform to schema
  * `block` - deny access to the site when request does not conform to schema

A special value of of `none` will skip running schema validation entirely for the request when there is no mitigation action defined on the operation

`null` will have no effect.

```yaml
{"description": "The default mitigation action used when there is no mitigation action defined on the operation\nMitigation actions are as follows:\n\n  * `log` - log request when request does not conform to schema\n  * `block` - deny access to the site when request does not conform to schema\n\nA special value of of `none` will skip running schema validation entirely for the request when there is no mitigation action defined on the operation\n\n`null` will have no effect.\n", "type": "string", "example": "block", "enum": ["none", "log", "block", null], "nullable": true, "x-auditable": true}
```
