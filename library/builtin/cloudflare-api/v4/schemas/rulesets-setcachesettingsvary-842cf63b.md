---
title: rulesets_SetCacheSettingsVary
page_id: schema-rulesets-setcachesettingsvary-842cf63b
path: schemas
description: Controls how cached responses vary based on request headers. `default` is required by the API and applies to any Vary response header that does not have a per-header override.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsVary

Controls how cached responses vary based on request headers. `default` is required by the API and applies to any Vary response header that does not have a per-header override.

```yaml
{"description": "Controls how cached responses vary based on request headers. `default` is required by the API and applies to any Vary response header that does not have a per-header override.", "type": "object", "properties": {"default": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsVaryDefault"}, "headers": {"description": "A mapping of lowercase request header names to their vary configuration.", "type": "object", "example": {"accept": {"action": "normalize", "media_types": ["image/webp", "image/png"]}}, "additionalProperties": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsVaryHeader"}, "maxProperties": 50, "title": "Headers"}}, "minProperties": 1, "title": "Vary"}
```
