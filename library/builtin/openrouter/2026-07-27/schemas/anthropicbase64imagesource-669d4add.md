---
title: AnthropicBase64ImageSource
page_id: schema-anthropicbase64imagesource-669d4add
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicBase64ImageSource

```yaml
{"example": {"data": "/9j/4AAQ...", "media_type": "image/jpeg", "type": "base64"}, "properties": {"data": {"type": "string"}, "media_type": {"$ref": "#/components/schemas/AnthropicImageMimeType"}, "type": {"enum": ["base64"], "type": "string"}}, "required": ["type", "media_type", "data"], "type": "object"}
```
