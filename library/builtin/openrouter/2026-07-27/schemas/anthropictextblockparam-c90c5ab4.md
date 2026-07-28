---
title: AnthropicTextBlockParam
page_id: schema-anthropictextblockparam-c90c5ab4
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicTextBlockParam

```yaml
{"example": {"text": "Hello, world!", "type": "text"}, "properties": {"cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "citations": {"items": {"discriminator": {"mapping": {"char_location": "#/components/schemas/AnthropicCitationCharLocationParam", "content_block_location": "#/components/schemas/AnthropicCitationContentBlockLocationParam", "page_location": "#/components/schemas/AnthropicCitationPageLocationParam", "search_result_location": "#/components/schemas/AnthropicCitationSearchResultLocationParam", "web_search_result_location": "#/components/schemas/AnthropicCitationWebSearchResultLocationParam"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/AnthropicCitationCharLocationParam"}, {"$ref": "#/components/schemas/AnthropicCitationPageLocationParam"}, {"$ref": "#/components/schemas/AnthropicCitationContentBlockLocationParam"}, {"$ref": "#/components/schemas/AnthropicCitationWebSearchResultLocationParam"}, {"$ref": "#/components/schemas/AnthropicCitationSearchResultLocationParam"}]}, "type": ["array", "null"]}, "text": {"type": "string"}, "type": {"enum": ["text"], "type": "string"}}, "required": ["type", "text"], "type": "object"}
```
