---
title: InputText
page_id: schema-inputtext-6b863c21
path: schemas
description: Text input content item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InputText

Text input content item

```yaml
{"description": "Text input content item", "example": {"text": "Hello, how can I help you?", "type": "input_text"}, "properties": {"prompt_cache_breakpoint": {"$ref": "#/components/schemas/PromptCacheBreakpoint"}, "text": {"type": "string"}, "type": {"enum": ["input_text"], "type": "string"}}, "required": ["type", "text"], "type": "object"}
```
