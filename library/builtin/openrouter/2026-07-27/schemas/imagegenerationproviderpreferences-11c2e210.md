---
title: ImageGenerationProviderPreferences
page_id: schema-imagegenerationproviderpreferences-11c2e210
path: schemas
description: Provider routing preferences and provider-specific passthrough configuration.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenerationProviderPreferences

Provider routing preferences and provider-specific passthrough configuration.

```yaml
{"description": "Provider routing preferences and provider-specific passthrough configuration.", "example": {"allow_fallbacks": false, "only": ["google-ai-studio"]}, "properties": {"allow_fallbacks": {"description": "Whether to allow backup providers to serve requests\n- true: (default) when the primary provider (or your custom providers in \"order\") is unavailable, use the next best provider.\n- false: use only the primary/custom provider, and return the upstream error if it's unavailable.\n", "type": ["boolean", "null"]}, "ignore": {"description": "List of provider slugs to ignore. If provided, this list is merged with your account-wide ignored provider settings for this request.", "example": ["openai", "anthropic"], "items": {"anyOf": [{"$ref": "#/components/schemas/ProviderName"}, {"type": "string"}]}, "type": ["array", "null"]}, "only": {"description": "List of provider slugs to allow. If provided, this list is merged with your account-wide allowed provider settings for this request.", "example": ["openai", "anthropic"], "items": {"anyOf": [{"$ref": "#/components/schemas/ProviderName"}, {"type": "string"}]}, "type": ["array", "null"]}, "options": {"allOf": [{"$ref": "#/components/schemas/ProviderOptions"}, {"example": {"black-forest-labs": {"guidance": 3, "steps": 40}}}]}, "order": {"description": "An ordered list of provider slugs. The router will attempt to use the first provider in the subset of this list that supports your requested model, and fall back to the next if it is unavailable. If no providers are available, the request will fail with an error message.", "example": ["openai", "anthropic"], "items": {"anyOf": [{"$ref": "#/components/schemas/ProviderName"}, {"type": "string"}]}, "type": ["array", "null"]}, "sort": {"anyOf": [{"$ref": "#/components/schemas/ProviderSort"}, {"$ref": "#/components/schemas/ProviderSortConfig"}, {"type": "null"}], "description": "The sorting strategy to use for this request, if \"order\" is not specified. When set, no load balancing is performed.", "example": "price"}}, "type": "object"}
```
