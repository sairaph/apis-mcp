---
title: rulesets_SetCacheControlDirectiveWithValue
page_id: schema-rulesets-setcachecontroldirectivewithvalue-91879418
path: schemas
description: A cache-control directive configuration that accepts a duration value in seconds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheControlDirectiveWithValue

A cache-control directive configuration that accepts a duration value in seconds.

```yaml
{"description": "A cache-control directive configuration that accepts a duration value in seconds.", "oneOf": [{"description": "Set the directive with a duration value in seconds.", "properties": {"cloudflare_only": {"$ref": "#/components/schemas/rulesets_SetCacheControlCloudflareOnly"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_SetCacheControlOperation"}, {"enum": ["set"]}]}, "value": {"description": "The duration value in seconds for the directive.", "type": "integer", "example": 3600, "minimum": 0, "title": "Value"}}, "required": ["operation", "value"], "title": "Set Directive", "type": "object"}, {"description": "Remove the directive.", "properties": {"cloudflare_only": {"$ref": "#/components/schemas/rulesets_SetCacheControlCloudflareOnly"}, "operation": {"allOf": [{"$ref": "#/components/schemas/rulesets_SetCacheControlOperation"}, {"enum": ["remove"]}]}}, "required": ["operation"], "title": "Remove Directive", "type": "object"}], "title": "Cache Control Directive With Duration Seconds Value"}
```
