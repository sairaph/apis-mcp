---
title: PromptCacheOptions
page_id: schema-promptcacheoptions-a03aa059
path: schemas
description: 'Request-level prompt-cache controls. `mode: "explicit"` disables OpenAI-managed breakpoints so only blocks marked with `prompt_cache_breakpoint` are cached. Only supported by OpenAI GPT-5.6 and newer.'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PromptCacheOptions

Request-level prompt-cache controls. `mode: "explicit"` disables OpenAI-managed breakpoints so only blocks marked with `prompt_cache_breakpoint` are cached. Only supported by OpenAI GPT-5.6 and newer.

```yaml
{"description": "Request-level prompt-cache controls. `mode: \"explicit\"` disables OpenAI-managed breakpoints so only blocks marked with `prompt_cache_breakpoint` are cached. Only supported by OpenAI GPT-5.6 and newer.", "example": {"mode": "explicit", "ttl": "30m"}, "properties": {"mode": {"enum": ["explicit"], "type": "string"}, "ttl": {"type": ["string", "null"]}}, "required": ["mode"], "type": ["object", "null"]}
```
