---
title: rulesets_SetCacheSettingsRule
page_id: schema-rulesets-setcachesettingsrule-6c30e532
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["set_cache_settings"]}, "action_parameters": {"minProperties": 1, "properties": {"additional_cacheable_ports": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsAdditionalCacheablePorts"}, "browser_ttl": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsBrowserTTL"}, "cache": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCache"}, "cache_key": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCacheKey"}, "cache_reserve": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCacheReserve"}, "edge_ttl": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsEdgeTTL"}, "origin_cache_control": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsOriginCacheControl"}, "origin_error_page_passthru": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsOriginErrorPagePassthru"}, "read_timeout": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsReadTimeout"}, "respect_strong_etags": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsRespectStrongEtags"}, "serve_stale": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsServeStale"}, "shared_dictionary": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsSharedDictionary"}, "strip_etags": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsStripETags"}, "strip_last_modified": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsStripLastModified"}, "strip_set_cookie": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsStripSetCookie"}, "vary": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsVary"}}}, "description": {"example": "Configure settings for how the response is cached."}}, "title": "Set Cache Settings Rule"}]}
```
