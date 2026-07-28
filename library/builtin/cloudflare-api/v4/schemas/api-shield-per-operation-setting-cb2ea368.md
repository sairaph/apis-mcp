---
title: api-shield_per_operation_setting
page_id: schema-api-shield-per-operation-setting-cb2ea368
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_per_operation_setting

```yaml
{"type": "object", "properties": {"mitigation_action": {"description": "When set, this applies a mitigation action to this operation which supersedes a global schema validation setting just for this operation\n\n  - `\"log\"` - log request when request does not conform to schema for this operation\n  - `\"block\"` - deny access to the site when request does not conform to schema for this operation\n  - `\"none\"` - will skip mitigation for this operation\n", "type": "string", "example": "block", "enum": ["log", "block", "none"], "x-auditable": true}, "operation_id": {"$ref": "#/components/schemas/api-shield_uuid-2"}}, "required": ["operation_id", "mitigation_action"]}
```
