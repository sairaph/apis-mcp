---
title: AnthropicImageBlockParam
page_id: schema-anthropicimageblockparam-555c9c65
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicImageBlockParam

```yaml
{"example": {"source": {"data": "/9j/4AAQ...", "media_type": "image/jpeg", "type": "base64"}, "type": "image"}, "properties": {"cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "source": {"discriminator": {"mapping": {"base64": "#/components/schemas/AnthropicBase64ImageSource", "url": "#/components/schemas/AnthropicUrlImageSource"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicBase64ImageSource"}, {"$ref": "#/components/schemas/AnthropicUrlImageSource"}]}, "type": {"enum": ["image"], "type": "string"}}, "required": ["type", "source"], "type": "object"}
```
