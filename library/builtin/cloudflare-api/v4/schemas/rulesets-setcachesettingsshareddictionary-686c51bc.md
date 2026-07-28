---
title: rulesets_SetCacheSettingsSharedDictionary
page_id: schema-rulesets-setcachesettingsshareddictionary-686c51bc
path: schemas
description: Configuration for shared dictionary compression. When set, Cloudflare injects Use-As-Dictionary headers on matching cacheable responses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsSharedDictionary

Configuration for shared dictionary compression. When set, Cloudflare injects Use-As-Dictionary headers on matching cacheable responses.

```yaml
{"description": "Configuration for shared dictionary compression. When set, Cloudflare injects Use-As-Dictionary headers on matching cacheable responses.", "type": "object", "properties": {"match_pattern": {"description": "URL pattern for the Use-As-Dictionary match field. This pattern specifies which URLs can use this response as a dictionary.", "type": "string", "example": "/static/js/*.js", "maxLength": 1024, "minLength": 1, "title": "Match Pattern"}}, "required": ["match_pattern"], "title": "Shared Dictionary"}
```
