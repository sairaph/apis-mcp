---
title: rulesets_SetCacheControlRule
page_id: schema-rulesets-setcachecontrolrule-572fc593
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheControlRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["set_cache_control"]}, "action_parameters": {"minProperties": 1, "properties": {"immutable": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "max-age": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirectiveWithValue"}, "must-revalidate": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "must-understand": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "no-cache": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirectiveWithQualifiers"}, "no-store": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "no-transform": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "private": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirectiveWithQualifiers"}, "proxy-revalidate": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "public": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirective"}, "s-maxage": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirectiveWithValue"}, "stale-if-error": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirectiveWithValue"}, "stale-while-revalidate": {"$ref": "#/components/schemas/rulesets_SetCacheControlDirectiveWithValue"}}}, "description": {"example": "Modify the cache-control header directives in an Origin response."}}, "title": "Set Cache Control Rule"}]}
```
