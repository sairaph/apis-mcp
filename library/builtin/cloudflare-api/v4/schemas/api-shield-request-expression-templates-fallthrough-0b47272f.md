---
title: api-shield_request_expression_templates_fallthrough
page_id: schema-api-shield-request-expression-templates-fallthrough-0b47272f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_request_expression_templates_fallthrough

```yaml
{"type": "object", "properties": {"hosts": {"description": "List of hosts to be targeted in the expression", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["{zone}.domain1.tld", "domain2.tld"]}}, "additionalProperties": false, "required": ["hosts"]}
```
