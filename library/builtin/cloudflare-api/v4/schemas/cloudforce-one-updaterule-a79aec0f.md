---
title: cloudforce-one_UpdateRule
page_id: schema-cloudforce-one-updaterule-a79aec0f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_UpdateRule

```yaml
{"type": "object", "properties": {"commit_message": {"description": "Human-readable justification for this change. Required for internal-account submissions; optional for customer accounts and automated sync.", "type": "string", "example": "Reduce false positives on legit CI workers.", "maxLength": 1000}, "content": {"example": "rule example { condition: true }", "minLength": 1, "type": "string"}, "description": {"description": "Human-readable description of the rule. Auto-extracted from YARA meta if present.", "type": "string", "example": "Detects malicious proxy workers", "maxLength": 1000}, "enabled": {"description": "Whether this rule is active for dice consumers.", "type": "boolean", "example": true}, "is_public": {"description": "Whether this rule is visible to other internal accounts.", "type": "boolean", "example": false, "x-auditable": true}, "meta": {"description": "Additional YARA meta entries appended to the rule's meta block (and stored in rule_meta alongside meta parsed from the content). Keys must be valid YARA identifiers and must not be 'name', 'enabled', or 'description'. Duplicate keys are allowed.", "type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_MetaInputEntry"}, "maxItems": 50}, "name": {"type": "string", "example": "block-malicious-workers", "maxLength": 255, "minLength": 1, "x-auditable": true}, "namespaces": {"type": "array", "items": {"maxLength": 255, "minLength": 1, "type": "string"}, "example": ["yara/workers"], "x-auditable": true}, "path": {"description": "Path change goes through approval workflow.", "type": "string", "example": "yara/workers", "minLength": 1, "x-auditable": true}}}
```
