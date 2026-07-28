---
title: ChatContentText
page_id: schema-chatcontenttext-26f6ae79
path: schemas
description: Text content part
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentText

Text content part

```yaml
{"description": "Text content part", "example": {"text": "Hello, world!", "type": "text"}, "properties": {"cache_control": {"$ref": "#/components/schemas/ChatContentCacheControl"}, "prompt_cache_breakpoint": {"$ref": "#/components/schemas/PromptCacheBreakpoint"}, "text": {"type": "string"}, "type": {"enum": ["text"], "type": "string"}}, "required": ["type", "text"], "type": "object"}
```
