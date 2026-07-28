---
title: api-shield_response_expression_templates_fallthrough
page_id: schema-api-shield-response-expression-templates-fallthrough-f6ead379
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_response_expression_templates_fallthrough

```yaml
{"type": "object", "properties": {"expression": {"description": "WAF Expression for fallthrough", "type": "string", "example": "(cf.api_gateway.fallthrough_detected)", "x-auditable": true}, "title": {"description": "Title for the expression", "type": "string", "example": "Fallthrough Expression for [zone.domain.tld]", "x-auditable": true}}, "required": ["title", "expression"]}
```
