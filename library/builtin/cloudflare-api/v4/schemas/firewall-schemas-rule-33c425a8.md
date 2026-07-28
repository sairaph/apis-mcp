---
title: firewall_schemas-rule
page_id: schema-firewall-schemas-rule-33c425a8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_schemas-rule

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/firewall_rule"}, {"properties": {"scope": {"description": "All zones owned by the user will have the rule applied.", "type": "object", "properties": {"email": {"$ref": "#/components/schemas/firewall_email"}, "id": {"$ref": "#/components/schemas/firewall_identifier"}, "type": {"description": "Defines the scope of the rule.", "type": "string", "example": "user", "enum": ["user", "organization"], "readOnly": true}}, "readOnly": true}}, "type": "object"}], "required": ["id", "mode", "allowed_modes", "configuration", "scope"]}
```
