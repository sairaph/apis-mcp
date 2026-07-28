---
title: rulesets_SetCacheControlDirectiveWithQualifiers
page_id: schema-rulesets-setcachecontroldirectivewithqualifiers-73e017e2
path: schemas
description: A cache-control directive configuration that accepts optional qualifiers (header names).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheControlDirectiveWithQualifiers

A cache-control directive configuration that accepts optional qualifiers (header names).

```yaml
{"description": "A cache-control directive configuration that accepts optional qualifiers (header names).", "oneOf": [{"description": "Set the directive with optional qualifiers.", "properties": {"cloudflare_only": {"$ref": "#/components/schemas/rulesets_SetCacheControlCloudflareOnly"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_SetCacheControlOperation"}, {"enum": ["set"]}]}, "qualifiers": {"description": "Optional list of header names to qualify the directive (e.g., for \"private\" or \"no-cache\" directives).", "type": "array", "items": {"description": "A header name to qualify the directive.", "example": "X-Custom-Header", "minLength": 1, "title": "Qualifier", "type": "string"}, "title": "Qualifiers", "uniqueItems": true}}, "required": ["operation"], "title": "Set Directive", "type": "object"}, {"description": "Remove the directive.", "properties": {"cloudflare_only": {"$ref": "#/components/schemas/rulesets_SetCacheControlCloudflareOnly"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_SetCacheControlOperation"}, {"enum": ["remove"]}]}}, "required": ["operation"], "title": "Remove Directive", "type": "object"}], "title": "Cache Control Directive With Optional Qualifiers"}
```
