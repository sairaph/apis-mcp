---
title: ChatContentCacheControl
page_id: schema-chatcontentcachecontrol-6fab13f7
path: schemas
description: 'Anthropic-style cache breakpoint for the content part. Interchangeable with the OpenAI-style `prompt_cache_breakpoint` marker: OpenRouter converts between the two based on the provider serving the request.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentCacheControl

Anthropic-style cache breakpoint for the content part. Interchangeable with the OpenAI-style `prompt_cache_breakpoint` marker: OpenRouter converts between the two based on the provider serving the request.

```yaml
{"allOf": [{"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, {"properties": {}, "type": "object"}], "description": "Anthropic-style cache breakpoint for the content part. Interchangeable with the OpenAI-style `prompt_cache_breakpoint` marker: OpenRouter converts between the two based on the provider serving the request.", "example": {"ttl": "5m", "type": "ephemeral"}}
```
