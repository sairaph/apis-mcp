---
title: firewall_rule
page_id: schema-firewall-rule-5fec3f9b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_rule

```yaml
{"type": "object", "properties": {"allowed_modes": {"description": "The available actions that a rule can apply to a matched request.", "type": "array", "items": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "example": ["whitelist", "block", "challenge", "js_challenge", "managed_challenge"], "readOnly": true}, "configuration": {"$ref": "#/components/schemas/firewall_configuration"}, "created_on": {"description": "The timestamp of when the rule was created.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}, "id": {"$ref": "#/components/schemas/firewall_schemas-identifier"}, "mode": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "modified_on": {"description": "The timestamp of when the rule was last modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}, "notes": {"$ref": "#/components/schemas/firewall_notes"}}, "required": ["id", "mode", "allowed_modes", "configuration"]}
```
