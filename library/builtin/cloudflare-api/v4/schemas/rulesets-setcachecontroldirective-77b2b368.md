---
title: rulesets_SetCacheControlDirective
page_id: schema-rulesets-setcachecontroldirective-77b2b368
path: schemas
description: A cache-control directive configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheControlDirective

A cache-control directive configuration.

```yaml
{"description": "A cache-control directive configuration.", "oneOf": [{"description": "Set the directive.", "properties": {"cloudflare_only": {"$ref": "#/components/schemas/rulesets_SetCacheControlCloudflareOnly"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_SetCacheControlOperation"}, {"enum": ["set"]}]}}, "required": ["operation"], "title": "Set Directive", "type": "object"}, {"description": "Remove the directive.", "properties": {"cloudflare_only": {"$ref": "#/components/schemas/rulesets_SetCacheControlCloudflareOnly"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_SetCacheControlOperation"}, {"enum": ["remove"]}]}}, "required": ["operation"], "title": "Remove Directive", "type": "object"}], "title": "Cache Control Directive"}
```
