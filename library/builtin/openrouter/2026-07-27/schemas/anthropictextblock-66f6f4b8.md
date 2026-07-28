---
title: AnthropicTextBlock
page_id: schema-anthropictextblock-66f6f4b8
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextBlock

```yaml
{"example": {"citations": null, "text": "Hello, world!", "type": "text"}, "properties": {"citations": {"items": {"$ref": "#/components/schemas/AnthropicTextCitation"}, "type": ["array", "null"]}, "text": {"type": "string"}, "type": {"enum": ["text"], "type": "string"}}, "required": ["type", "text", "citations"], "type": "object"}
```
