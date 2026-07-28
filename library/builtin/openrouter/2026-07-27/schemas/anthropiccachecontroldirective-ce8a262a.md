---
title: AnthropicCacheControlDirective
page_id: schema-anthropiccachecontroldirective-ce8a262a
path: schemas
description: Enable automatic prompt caching. When set at the top level, the system automatically applies cache breakpoints to the last cacheable block in the request. When set on an individual content block, it marks an explicit cache breakpoint; block-level markers also work on OpenAI models that support explicit prompt caching — OpenRouter converts them to the provider's native format.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCacheControlDirective

Enable automatic prompt caching. When set at the top level, the system automatically applies cache breakpoints to the last cacheable block in the request. When set on an individual content block, it marks an explicit cache breakpoint; block-level markers also work on OpenAI models that support explicit prompt caching — OpenRouter converts them to the provider's native format.

```yaml
{"description": "Enable automatic prompt caching. When set at the top level, the system automatically applies cache breakpoints to the last cacheable block in the request. When set on an individual content block, it marks an explicit cache breakpoint; block-level markers also work on OpenAI models that support explicit prompt caching — OpenRouter converts them to the provider's native format.", "example": {"type": "ephemeral"}, "properties": {"ttl": {"$ref": "#/components/schemas/AnthropicCacheControlTtl"}, "type": {"enum": ["ephemeral"], "type": "string"}}, "required": ["type"], "type": "object"}
```
