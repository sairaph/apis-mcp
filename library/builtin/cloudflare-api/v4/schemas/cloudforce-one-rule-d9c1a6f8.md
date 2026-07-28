---
title: cloudforce-one_Rule
page_id: schema-cloudforce-one-rule-d9c1a6f8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_Rule

```yaml
{"type": "object", "properties": {"content": {"example": "rule example { condition: true }", "type": "string"}, "created_at": {"type": "number", "example": 1679529600000, "x-auditable": true}, "created_by": {"type": "string", "example": "user@example.com", "x-auditable": true}, "description": {"type": "string", "example": "Detects malicious proxy workers"}, "enabled": {"description": "Whether this rule is active for dice consumers.", "type": "boolean", "example": true}, "id": {"type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000", "x-auditable": true}, "is_public": {"description": "Whether this rule is visible to other internal accounts.", "type": "boolean", "example": false, "x-auditable": true}, "meta": {"description": "Structured meta entries for the rule (parsed from content plus any request-supplied meta). Returned in source order.", "type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_RuleMetaEntry"}}, "name": {"type": "string", "example": "block-malicious-workers", "x-auditable": true}, "namespaces": {"type": "array", "items": {"type": "string"}, "example": ["yara/workers"], "x-auditable": true}, "path": {"type": "string", "example": "yara/workers", "x-auditable": true}, "pending_approval_id": {"description": "ID of an open approval workflow targeting this rule, or null if none is pending.", "type": "number", "nullable": true}, "structured_source": {"description": "Original JSON payload for rules created via the structured rules API. Null for hand-written rules.", "type": "string", "nullable": true}, "updated_at": {"type": "number", "example": 1679529600000, "x-auditable": true}, "updated_by": {"type": "string", "example": "user@example.com", "x-auditable": true}}, "required": ["id", "name", "path", "description", "namespaces", "content", "is_public", "enabled", "created_at", "updated_at", "created_by", "updated_by", "pending_approval_id"]}
```
