---
title: api-shield_global_settings
page_id: schema-api-shield-global-settings-71677420
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_global_settings

```yaml
{"type": "object", "properties": {"validation_default_mitigation_action": {"description": "The default mitigation action used\n\nMitigation actions are as follows:\n\n  - `log` - log request when request does not conform to schema\n  - `block` - deny access to the site when request does not conform to schema\n  - `none` - skip running schema validation\n", "type": "string", "example": "block", "enum": ["none", "log", "block"], "x-auditable": true}, "validation_override_mitigation_action": {"description": "When not null, this overrides global both zone level and operation level mitigation actions. This can serve as a quick way to disable schema validation for the whole zone.\n\n  - `\"none\"` will skip running schema validation entirely for the request\n", "type": "string", "enum": ["none"], "x-auditable": true}}, "required": ["validation_default_mitigation_action"]}
```
