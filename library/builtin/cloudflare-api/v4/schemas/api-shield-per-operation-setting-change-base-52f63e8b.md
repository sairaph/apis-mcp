---
title: api-shield_per_operation_setting_change_base
page_id: schema-api-shield-per-operation-setting-change-base-52f63e8b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_per_operation_setting_change_base

```yaml
{"type": "object", "properties": {"mitigation_action": {"description": "When set, this applies a mitigation action to this operation\n\n  - `\"log\"` - log request when request does not conform to schema for this operation\n  - `\"block\"` - deny access to the site when request does not conform to schema for this operation\n  - `\"none\"` - will skip mitigation for this operation\n  - `null` - clears any mitigation action\n", "type": "string", "example": "block", "enum": ["log", "block", "none", null], "nullable": true, "x-auditable": true}}}
```
