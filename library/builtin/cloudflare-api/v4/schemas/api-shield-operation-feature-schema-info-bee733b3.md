---
title: api-shield_operation_feature_schema_info
page_id: schema-api-shield-operation-feature-schema-info-bee733b3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_feature_schema_info

```yaml
{"type": "object", "properties": {"schema_info": {"type": "object", "properties": {"active_schema": {"description": "Schema active on endpoint.", "type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/api-shield_timestamp"}, "id": {"$ref": "#/components/schemas/api-shield_uuid-2"}, "is_learned": {"description": "True if schema is Cloudflare-provided.", "type": "boolean", "example": true, "x-auditable": true}, "name": {"description": "Schema file name.", "type": "string", "example": "api-endpoints-8694824bf5c04d019edcbf399c03c103-api-discovery.example.com-thresholds.json", "x-auditable": true}}}, "learned_available": {"description": "Deprecated. Always false.", "type": "boolean", "example": false, "x-auditable": true}, "mitigation_action": {"description": "Action taken on requests failing validation.", "type": "string", "example": "block", "enum": ["none", "log", "block"], "nullable": true, "x-auditable": true}}}}, "readOnly": true}
```
