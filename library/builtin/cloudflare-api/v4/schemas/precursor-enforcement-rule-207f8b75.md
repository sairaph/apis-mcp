---
title: precursor_enforcement_rule
page_id: schema-precursor-enforcement-rule-207f8b75
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# precursor_enforcement_rule

```yaml
{"type": "object", "properties": {"description": {"description": "An informative description of the rule.", "type": "string", "example": "Ease friction on the login path", "default": "", "x-auditable": true}, "enabled": {"description": "Whether the rule is active.", "type": "boolean", "example": true, "default": true, "x-auditable": true}, "expression": {"description": "The filter expression that determines which requests the rule matches.", "type": "string", "example": "http.request.uri.path eq \"/login\"", "x-auditable": true}, "id": {"description": "The read-only identifier that Cloudflare assigns to the rule.", "type": "string", "example": "3a03d665bac043e3a684e0d385a4b1e2", "readOnly": true, "x-auditable": true}, "mode": {"$ref": "#/components/schemas/precursor_enforcement_rule_mode"}}, "required": ["expression", "mode"], "title": "Enforcement Rule"}
```
