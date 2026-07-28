---
title: api-shield_global_setting_change_base
page_id: schema-api-shield-global-setting-change-base-6d52d898
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_global_setting_change_base

```yaml
{"type": "object", "properties": {"validation_default_mitigation_action": {"description": "The default mitigation action used\nMitigation actions are as follows:\n\n  - `\"log\"` - log request when request does not conform to schema\n  - `\"block\"` - deny access to the site when request does not conform to schema\n  - `\"none\"` - skip running schema validation\n", "type": "string", "example": "block", "enum": ["none", "log", "block"], "x-auditable": true}, "validation_override_mitigation_action": {"description": "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `\"none\"` - skip running schema validation entirely for the request\n  - `null` - clears any existing override\n", "type": "string", "enum": ["none", null], "nullable": true, "x-auditable": true}}}
```
