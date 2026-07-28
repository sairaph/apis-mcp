---
title: zaraz_base-mc
page_id: schema-zaraz-base-mc-2ebe640d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_base-mc

```yaml
{"allOf": [{"$ref": "#/components/schemas/zaraz_base-tool"}, {"properties": {"actions": {"description": "Actions configured on a tool. Either this or neoEvents field is required.", "type": "object", "additionalProperties": {"properties": {"actionType": {"description": "Tool event type.", "type": "string", "x-auditable": true}, "blockingTriggers": {"description": "List of blocking triggers IDs.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "data": {"description": "Event payload.", "type": "object"}, "firingTriggers": {"description": "List of firing triggers IDs.", "type": "array", "items": {"type": "string", "x-auditable": true}, "minItems": 1}}, "required": ["actionType", "blockingTriggers", "firingTriggers", "data"], "type": "object"}}, "component": {"description": "Tool's internal name.", "type": "string", "x-auditable": true}, "neoEvents": {"description": "DEPRECATED - List of actions configured on a tool. Either this or actions field is required. If both are present, actions field will take precedence.", "type": "array", "items": {"properties": {"actionType": {"description": "Tool event type.", "type": "string", "x-auditable": true}, "blockingTriggers": {"description": "List of blocking triggers IDs.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "data": {"description": "Event payload.", "type": "object"}, "firingTriggers": {"description": "List of firing triggers IDs.", "type": "array", "items": {"type": "string", "x-auditable": true}, "minItems": 1}}, "required": ["actionType", "blockingTriggers", "firingTriggers", "data"], "type": "object"}}, "permissions": {"description": "List of permissions granted to the component.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "settings": {"description": "Tool's settings.", "type": "object", "additionalProperties": {"anyOf": [{"type": "string"}, {"type": "boolean"}], "x-auditable": true}}}, "required": ["component", "settings", "permissions"], "type": "object"}]}
```
