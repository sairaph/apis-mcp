---
title: rulesets_SetCacheSettingsVaryHeader
page_id: schema-rulesets-setcachesettingsvaryheader-01d7267f
path: schemas
description: Controls how a single request header contributes to the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsVaryHeader

Controls how a single request header contributes to the cache key.

```yaml
{"description": "Controls how a single request header contributes to the cache key.", "type": "object", "properties": {"action": {"description": "How the header value is treated when building the cache key.", "type": "string", "example": "normalize", "enum": ["bypass", "passthrough", "normalize"], "title": "Action"}, "languages": {"description": "The set of languages to normalize against. Only valid for the `accept-language` header.", "type": "array", "items": {"description": "A language tag to normalize against. Must be printable ASCII.", "example": "en", "maxLength": 64, "minLength": 1, "pattern": "^[ -~]+$", "title": "Language", "type": "string"}, "maxItems": 20, "title": "Languages"}, "media_types": {"description": "The set of media types to normalize against. Only valid for the `accept` header.", "type": "array", "items": {"description": "A media type to normalize against. Must be printable ASCII.", "example": "image/webp", "maxLength": 255, "minLength": 1, "pattern": "^[ -~]+$", "title": "Media Type", "type": "string"}, "maxItems": 10, "title": "Media Types"}}, "required": ["action"], "title": "Vary Header"}
```
